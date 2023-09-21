// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package library

import (
	"context"

	"github.com/qclaogui/golang-api-server/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/golang-api-server/pkg/service/library/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// The Server type implements a library server.
//
// Server is used to implement LibraryServiceServer.
type Server struct {
	librarypb.UnimplementedLibraryServiceServer

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
	s.RegisterService(&librarypb.LibraryService_ServiceDesc, srv)
}

func (srv *Server) CreateShelf(context.Context, **librarypb.CreateShelfRequest) (**librarypb.Shelf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShelf not implemented")
}

func (srv *Server) ListShelves(ctx context.Context, req *librarypb.ListShelvesRequest) (*librarypb.ListShelvesResponse, error) {
	return srv.repo.ListShelves(ctx, req)
}
