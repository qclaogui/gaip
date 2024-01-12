// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"context"

	pb "github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (srv *Server) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*pb.Project, error) {
	return srv.repoProject.CreateProject(ctx, req)
}

func (srv *Server) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.Project, error) {
	return srv.repoProject.GetProject(ctx, req)
}

func (srv *Server) ListProjects(ctx context.Context, req *pb.ListProjectsRequest) (*pb.ListProjectsResponse, error) {
	return srv.repoProject.ListProjects(ctx, req)
}

func (srv *Server) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*emptypb.Empty, error) {
	return srv.repoProject.DeleteProject(ctx, req)
}
