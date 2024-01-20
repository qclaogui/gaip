// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.
//go:build requires_docker

package generativeai_test

import (
	"context"
	"testing"

	generativelanguage "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1"
	pb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1/generativelanguagepb"
)

// Clients are initialized in main_test.go.
var (
	modelGRPC *generativelanguage.ModelClient
	modelREST *generativelanguage.ModelClient
)

func TestListModels(t *testing.T) {
	t.Skip()

	list := &pb.ListModelsRequest{
		PageSize: 5,
	}

	for typ, client := range map[string]*generativelanguage.ModelClient{
		"grpc": modelGRPC,
		"rest": modelREST,
	} {

		it := client.ListModels(context.Background(), list)
		if mSize := it.PageInfo().MaxSize; mSize != int(list.PageSize) {
			t.Errorf("%s client.ListModels(): it.PageInfo().MaxSize = %d, want %d", typ, mSize, list.PageSize)
		}
	}
}
