// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package todo

import (
	"context"
	"fmt"

	"github.com/go-kit/log/level"
	pb "github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	APIVersion = "v1"
)

func (s *Server) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if APIVersion != api {
			return fmt.Errorf("unsupported API version: service implements API version '%s', but asked for '%s'", APIVersion, api)
		}
	}
	return nil
}

func (s *Server) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	_ = level.Info(s.logger).Log("msg", "[Create] received request")
	defer func() { _ = level.Info(s.logger).Log("msg", "[Create] completed request") }()

	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if err := req.GetItem().GetCreateTime().CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	return s.repo.CreateTodo(ctx, req)
}

func (s *Server) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return s.repo.UpdateTodo(ctx, req)
}

func (s *Server) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return s.repo.GetTodo(ctx, req)
}

func (s *Server) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return s.repo.DeleteTodo(ctx, req)
}

func (s *Server) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return s.repo.ListTodo(ctx, req)
}
