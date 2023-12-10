// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"context"

	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/gaip/pkg/service"
	"github.com/qclaogui/gaip/pkg/service/project/repository"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Service Project server
type Service interface {
	service.Backend

	projectpb.ProjectServiceServer
}

// The projectServiceImpl type implements a project server.
type projectServiceImpl struct {
	Cfg Config

	repo repository.Repository
}

func NewProjectService(cfg Config) (Service, error) {
	s := &projectServiceImpl{Cfg: cfg}
	if err := s.setupRepo(); err != nil {
		return nil, err
	}

	return s, nil
}

func (srv *projectServiceImpl) setupRepo() error {
	repo, err := repository.NewRepository(srv.Cfg.RepoCfg)
	if err != nil {
		return err
	}

	srv.repo = repo
	return nil
}

func (srv *projectServiceImpl) RegisterGRPC(s *grpc.Server) {
	s.RegisterService(&projectpb.ProjectService_ServiceDesc, srv)
}

func (srv *projectServiceImpl) CreateProject(ctx context.Context, req *projectpb.CreateProjectRequest) (*projectpb.Project, error) {
	return srv.repo.CreateProject(ctx, req)
}

func (srv *projectServiceImpl) GetProject(ctx context.Context, req *projectpb.GetProjectRequest) (*projectpb.Project, error) {
	return srv.repo.GetProject(ctx, req)
}

func (srv *projectServiceImpl) ListProjects(ctx context.Context, req *projectpb.ListProjectsRequest) (*projectpb.ListProjectsResponse, error) {
	return srv.repo.ListProjects(ctx, req)
}

func (srv *projectServiceImpl) DeleteProject(ctx context.Context, req *projectpb.DeleteProjectRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteProject(ctx, req)
}
