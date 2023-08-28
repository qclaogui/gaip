package todo

import (
	"context"
	"sync"

	"github.com/google/uuid"
	pb "github.com/qclaogui/golang-api-server/pkg/api/todopb/v1"
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
	todo := req.GetToDo()
	if todo.GetTitle() == "" && todo.GetDescription() == "" {
		return nil, status.Error(codes.Unknown, ErrFailedToCreate.Error())
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	id := uuid.NewSHA1(uuid.NameSpaceOID, []byte(todo.GetTitle()+todo.GetDescription()))
	m.mem[id] = todo
	return &pb.CreateResponse{Api: apiVersion, Id: id.String()}, nil
}

func (m *MemoryRepository) Read(_ context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id, _ := uuid.Parse(req.GetId())
	todo, ok := m.mem[id]
	if !ok {
		return nil, status.Error(codes.Unknown, ErrNotFound.Error())
	}

	return &pb.ReadResponse{Api: apiVersion, ToDo: todo}, nil
}

func (m *MemoryRepository) Update(_ context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	todo := req.GetToDo()
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

func (m *MemoryRepository) ReadAll(_ context.Context, _ *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var todos []*pb.ToDo
	for _, todo := range m.mem {
		todos = append(todos, todo)
	}

	if len(todos) < 1 {
		return nil, status.Error(codes.Unknown, ErrNotFound.Error())
	}
	return &pb.ReadAllResponse{Api: apiVersion, ToDos: todos}, nil
}
