// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"context"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/gaip/pkg/service"
	"github.com/qclaogui/gaip/pkg/service/project/repository"
	"google.golang.org/protobuf/types/known/emptypb"
)

// The Project type implements a projectpb server.
type Project struct {
	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo repository.Repository
}

func New(cfg Config, s *service.Server) (*Project, error) {
	srv := &Project{
		Cfg:        cfg,
		logger:     s.Log,
		Registerer: s.Registerer,
	}

	if err := srv.setupRepo(); err != nil {
		return nil, err
	}

	projectpb.RegisterProjectServiceServer(s.GRPCServer, srv)
	//projectpb.RegisterIdentityServiceServer(g.Server.GRPCServer, nil)
	//projectpb.RegisterEchoServiceServer(g.Server.GRPCServer, nil)
	return srv, nil
}

func (srv *Project) setupRepo() error {
	var err error
	if srv.repo, err = repository.NewRepository(srv.Cfg.RepoCfg); err != nil {
		return err
	}
	return nil
}

func (srv *Project) CreateProject(ctx context.Context, req *projectpb.CreateProjectRequest) (*projectpb.Project, error) {
	return srv.repo.CreateProject(ctx, req)
}

func (srv *Project) GetProject(ctx context.Context, req *projectpb.GetProjectRequest) (*projectpb.Project, error) {
	return srv.repo.GetProject(ctx, req)
}

func (srv *Project) ListProjects(ctx context.Context, req *projectpb.ListProjectsRequest) (*projectpb.ListProjectsResponse, error) {
	return srv.repo.ListProjects(ctx, req)
}

func (srv *Project) DeleteProject(ctx context.Context, req *projectpb.DeleteProjectRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteProject(ctx, req)
}
