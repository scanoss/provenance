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

package dtos

import (
	"encoding/json"
	"errors"
	"fmt"

	zlog "scanoss.com/provenance/pkg/logger"
)

type ProvenanceOutput struct {
	Provenance []ProvenanceOutputItem `json:"purls"`
}

type ProvenanceOutputItem struct {
	Purl              string                   `json:"purl"`
	DeclaredLocations []DeclaredProvenanceItem `json:"declared_locations"`
	CuratedLocations  []CuratedProvenanceItem  `json:"curated_locations"`
}

type DeclaredProvenanceItem struct {
	Type     string `json:"type"`
	Location string `json:"location"`
}

type CuratedProvenanceItem struct {
	Country string `json:"country"`
	Count   int    `json:"count"`
}

// ExportProvenanceOutput converts the CryptoOutput structure to a byte array
func ExportProvenanceOutput(output ProvenanceOutput) ([]byte, error) {
	data, err := json.Marshal(output)
	if err != nil {
		zlog.S.Errorf("Parse failure: %v", err)
		return nil, errors.New("failed to produce JSON from provenance output data")
	}
	return data, nil
}

// ParseCryptoOutput converts the input byte array to a CryptoOutput structure
func ParseCryptoOutput(input []byte) (ProvenanceOutput, error) {
	if input == nil || len(input) == 0 {
		return ProvenanceOutput{}, errors.New("no output Cryptography data supplied to parse")
	}
	var data ProvenanceOutput
	err := json.Unmarshal(input, &data)
	if err != nil {
		zlog.S.Errorf("Parse failure: %v", err)
		return ProvenanceOutput{}, errors.New(fmt.Sprintf("failed to parse Cryptography output data: %v", err))
	}
	zlog.S.Debugf("Parsed data2: %v", data)
	return data, nil
}
