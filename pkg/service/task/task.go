// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package task

import (
	"context"

	"github.com/qclaogui/gaip/genproto/task/apiv1/taskpb"
)

func (s *Server) GetTask(ctx context.Context, in *taskpb.GetTaskRequest) (*taskpb.Task, error) {
	return s.repo.GetTask(ctx, in)
}

func (s *Server) ListTasks(ctx context.Context, in *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	return s.repo.ListTasks(ctx, in)
}

func (s *Server) CreateTask(ctx context.Context, in *taskpb.CreateTaskRequest) (*taskpb.Task, error) {
	return s.repo.CreateTask(ctx, in)
}

func (s *Server) DeleteTask(ctx context.Context, in *taskpb.DeleteTaskRequest) (*taskpb.Task, error) {
	return s.repo.DeleteTask(ctx, in)
}

func (s *Server) UndeleteTask(ctx context.Context, in *taskpb.UndeleteTaskRequest) (*taskpb.Task, error) {
	return s.repo.UndeleteTask(ctx, in)
}

func (s *Server) UpdateTask(ctx context.Context, in *taskpb.UpdateTaskRequest) (*taskpb.Task, error) {
	return s.repo.UpdateTask(ctx, in)
}
