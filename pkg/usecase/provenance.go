// SPDX-License-Identifier: GPL-2.0-or-later
/*
 * Copyright (C) 2018-2022 SCANOSS.COM
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 2 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	"scanoss.com/provenance/pkg/dtos"
	zlog "scanoss.com/provenance/pkg/logger"
	"scanoss.com/provenance/pkg/models"
	i "scanoss.com/provenance/pkg/models"
	utils "scanoss.com/provenance/pkg/utils"
)

type ProvenanceUseCase struct {
	ctx     context.Context
	conn    *sqlx.Conn
	allUrls *models.AllUrlsModel
	prov    *models.ContributorProvenanceModel
}

func NewProvenance(ctx context.Context, conn *sqlx.Conn) *ProvenanceUseCase {
	return &ProvenanceUseCase{ctx: ctx, conn: conn,
		allUrls: models.NewAllUrlModel(ctx, conn, models.NewProjectModel(ctx, conn)),
		prov:    models.NewContributorProvenanceModel(ctx, conn),
	}
}

// GetProvenance takes the Provenance Input request, searches for Provenance usages and returns a ProvenanceOutput struct
func (p ProvenanceUseCase) GetProvenance(request dtos.ProvenanceInput) (dtos.ProvenanceOutput, int, error) {

	notFound := 0
	if len(request.Purls) == 0 {
		zlog.S.Info("Empty List of Purls supplied")
		return dtos.ProvenanceOutput{}, 0, errors.New("empty list of purls")
	}

	query := []i.InternalQuery{}
	purlsToQuery := []utils.PurlReq{}
	justPurlNames := []string{}
	//Prepare purls to query
	for _, purl := range request.Purls {

		purlReq := strings.Split(purl.Purl, "@") // Remove any version specific info from the PURL
		if purlReq[0] == "" {
			continue
		}
		if len(purlReq) > 1 {
			purl.Requirement = purlReq[1]
		}

		purlName, err := utils.PurlNameFromString(purl.Purl) // Make sure we just have the bare minimum for a Purl Name
		if err == nil {
			purlsToQuery = append(purlsToQuery, utils.PurlReq{Purl: purlName, Version: purl.Requirement})
			justPurlNames = append(justPurlNames, purlName)
		}
		query = append(query, i.InternalQuery{CompletePurl: purl.Purl, Requirement: purl.Requirement, PurlName: purlName})
	}

	clResults := models.QueryBulkPurlLDB(query)

	url, err := p.allUrls.GetUrlsByPurlList(purlsToQuery)
	if len(url) == 0 {
		return dtos.ProvenanceOutput{}, 0, errors.New("Error Processing input")
	}
	zlog.S.Info("\n\nAbout to query contributors\n", justPurlNames)
	contributorProvs, _ := p.prov.GetContributorsByPurlList(justPurlNames)
	mapContributors := make(map[string][]string)
	for _, p := range contributorProvs {
		mapContributors[p.PurlName] = append(mapContributors[p.PurlName], p.Country)
	}

	purlMap := make(map[string][]models.AllUrl)

	///Order Urls in a map for fast access by purlname
	for r := range url {
		purlMap[url[r].PurlName] = append(purlMap[url[r].PurlName], url[r])
	}
	urlHashes := []string{}
	// For all the requested purls, choose the closest urls that match
	for r := range query {
		query[r].SelectedURLS, err = models.PickClosestUrls(purlMap[query[r].PurlName], query[r].PurlName, "", query[r].Requirement)
		if err != nil {
			return dtos.ProvenanceOutput{}, 0, err
		}
		if len(query[r].SelectedURLS) > 0 {
			query[r].SelectedVersion = query[r].SelectedURLS[0].Version
			for h := range query[r].SelectedURLS {
				urlHashes = append(urlHashes, query[r].SelectedURLS[h].UrlHash)

			}
		} else {
			// NO URL linked to that purl
			notFound++
		}
	}
	//Create a map containing the files for each url
	files, errFiles := models.QueryBulkPivotLDB(urlHashes)
	if errFiles != nil {
		return dtos.ProvenanceOutput{}, 0, errFiles
	}
	//Create a map containing the Provenance usage for each file
	prov := models.QueryBulkProvenanceLDB(files)

	mapProv := make(map[string][]models.ProvenanceItem)

	//Remove duplicate algorithms for the same file
	for k, v := range files {
		for f := range v {
			mapProv[k] = append(mapProv[k], prov[v[f]]...)
		}
	}
	retV := dtos.ProvenanceOutput{}

	//Create the response
	for r := range query {
		var provOutItem dtos.ProvenanceOutputItem
		countries := make(map[string]bool)
		relatedURLs := query[r].SelectedURLS
		provOutItem.Version = query[r].SelectedVersion
		provOutItem.Purl = query[r].CompletePurl
		for u := range relatedURLs {

			hash := relatedURLs[u].UrlHash
			items := mapProv[hash]
			//remove duplicates for the same URL
			for i := range items {
				if _, exist := countries[items[i].Country]; !exist {
					provOutItem.Countries = append(provOutItem.Countries, dtos.ProvenanceItem{Country: items[i].Country, Source: items[i].Source})
					countries[items[i].Country] = true
				}
			}
		}
		if clRes, exist := clResults[provOutItem.Purl]; exist {
			for _, v := range clRes {
				if v.Country != "" {
					provOutItem.Countries = append(provOutItem.Countries, dtos.ProvenanceItem{Country: v.Country, Source: "Component Declared"})
				}

			}
		}

		pn, _ := utils.PurlNameFromString(provOutItem.Purl)
		ctrs := mapContributors[pn]
		for _, ctr := range ctrs {
			if ctr != "" {
				provOutItem.Countries = append(provOutItem.Countries, dtos.ProvenanceItem{Country: ctr, Source: "Contributor Declared"})
			}
		}

		retV.Provenance = append(retV.Provenance, provOutItem)
	}
	return retV, notFound, nil
}
