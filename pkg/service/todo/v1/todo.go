// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package v1

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-kit/log"
	pb "github.com/qclaogui/golang-api-server/genproto/todo/apiv1/todopb"
	"github.com/qclaogui/golang-api-server/pkg/service/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

// Option is an alias for a function that will take in a pointer to
// an ServiceServer and modify it
type Option func(*ServiceServer) error

// WithRepository applies a given repository to the ServiceServer
func WithRepository(repo todo.Repository) Option {
	return func(srv *ServiceServer) error {
		srv.repo = repo
		return nil
	}
}

// WithMemoryRepository applies a memory repository to the Option
func WithMemoryRepository() Option {
	repo := todo.NewMemoryRepository()
	return WithRepository(repo)
}

// WithMysqlRepository applies a memory repository to the Option
func WithMysqlRepository(dsn string) Option {
	// Create the memory repo, if we needed parameters, such as connection
	// strings they could be inputted here
	return func(svc *ServiceServer) error {
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

type ServiceServer struct {
	pb.UnimplementedToDoServiceServer
	repo   todo.Repository
	logger log.Logger
}

func NewServiceServer(logger log.Logger, opts ...Option) (*ServiceServer, error) {
	// Create the ServiceServer
	s := &ServiceServer{logger: logger}
	// Apply all Configurations passed in
	for _, opt := range opts {
		// Pass the service into the config option function
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (srv *ServiceServer) RegisterGRPC(s *grpc.Server) {
	s.RegisterService(&pb.ToDoService_ServiceDesc, srv)
}

func (srv *ServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return fmt.Errorf("unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (srv *ServiceServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if err := req.GetItem().GetCreatedAt().CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	return srv.repo.Create(ctx, req)
}

func (srv *ServiceServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return srv.repo.Update(ctx, req)
}

func (srv *ServiceServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return srv.repo.Get(ctx, req)
}

func (srv *ServiceServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return srv.repo.Delete(ctx, req)
}

func (srv *ServiceServer) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return srv.repo.List(ctx, req)
}
