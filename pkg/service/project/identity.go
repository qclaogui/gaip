// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"context"

	pb "github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateUser Create User
func (srv *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	return srv.repoIdentity.CreateUser(ctx, req)
}

// GetUser Get User
func (srv *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	return srv.repoIdentity.GetUser(ctx, req)
}

// ListUsers List Users
func (srv *Server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return srv.repoIdentity.ListUsers(ctx, req)
}

// UpdateUser Update User
func (srv *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	return srv.repoIdentity.UpdateUser(ctx, req)
}

// DeleteUser Delete User
func (srv *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	return srv.repoIdentity.DeleteUser(ctx, req)
}
