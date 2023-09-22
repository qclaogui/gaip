// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"context"

	"github.com/qclaogui/golang-api-server/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/golang-api-server/pkg/service"
	"github.com/qclaogui/golang-api-server/pkg/service/project/repository"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Server Project server
type Server interface {
	service.Backend

	projectpb.ProjectServiceServer
}

// The projectServerImpl type implements a project server.
type projectServerImpl struct {
	Cfg Config

	repo repository.Repository
}

func NewProjectServer(cfg Config) (Server, error) {
	s := &projectServerImpl{Cfg: cfg}
	if err := s.setupRepo(); err != nil {
		return nil, err
	}

	return s, nil
}

func (srv *projectServerImpl) setupRepo() error {
	repo, err := repository.NewRepository(srv.Cfg.RepoCfg)
	if err != nil {
		return err
	}

	srv.repo = repo
	return nil
}

func (srv *projectServerImpl) RegisterGRPC(s *grpc.Server) {
	s.RegisterService(&projectpb.ProjectService_ServiceDesc, srv)
}

func (srv *projectServerImpl) CreateProject(ctx context.Context, req *projectpb.CreateProjectRequest) (*projectpb.Project, error) {
	return srv.repo.CreateProject(ctx, req)
}

func (srv *projectServerImpl) GetProject(ctx context.Context, req *projectpb.GetProjectRequest) (*projectpb.Project, error) {
	return srv.repo.GetProject(ctx, req)
}

func (srv *projectServerImpl) ListProjects(ctx context.Context, req *projectpb.ListProjectsRequest) (*projectpb.ListProjectsResponse, error) {
	return srv.repo.ListProjects(ctx, req)
}

func (srv *projectServerImpl) DeleteProject(ctx context.Context, req *projectpb.DeleteProjectRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteProject(ctx, req)
}
