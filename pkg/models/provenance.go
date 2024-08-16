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

// Handle all interaction with the many tables to get provenance

package models

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	zlog "scanoss.com/provenance/pkg/logger"
)

type provenanceModel struct {
	ctx  context.Context
	conn *sqlx.Conn
}

type Provenance struct {
	Type             string `db:"type"`
	PurlName         string `db:"purl_name"`
	VendorName       string `db:"vendor_name"`
	DeclaredLocation string `db:"declared_location"`
	CountriesId      string `db:"countries_id"`
}

// NewProvenanceModel creates a new instance of a provenance Model
func NewProvenanceModel(ctx context.Context, conn *sqlx.Conn) *provenanceModel {
	return &provenanceModel{ctx: ctx, conn: conn}
}

// ProcessCuratedVendors assigns a list of country name to given set of id's of a set of provenance records
func (m *provenanceModel) ProcessCuratedVendors(vendors []Provenance) map[string]map[string]int {
	curatedCountries := make(map[string]map[string]int)
	for _, v := range vendors {
		if v.CountriesId != "" {
			listStr := strings.ReplaceAll(v.CountriesId, "{", "")
			listStr = strings.ReplaceAll(listStr, "}", "")
			list := strings.Split(listStr, ",")
			if len(list) == 0 {
				list = append(list, listStr)
			}
			_, exist := curatedCountries[v.PurlName]
			if !exist {
				curatedCountries[v.PurlName] = make(map[string]int)
			}
			curatedCountries[v.PurlName][list[0]]++
		}
	}
	return (curatedCountries)
}

// GetProvenanceByPurlName get declared and curated locations for contributors and authors from a list of purlnames
func (m *provenanceModel) GetProvenanceByPurlNames(purlNames []string, purlType string) ([]Provenance, error) {
	list := ""
	list = strings.Join(purlNames, "','")
	list = "('" + list + "')"
	var allSources []Provenance
	query := `
		    SELECT DISTINCT
		        gc.purl_name AS purl_name,
		        vd.type AS type,
		        vd.username AS vendor_name,
		        CASE
					WHEN vl.declared_location IS NULL THEN ''
					ELSE
						vl.declared_location
				END AS declared_location,
		        CASE
		            WHEN vl.curated_countries_ids IS NULL THEN ''
		            ELSE
		                CASE
		                    WHEN vl.curated_countries_ids = '{}' THEN ''
		                    ELSE concat(vl.curated_countries_ids)
		                END
		        END AS countries_id
		    FROM vendors vd
		    left JOIN github_contributors gc ON gc.contributor = vd.username
		    left JOIN vendor_locations vl ON vl.vendor_id = vd.id
		    WHERE gc.purl_name IN ` + list + `
		      AND vd.type IS NOT NULL
		      AND vd.mine_id = 5
		      AND vl.declared_location IS NOT NULL;`
	err := m.conn.SelectContext(m.ctx, &allSources, query)
	if err != nil {
		zlog.S.Errorf("Error: Failed to query %v: %+v", purlNames, err)
		return nil, fmt.Errorf("failed to query : %v", err)
	}
	return allSources, nil
}
