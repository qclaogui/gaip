// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.
//go:build requires_docker

package todo_test

import (
	"context"
	"testing"
	"time"

	todo "github.com/qclaogui/gaip/genproto/todo/apiv1"
	pb "github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Clients are initialized in main_test.go.
var (
	todoGRPC *todo.Client
	todoREST *todo.Client
)

func TestTodoCRUD(t *testing.T) {
	ctx := context.Background()

	for _, client := range map[string]*todo.Client{
		"grpc": todoGRPC,
		// "rest": todoREST,
	} {
		tm := time.Now().UTC().Add(time.Minute)
		reminder := timestamppb.New(tm)
		item := &pb.ToDo{
			Title:       "title",
			Description: "description",
			CreatedAt:   reminder,
		}
		create := &pb.CreateTodoRequest{
			Api:  "v1",
			Item: item,
		}

		td, err := client.CreateTodo(ctx, create)
		if err != nil {
			t.Fatalf("client.Create() failed: %v", err)
		}

		if td.GetId() == "" {
			t.Errorf("client.Create().Id was unexpectedly empty")
		}
	}
}
