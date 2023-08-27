package v1

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/qclaogui/golang-api-server/pkg/service/todo"

	pb "github.com/qclaogui/golang-api-server/pkg/api/todopb/v1"

	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

// TodoOption is an alias for a function that will take in a pointer to
// an ToDoService and modify it
type TodoOption func(svc *ToDoService) error

// WithRepository applies a given todo repository to the ToDoService
func WithRepository(repo todo.Repository) TodoOption {
	return func(svc *ToDoService) error {
		svc.repo = repo
		return nil
	}
}

// WithMemoryRepository applies a memory todo repository to the ConfigOption
func WithMemoryRepository() TodoOption {
	repo := todo.NewMemoryRepository()
	return WithRepository(repo)
}

// WithMysqlRepository applies a memory todo repository to the ConfigOption
func WithMysqlRepository(dsn string) TodoOption {
	// Create the memory repo, if we needed parameters, such as connection
	// strings they could be inputted here
	return func(svc *ToDoService) error {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return fmt.Errorf("failed to open database: %v", err)
		}

		repo, err := todo.NewMysqlRepository(db)
		if err != nil {
			return err
		}

		svc.repo = repo
		return nil
	}
}

type ToDoService struct {
	repo todo.Repository
	db   *sql.DB
}

func NewToDoService(opts ...TodoOption) (pb.ToDoServiceServer, error) {
	// Create the ToDoService
	s := &ToDoService{}
	// Apply all Configurations passed in
	for _, opt := range opts {
		// Pass the service into the configuration function
		if err := opt(s); err != nil {
			return nil, err
		}

	}
	return s, nil
}

func (s *ToDoService) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return fmt.Errorf("unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (s *ToDoService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if err := req.ToDo.Reminder.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	id, _ := uuid.Parse(req.ToDo.Id)
	td, err := s.repo.Create(ctx, &todo.Todo{
		ID:          id,
		Title:       req.ToDo.Title,
		Description: req.ToDo.Description,
		Reminder:    req.ToDo.Reminder.AsTime(),
	})

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &pb.CreateResponse{Api: apiVersion, Id: td.ID.String()}, nil
}

func (s *ToDoService) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	id, _ := uuid.Parse(req.ToDo.Id)
	rows, err := s.repo.Update(ctx, &todo.Todo{
		ID:          id,
		Title:       req.ToDo.Title,
		Description: req.ToDo.Description,
		Reminder:    req.ToDo.Reminder.AsTime(),
	})

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &pb.UpdateResponse{Api: apiVersion, Updated: rows}, nil
}

func (s *ToDoService) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	id, _ := uuid.Parse(req.Id)
	td, err := s.repo.Read(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	ptd := &pb.ToDo{
		Id:          td.ID.String(),
		Title:       td.Title,
		Description: td.Description,
		Reminder:    timestamppb.New(td.Reminder),
	}

	return &pb.ReadResponse{Api: apiVersion, ToDo: ptd}, nil
}

func (s *ToDoService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	id, _ := uuid.Parse(req.Id)
	rows, err := s.repo.Delete(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &pb.DeleteResponse{Api: apiVersion, Deleted: rows}, nil
}

func (s *ToDoService) ReadAll(ctx context.Context, req *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {

	// 1.check api version
	if err := s.checkAPI(req.Api); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	todos, err := s.repo.ReadAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	var list []*pb.ToDo
	for _, td := range todos {
		ptd := &pb.ToDo{
			Id:          td.ID.String(),
			Title:       td.Title,
			Description: td.Description,
			Reminder:    timestamppb.New(td.Reminder),
		}
		list = append(list, ptd)
	}

	return &pb.ReadAllResponse{Api: apiVersion, ToDos: list}, nil
}
