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

package models

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	zlog "scanoss.com/provenance/pkg/logger"
)

type ContributorProvenanceModel struct {
	ctx  context.Context
	conn *sqlx.Conn
}

type CountributorProvenance struct {
	Contributor string `db:"contributor"`
	Country     string `db:"country"`
	PurlName    string `db:"purl_name"`
}

// NewContributorProvenanceModel creates a new instance of the All URL Model
func NewContributorProvenanceModel(ctx context.Context, conn *sqlx.Conn) *ContributorProvenanceModel {
	return &ContributorProvenanceModel{ctx: ctx, conn: conn}
}

/*
func (c *ContributorProvenanceModel) GetContributorProvenanceByPurlList(list []utils.PurlReq) ([]AllUrl, error) {
	if len(list) == 0 {
		zlog.S.Errorf("Please specify a valid Purl list to query")
		return []AllUrl{}, errors.New("please specify a valid Purl list to query")
	}
	purlNames := []string{}

	for p := range list {
		purlNames = append(purlNames, "'"+list[p].Purl+"'")
	}
	inStmt := strings.Join(purlNames, ",")
	inStmt = "(" + inStmt + ")"
	stmt := "SELECT package_hash AS url_hash, component, v.version_name AS version, v.semver AS semver, m.purl_type as purl_type, " +
		"purl_name, mine_id FROM all_urls u " +
		"LEFT JOIN mines m ON u.mine_id = m.id " +
		"LEFT JOIN versions v ON u.version_id = v.id " +
		"WHERE u.purl_name in " + inStmt +
		" and package_hash!= '' ORDER BY date DESC;"

	var allUrls []AllUrl
	err := m.conn.SelectContext(m.ctx, &allUrls, stmt)
	if err != nil {
		zlog.S.Errorf("Failed to query a list of urls:  %v", err)
		return []AllUrl{}, fmt.Errorf("failed to query the all urls table: %v", err)
	}
	//zlog.S.Debugf("Found %v results for %v, %v.", len(allUrls), purlType, purlName)
	return allUrls, nil

}
*/

// GetUrlsByPurlString searches for component details of the specified Purl string (and optional requirement)
func (c *ContributorProvenanceModel) GetContributorsByPurlList(purlNames []string) ([]CountributorProvenance, error) {
	if len(purlNames) == 0 {
		zlog.S.Errorf("Please specify a valid Purl String to query")
		return []CountributorProvenance{}, errors.New("please specify a valid Purl String to query")
	}

	if len(purlNames) > 0 { // No version specified, but we might have a specific version in the Requirement
		inStatement := "('" + strings.Join(purlNames, "','") + "')"
		var prov []CountributorProvenance
		err := c.conn.SelectContext(c.ctx, &prov,
			"select a.contributor as contributor, a.purl_name as purl_name, c.country as country "+
				"from github_contributors a, vendors b, locations c "+
				"where a.contributor = b.username and b.location_id=c.id and a.purl_name in "+inStatement+" and c.country!='';")

		zlog.S.Info("select a.purl_name,a.contributor, b.location, c.id, c.country " +
			"from github_contributors a, vendors b, locations c " +
			"where a.contributor = b.username and b.location_id=c.id and a.purl_name in " + inStatement + ";")

		if err != nil {
			//zlog.S.Errorf("Failed to query Location table for %v  %v", purlName, err)
			return []CountributorProvenance{}, fmt.Errorf("failed to query the Location table: %v", err)
		}
		return prov, nil
	}
	return []CountributorProvenance{}, nil

}
