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

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	common "github.com/scanoss/papi/api/commonv2"
	pb "github.com/scanoss/papi/api/provenancev2"
	myconfig "scanoss.com/provenance/pkg/config"
	"scanoss.com/provenance/pkg/dtos"
	zlog "scanoss.com/provenance/pkg/logger"
	"scanoss.com/provenance/pkg/models"
)

func TestCProvenanceServer_Echo(t *testing.T) {
	ctx := context.Background()
	err := zlog.NewSugaredDevLogger()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a sugared logger", err)
	}
	defer zlog.SyncZap()
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer models.CloseDB(db)
	myConfig, err := myconfig.NewServerConfig(nil)
	if err != nil {
		t.Fatalf("failed to load Config: %v", err)
	}
	s := NewProvenanceServer(db, myConfig)

	type args struct {
		ctx context.Context
		req *common.EchoRequest
	}
	tests := []struct {
		name    string
		s       pb.ProvenanceServer
		args    args
		want    *common.EchoResponse
		wantErr bool
	}{
		{
			name: "Echo",
			s:    s,
			args: args{
				ctx: ctx,
				req: &common.EchoRequest{Message: "Hello there!"},
			},
			want: &common.EchoResponse{Message: "Hello there!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Echo(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Echo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Echo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCProvenanceServer_GetProvenance(t *testing.T) {
	ctx := context.Background()
	err := zlog.NewSugaredDevLogger()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a sugared logger", err)
	}
	defer zlog.SyncZap()
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer models.CloseDB(db)
	ctx = ctxzap.ToContext(ctx, zlog.L)
	models.RegisterConcat(db, ctx)
	err = models.LoadTestSqlData(db, nil, nil)
	if err != nil {
		fmt.Println(err)
	}

	myConfig, err := myconfig.NewServerConfig(nil)
	if err != nil {
		t.Fatalf("failed to load Config: %v", err)
	}

	s := NewProvenanceServer(db, myConfig)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when loading test data", err)
	}
	request := common.PurlRequest{Purls: []*common.PurlRequest_Purls{&common.PurlRequest_Purls{Purl: "pkg:github/scanoss/engine"}, &common.PurlRequest_Purls{Purl: "pkg:github/torvalds/uemacs"}}}

	got, errReq := s.GetComponentProvenance(ctx, &request)
	if errReq != nil {
		t.Logf("unexpected error on request %+v", errReq)
	}
	var rcv dtos.ProvenanceOutput
	jsonOut, errResp := json.Marshal(got)
	if errResp != nil {
		t.Logf("unexpected error on unmarshalling response %+v", errResp)
	}
	err = json.Unmarshal(jsonOut, &rcv)
	if err != nil {
		t.Logf("unexpected error on unmarshalling to a dto %+v", err)
	}
	if len(rcv.Provenance) == 0 {
		t.Error("expected to get 1 result")

	} else {
		firstPurl := rcv.Provenance[0]
		if len(firstPurl.DeclaredLocations) == 0 {
			t.Error("expected to get at least 1 declared location")
		} else if len(firstPurl.CuratedLocations) == 0 {
			t.Error("expected to get at least 1 curated location")
		} else {
			firstCuratedCountry := firstPurl.CuratedLocations[0]
			if firstCuratedCountry.Country != "Argentina" && firstCuratedCountry.Country != "Spain" && firstCuratedCountry.Country != "Afghanistan" {
				t.Errorf("Curated country (%s) was not expected", firstCuratedCountry.Country)
			}
		}

	}

}
