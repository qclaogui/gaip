// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package todo

import (
	"context"
	"log/slog"
	"sync"

	"github.com/google/uuid"
	pb "github.com/qclaogui/golang-api-server/api/gen/proto/todo/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var apiVersion = "v1"

// MemoryRepository fulfills the Repository interface
type MemoryRepository struct {
	mem map[uuid.UUID]*pb.ToDo
	mu  sync.Mutex
}

// NewMemoryRepository is a factory function to generate a new repository
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		mem: make(map[uuid.UUID]*pb.ToDo),
	}
}

func (m *MemoryRepository) Create(_ context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
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
	return &pb.CreateResponse{Api: apiVersion, Id: todo.Id}, nil
}

func (m *MemoryRepository) Get(_ context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id, _ := uuid.Parse(req.GetId())

	slog.Warn("RepoCfg info", "req_id", req.GetId(), "id", id, "mem", m.mem)
	todo, ok := m.mem[id]
	if !ok {
		return nil, status.Error(codes.Unknown, ErrNotFound.Error())
	}

	return &pb.GetResponse{Api: apiVersion, Item: todo}, nil
}

func (m *MemoryRepository) Update(_ context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	todo := req.GetItem()
	id, _ := uuid.Parse(todo.GetId())

	if _, ok := m.mem[id]; !ok {
		return nil, status.Error(codes.Unknown, ErrNotFound.Error())
	}

	m.mem[id] = todo
	return &pb.UpdateResponse{Api: apiVersion, Updated: 1}, nil
}

func (m *MemoryRepository) Delete(_ context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id, _ := uuid.Parse(req.GetId())

	if _, ok := m.mem[id]; !ok {
		return nil, status.Error(codes.Unknown, ErrNotFound.Error())
	}
	delete(m.mem, id)
	return &pb.DeleteResponse{Api: apiVersion, Deleted: 1}, nil
}

func (m *MemoryRepository) List(_ context.Context, _ *pb.ListRequest) (*pb.ListResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var todos []*pb.ToDo
	for _, todo := range m.mem {
		todos = append(todos, todo)
	}

	if len(todos) < 1 {
		return nil, status.Error(codes.Unknown, ErrNotFound.Error())
	}
	return &pb.ListResponse{Api: apiVersion, Items: todos}, nil
}
