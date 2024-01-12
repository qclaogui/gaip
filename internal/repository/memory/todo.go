// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFound = errors.New("the item was not found in the repository")

	ErrFailedToCreate = errors.New("failed to add the todo to the repository")
)

// NewTodo is a factory function to generate a new repository
func NewTodo() (todopb.ToDoServiceServer, error) {
	m := &todoMemImpl{
		mem: make(map[uuid.UUID]*todopb.ToDo),
	}
	return m, nil
}

// todoMemImpl fulfills the Repository todoMemImpl interface
// All objects are managed in an in-memory non-persistent store.
//
// todoMemImpl is used to implement todopb.ToDoServiceServer.
type todoMemImpl struct {
	todopb.UnimplementedToDoServiceServer

	mem map[uuid.UUID]*todopb.ToDo
	mu  sync.Mutex
}

func (m *todoMemImpl) Create(_ context.Context, req *todopb.CreateRequest) (*todopb.CreateResponse, error) {
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

func (m *todoMemImpl) Get(_ context.Context, req *todopb.GetRequest) (*todopb.GetResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id, _ := uuid.Parse(req.GetId())

	//slog.Warn("Get todo from todoMemImpl", "req_id", req.GetId(), "id", id, "mem", m.mem)
	todo, ok := m.mem[id]
	if !ok {
		return nil, status.Error(codes.Unknown, ErrNotFound.Error())
	}

	return &todopb.GetResponse{Api: "v1", Item: todo}, nil
}

func (m *todoMemImpl) Update(_ context.Context, req *todopb.UpdateRequest) (*todopb.UpdateResponse, error) {
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

func (m *todoMemImpl) Delete(_ context.Context, req *todopb.DeleteRequest) (*todopb.DeleteResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id, _ := uuid.Parse(req.GetId())

	if _, ok := m.mem[id]; !ok {
		return nil, status.Error(codes.Unknown, ErrNotFound.Error())
	}
	delete(m.mem, id)
	return &todopb.DeleteResponse{Api: "v1", Deleted: 1}, nil
}

func (m *todoMemImpl) List(_ context.Context, _ *todopb.ListRequest) (*todopb.ListResponse, error) {
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
