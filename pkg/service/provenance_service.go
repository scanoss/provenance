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

// Package service implements the gRPC service endpoints
package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	common "github.com/scanoss/papi/api/commonv2"
	pb "github.com/scanoss/papi/api/provenancev2"
	myconfig "scanoss.com/provenance/pkg/config"
	zlog "scanoss.com/provenance/pkg/logger"
	"scanoss.com/provenance/pkg/usecase"
)

type provenanceServer struct {
	pb.ProvenanceServer
	db     *sqlx.DB
	config *myconfig.ServerConfig
}

// NewProvenanceServer creates a new instance of Provenance Server
func NewProvenanceServer(db *sqlx.DB, config *myconfig.ServerConfig) pb.ProvenanceServer {
	return &provenanceServer{db: db, config: config}
}

// Echo sends back the same message received
func (p provenanceServer) Echo(ctx context.Context, request *common.EchoRequest) (*common.EchoResponse, error) {
	zlog.S.Infof("Received (%v): %v", ctx, request.GetMessage())
	return &common.EchoResponse{Message: request.GetMessage()}, nil
}

func (p provenanceServer) GetComponentProvenance(ctx context.Context, request *common.PurlRequest) (*pb.ProvenanceResponse, error) {

	// Make sure we have Provenance data to query
	reqPurls := request.GetPurls()
	if reqPurls == nil || len(reqPurls) == 0 {
		statusResp := common.StatusResponse{Status: common.StatusCode_FAILED, Message: "No purls in request data supplied"}
		return &pb.ProvenanceResponse{Status: &statusResp}, errors.New("no purl data supplied")
	}
	dtoRequest, err := convertProvenanceInput(request) // Convert to internal DTO for processing
	if err != nil {
		statusResp := common.StatusResponse{Status: common.StatusCode_FAILED, Message: "Problem parsing Provenance input data"}
		return &pb.ProvenanceResponse{Status: &statusResp}, errors.New("problem parsing Provenance input data")
	}
	conn, err := p.db.Connx(ctx) // Get a connection from the pool
	if err != nil {
		zlog.S.Errorf("Failed to get a database connection from the pool: %v", err)
		statusResp := common.StatusResponse{Status: common.StatusCode_FAILED, Message: "Failed to get database pool connection"}
		return &pb.ProvenanceResponse{Status: &statusResp}, errors.New("problem getting database pool connection")
	}
	defer closeDbConnection(conn)
	// Search the KB for information about each Provenance
	provUc := usecase.NewProvenance(ctx, conn)
	dtoProv, notFound, err := provUc.GetProvenance(dtoRequest)

	if err != nil {
		zlog.S.Errorf("Failed to get provenance: %v", err)
		statusResp := common.StatusResponse{Status: common.StatusCode_FAILED, Message: "Problems encountered extracting Provenance data"}
		return &pb.ProvenanceResponse{Status: &statusResp}, nil
	}
	//zlog.S.Debugf("Parsed Provenance: %+v", dtoProv)
	provResponse, err := convertProvenanceOutput(dtoProv) // Convert the internal data into a response object
	if err != nil {
		zlog.S.Errorf("Failed to covnert parsed dependencies: %v", err)
		statusResp := common.StatusResponse{Status: common.StatusCode_FAILED, Message: "Problems encountered extracting Provenance data"}
		return &pb.ProvenanceResponse{Status: &statusResp}, nil
	}
	// Set the status and respond with the data

	statusResp := common.StatusResponse{Status: common.StatusCode_SUCCESS, Message: "Success"}
	if notFound > 0 {
		statusResp.Status = common.StatusCode_SUCCEEDED_WITH_WARNINGS
		statusResp.Message = fmt.Sprintf("No information found for %d purl(s)", notFound)
	}
	return &pb.ProvenanceResponse{Purls: provResponse.Purls, Status: &statusResp}, nil
}

// closeDbConnection closes the specified database connection
func closeDbConnection(conn *sqlx.Conn) {
	zlog.S.Debugf("Closing DB Connection: %+v", conn)
	err := conn.Close()
	if err != nil {
		zlog.S.Warnf("Warning: Problem closing database connection: %v", err)
	}
}
