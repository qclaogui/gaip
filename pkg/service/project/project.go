// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"context"

	"github.com/qclaogui/golang-api-server/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/golang-api-server/pkg/service/project/repository"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	Cfg Config

	repo repository.Repository
}

func NewServer(cfg Config) (*Server, error) {
	// Create the Server
	s := &Server{Cfg: cfg}

	if err := s.setupRepo(); err != nil {
		return nil, err
	}

	return s, nil
}

func (srv *Server) setupRepo() error {
	repo, err := repository.NewRepository(srv.Cfg.RepoCfg)
	if err != nil {
		return err
	}

	srv.repo = repo
	return nil
}

func (srv *Server) RegisterGRPC(s *grpc.Server) {
	s.RegisterService(&projectpb.ProjectService_ServiceDesc, srv)
}

func (srv *Server) CreateProject(ctx context.Context, request *projectpb.CreateProjectRequest) (*projectpb.Project, error) {
	_ = ctx
	_ = request
	//TODO implement me
	panic("implement me")
}

func (srv *Server) GetProject(ctx context.Context, request *projectpb.GetProjectRequest) (*projectpb.Project, error) {
	_ = ctx
	_ = request
	//TODO implement me
	panic("implement me")
}

func (srv *Server) ListProjects(ctx context.Context, request *projectpb.ListProjectsRequest) (*projectpb.ListProjectsResponse, error) {
	_ = ctx
	_ = request
	//TODO implement me
	panic("implement me")
}

func (srv *Server) DeleteProject(ctx context.Context, request *projectpb.DeleteProjectRequest) (*emptypb.Empty, error) {
	_ = ctx
	_ = request
	//TODO implement me
	panic("implement me")
}
