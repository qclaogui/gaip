// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package memory

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/qclaogui/gaip/genproto/task/apiv1/taskpb"
)

func NewTasks(cfg Config) (taskpb.TasksServiceServer, error) {
	m := &taskImpl{
		mem: cfg.TaskMemDB,
	}

	return m, nil
}

// taskImpl fulfills the Repository taskImpl interface
// All objects are managed in an in-memory non-persistent store.
//
// taskImpl is used to implement taskpb.TasksServiceServer, taskpb.TasksReaderServiceServer.
type taskImpl struct {
	taskpb.UnimplementedTasksServiceServer
	mem map[uuid.UUID]*taskpb.Task
	mu  sync.Mutex
}

func (t *taskImpl) CreateTask(ctx context.Context, in *taskpb.CreateTaskRequest) (*taskpb.Task, error) {
	_ = ctx
	_ = in

	t.mu.Lock()
	defer t.mu.Unlock()

	// TODO implement me
	panic("implement me")
}

func (t *taskImpl) DeleteTask(ctx context.Context, in *taskpb.DeleteTaskRequest) (*taskpb.Task, error) {
	_ = ctx
	_ = in
	// TODO implement me
	panic("implement me")
}

func (t *taskImpl) UndeleteTask(ctx context.Context, in *taskpb.UndeleteTaskRequest) (*taskpb.Task, error) {
	_ = ctx
	_ = in
	// TODO implement me
	panic("implement me")
}

func (t *taskImpl) UpdateTask(ctx context.Context, in *taskpb.UpdateTaskRequest) (*taskpb.Task, error) {
	_ = ctx
	_ = in
	// TODO implement me
	panic("implement me")
}

func (t *taskImpl) GetTask(ctx context.Context, in *taskpb.GetTaskRequest) (*taskpb.Task, error) {
	_ = ctx
	_ = in
	// TODO implement me
	panic("implement me")
}

func (t *taskImpl) ListTasks(ctx context.Context, in *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	_ = ctx
	_ = in
	// TODO implement me
	panic("implement me")
}
