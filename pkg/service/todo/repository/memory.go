// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"context"
	"log/slog"
	"sync"

	"github.com/google/uuid"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MemoryRepo fulfills the Repository interface
type MemoryRepo struct {
	mem map[uuid.UUID]*todopb.ToDo
	mu  sync.Mutex
}

// NewMemoryRepo is a factory function to generate a new repository
func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		mem: make(map[uuid.UUID]*todopb.ToDo),
	}
}

func (m *MemoryRepo) Create(_ context.Context, req *todopb.CreateRequest) (*todopb.CreateResponse, error) {
	// request todo
	todo := req.GetItem()
	if todo.GetTitle() == "" && todo.GetDescription() == "" {
		return nil, status.Error(codes.Unknown, ErrFailedToCreate.Error())
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	id := uuid.NewSHA1(uuid.NameSpaceOID, []byte(todo.GetTitle()+todo.GetDescription()))
	todo.Id = id.String()

	m.mem[id] = todo
	return &todopb.CreateResponse{Api: "v1", Id: todo.Id}, nil
}

func (m *MemoryRepo) Get(_ context.Context, req *todopb.GetRequest) (*todopb.GetResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id, _ := uuid.Parse(req.GetId())

	slog.Warn("Get todo from MemoryRepo", "req_id", req.GetId(), "id", id, "mem", m.mem)
	todo, ok := m.mem[id]
	if !ok {
		return nil, status.Error(codes.Unknown, ErrNotFound.Error())
	}

	return &todopb.GetResponse{Api: "v1", Item: todo}, nil
}

func (m *MemoryRepo) Update(_ context.Context, req *todopb.UpdateRequest) (*todopb.UpdateResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	todo := req.GetItem()
	id, _ := uuid.Parse(todo.GetId())

	if _, ok := m.mem[id]; !ok {
		return nil, status.Error(codes.Unknown, ErrNotFound.Error())
	}

	m.mem[id] = todo
	return &todopb.UpdateResponse{Api: "v1", Updated: 1}, nil
}

func (m *MemoryRepo) Delete(_ context.Context, req *todopb.DeleteRequest) (*todopb.DeleteResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id, _ := uuid.Parse(req.GetId())

	if _, ok := m.mem[id]; !ok {
		return nil, status.Error(codes.Unknown, ErrNotFound.Error())
	}
	delete(m.mem, id)
	return &todopb.DeleteResponse{Api: "v1", Deleted: 1}, nil
}

func (m *MemoryRepo) List(_ context.Context, _ *todopb.ListRequest) (*todopb.ListResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var todos []*todopb.ToDo
	for _, todo := range m.mem {
		todos = append(todos, todo)
	}

	if len(todos) < 1 {
		return nil, status.Error(codes.Unknown, ErrNotFound.Error())
	}
	return &todopb.ListResponse{Api: "v1", Items: todos}, nil
}
