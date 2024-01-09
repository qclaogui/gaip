// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ListOperations returns a fixed response matching the given PageSize if the resource name is not blank
func (srv *Server) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest) (*longrunningpb.ListOperationsResponse, error) {
	_ = ctx

	if req.Name == "" {
		return nil, status.Error(codes.NotFound, "cannot list operation without a name.")
	}

	var operations []*longrunningpb.Operation
	if req.PageSize > 0 {
		for i := 1; i < int(req.PageSize); i++ {
			var result *longrunningpb.Operation_Response
			if i%2 == 0 {
				result = &longrunningpb.Operation_Response{}
			}

			operations = append(operations, &longrunningpb.Operation{
				Name:   "the/thing/" + strconv.Itoa(i),
				Done:   result != nil,
				Result: result,
			})
		}
	} else {
		operations = append(operations, &longrunningpb.Operation{
			Name: "a/pending/thing",
			Done: false,
		})
	}

	return &longrunningpb.ListOperationsResponse{
		Operations: operations,
	}, nil
}

func (srv *Server) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest) (*longrunningpb.Operation, error) {
	if op, err := srv.handleWaitOperation(ctx, req); op != nil || err != nil {
		return op, err
	}

	return nil, status.Errorf(codes.NotFound, "Operation %q not found.", req.Name)
}

// DeleteOperation returns a successful response if the resource name is not blank
func (srv *Server) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest) (*emptypb.Empty, error) {
	_ = ctx

	if req.Name == "" {
		return nil, status.Error(codes.NotFound, "cannot delete operation without a name.")
	}
	return &emptypb.Empty{}, nil
}

// CancelOperation returns a successful response if the resource name is not blank
func (srv *Server) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest) (*emptypb.Empty, error) {
	_ = ctx

	if req.Name == "" {
		return nil, status.Error(codes.NotFound, "cannot cancel operation without a name.")
	}
	return &emptypb.Empty{}, nil
}

// WaitOperation randomly waits and returns an operation with the same name
func (srv *Server) WaitOperation(ctx context.Context, req *longrunningpb.WaitOperationRequest) (*longrunningpb.Operation, error) {
	_ = ctx

	if req.Name == "" {
		return nil, status.Error(codes.NotFound, "cannot wait on a operation without a name.")
	}

	num := rand.Intn(500)
	time.Sleep(time.Duration(num) * time.Millisecond)

	var result *longrunningpb.Operation_Response
	if num%2 == 0 {
		result = &longrunningpb.Operation_Response{}
	}

	return &longrunningpb.Operation{
		Name:   req.Name,
		Done:   result != nil,
		Result: result,
	}, nil
}
