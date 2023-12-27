// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"context"

	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (srv *Server) CreateProject(ctx context.Context, req *projectpb.CreateProjectRequest) (*projectpb.Project, error) {
	return srv.repoProject.CreateProject(ctx, req)
}

func (srv *Server) GetProject(ctx context.Context, req *projectpb.GetProjectRequest) (*projectpb.Project, error) {
	return srv.repoProject.GetProject(ctx, req)
}

func (srv *Server) ListProjects(ctx context.Context, req *projectpb.ListProjectsRequest) (*projectpb.ListProjectsResponse, error) {
	return srv.repoProject.ListProjects(ctx, req)
}

func (srv *Server) DeleteProject(ctx context.Context, req *projectpb.DeleteProjectRequest) (*emptypb.Empty, error) {
	return srv.repoProject.DeleteProject(ctx, req)
}
