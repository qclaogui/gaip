// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package task_test

import (
	"context"
	"testing"

	task "github.com/qclaogui/gaip/genproto/task/apiv1"
	pb "github.com/qclaogui/gaip/genproto/task/apiv1/taskpb"
)

// Clients are initialized in main_test.go.
var (
	taskWriterGRPC *task.TasksWriterClient
	taskWriterREST *task.TasksWriterClient

	taskReaderGRPC *task.TasksReaderClient
	taskReaderREST *task.TasksReaderClient
)

func TestTaskCRUD(t *testing.T) {
	_ = taskReaderGRPC
	_ = taskReaderREST

	ctx := context.Background()

	for _, client := range map[string]*task.TasksWriterClient{
		"grpc": taskWriterGRPC,
		"rest": taskWriterREST,
	} {

		item := &pb.Task{
			Title:       "title",
			Description: "description",
		}

		create := &pb.CreateTaskRequest{
			Task: item,
		}

		tk, err := client.CreateTask(ctx, create)
		if err != nil {
			t.Fatalf("client.Create() failed: %v", err)
		}

		if tk.GetId() == 0 {
			t.Errorf("client.Create().Id was unexpectedly empty")
		}
	}
}
