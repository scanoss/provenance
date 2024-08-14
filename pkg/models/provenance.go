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

// Handle all interaction with the projects table

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
	ComponentVendor  string `db:"vendor"`
	PurlName         string `db:"purl_name"`
	VendorName       string `db:"vendor_name"`
	DeclaredLocation string `db:"declared_location"`
	CountriesId      string `db:"countries_id"`
}

// NewProjectModel creates a new instance of the Project Model
func NewProvenanceModel(ctx context.Context, conn *sqlx.Conn) *provenanceModel {
	return &provenanceModel{ctx: ctx, conn: conn}
}

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
	//fmt.Println("curated Countries", curatedCountries)
	return (curatedCountries)
}

// GetProjectsByPurlName searches the projects' table for details about Purl Name and Type
func (m *provenanceModel) GetProvenanceByPurlNames(purlNames []string, purlType string) ([]Provenance, error) {
	list := ""
	list = strings.Join(purlNames, "','")
	list = "('" + list + "')"
	//fmt.Println(list)
	var allSources []Provenance
	/*err := m.conn.SelectContext(m.ctx, &allSources, "SELECT  distinct gc.purl_name as purl_name, vd.type as type, vd.username as vendor_name , vl.declared_location  as declared_location, "+
	" (CASE 	WHEN vl.curated_countries_ids IS NULL THEN '' "+
	"ELSE vl.curated_countries_ids "+
	"END AS countries_id) as countries_id "+
	"FROM vendors vd "+
	"left join github_contributors gc on gc.contributor =vd.username  "+
	"left join vendor_locations vl on vl.vendor_id = vd.id	"+
	"WHERE gc.purl_name  in "+list+" and vd.type is not null and vd.mine_id = 5  and vl.declared_location is not null")
	*/

	query := `
    SELECT DISTINCT 
        gc.purl_name AS purl_name, 
        vd.type AS type, 
        vd.username AS vendor_name, 
        vl.declared_location AS declared_location, 
        CASE 
            WHEN vl.curated_countries_ids IS NULL THEN '' 
            ELSE 
                CASE 
                    WHEN vl.curated_countries_ids = '{}' THEN ''
                    ELSE concat(vl.curated_countries_ids)
                END
        END AS countries_id
    FROM vendors vd  
    LEFT JOIN github_contributors gc ON gc.contributor = vd.username  
    LEFT JOIN vendor_locations vl ON vl.vendor_id = vd.id
    WHERE gc.purl_name IN (` + list + `)
      AND vd.type IS NOT NULL 
      AND vd.mine_id = 5  
      AND vl.declared_location IS NOT NULL;`
	err := m.conn.SelectContext(m.ctx, &allSources, query)
	if err != nil {
		zlog.S.Errorf("Error: Failed to query %v, %v: %v", purlNames, purlType, err)
		return nil, fmt.Errorf("failed to query : %v", err)
	}
	return allSources, nil
}
