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
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"scanoss.com/provenance/pkg/dtos"
	zlog "scanoss.com/provenance/pkg/logger"
	"scanoss.com/provenance/pkg/models"
	"scanoss.com/provenance/pkg/utils"
)

type ProvenanceUseCase struct {
	ctx     context.Context
	conn    *sqlx.Conn
	allUrls *models.AllUrlsModel
}
type ProvenanceWorkerStruct struct {
	URLMd5  string
	Purl    string
	Version string
}
type InternalQuery struct {
	CompletePurl    string
	PurlName        string
	Requirement     string
	SelectedVersion string
	SelectedURLS    []models.AllUrl
}

func NewProvenance(ctx context.Context, conn *sqlx.Conn) *ProvenanceUseCase {
	return &ProvenanceUseCase{ctx: ctx, conn: conn,
		allUrls: models.NewAllUrlModel(ctx, conn, models.NewProjectModel(ctx, conn)),
	}
}

// GetCrypto takes the Crypto Input request, searches for Crytporaphic usages and returns a CrytoOutput struct
func (p ProvenanceUseCase) GetProvenance(request dtos.ProvenanceInput) (dtos.ProvenanceOutput, int, error) {

	notFound := 0
	if len(request.Purls) == 0 {
		zlog.S.Info("Empty List of Purls supplied")
		return dtos.ProvenanceOutput{}, 0, errors.New("empty list of purls")
	}

	query := []InternalQuery{}
	purlsToQuery := []utils.PurlReq{}
	purls := []string{}
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
			purls = append(purls, purlName)
		}
		query = append(query, InternalQuery{CompletePurl: purl.Purl, Requirement: purl.Requirement, PurlName: purlName})
	}
	prov := models.NewProvenanceModel(p.ctx, p.conn)
	countries := models.NewCountryMapModel(p.ctx, p.conn)

	vendors, err := prov.GetProvenanceByPurlNames(purls, "")
	if err != nil {
		return dtos.ProvenanceOutput{}, 0, err
	}
	curatedCountries := prov.ProcessCuratedVendors(vendors)

	vendorsMap := make(map[string][]models.Provenance)

	for _, v := range vendors {
		vendorsMap[v.PurlName] = append(vendorsMap[v.PurlName], v)
	}
	/*
		 mapProv := make(map[string][]models.ProvenanceItem)

		 //Remove duplicate algorithms for the same file
		 for k, v := range files {
			 for f := range v {
				 mapProv[k] = append(mapProv[k], prov[v[f]]...)
			 }
		 }*/
	retV := dtos.ProvenanceOutput{}

	//Create the response
	for k, listOfVendors := range vendorsMap {
		var provOutItem dtos.ProvenanceOutputItem

		//provOutItem.Version = query[r].SelectedVersion
		provOutItem.Purl = k //query[r].CompletePurl
		for _, vendor := range listOfVendors {
			if vendor.DeclaredLocation != "" {
				provOutItem.DeclaredLocations = append(provOutItem.DeclaredLocations, dtos.DeclaredProvenanceItem{Type: vendor.Type, Location: vendor.DeclaredLocation})
			}
		}

		//add curated values
		for k, v := range curatedCountries[provOutItem.Purl] {
			i, err := strconv.Atoi(k)
			if err == nil {
				countryName, err := countries.GeCountryById(i)
				if err == nil {
					provOutItem.CuratedLocations = append(provOutItem.CuratedLocations, dtos.CuratedProvenanceItem{Country: countryName, Count: v})
				}
			}
		}

		retV.Provenance = append(retV.Provenance, provOutItem)

	}
	return retV, notFound, nil
}
