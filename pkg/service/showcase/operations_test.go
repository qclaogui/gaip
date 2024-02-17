// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package showcase

import (
	"context"
	"testing"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCancelOperation(t *testing.T) {
	srv, _ := NewServer(Config{})
	_, err := srv.CancelOperation(context.Background(), &longrunningpb.CancelOperationRequest{
		Name: "a/thing",
	})

	if err != nil {
		t.Error("CancelOperation should have been successful")
	}
}

func TestCancelOperation_notFound(t *testing.T) {
	srv, _ := NewServer(Config{})
	_, err := srv.CancelOperation(context.Background(), &longrunningpb.CancelOperationRequest{})

	s, _ := status.FromError(err)
	if s.Code() != codes.NotFound {
		t.Errorf("CancelOperation expected code=%d, got %d", codes.NotFound, s.Code())
	}
}
