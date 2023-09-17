// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package library

import (
	"github.com/qclaogui/golang-api-server/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/golang-api-server/pkg/service/library/repository"
	"google.golang.org/grpc"
)

// The ServiceServer type implements a library server.
//
// ServiceServer is used to implement LibraryServiceServer.
type ServiceServer struct {
	librarypb.UnimplementedLibraryServiceServer

	Cfg Config

	repo repository.Repository
}

func NewServiceServer(cfg Config) (*ServiceServer, error) {
	// Create the ServiceServer
	s := &ServiceServer{Cfg: cfg}

	if err := s.setupRepo(); err != nil {
		return nil, err
	}

	return s, nil
}

func (srv *ServiceServer) setupRepo() error {
	repo, err := repository.NewRepository(srv.Cfg.RepoCfg)
	if err != nil {
		return err
	}

	srv.repo = repo
	return nil
}

func (srv *ServiceServer) RegisterGRPC(s *grpc.Server) {
	s.RegisterService(&librarypb.LibraryService_ServiceDesc, srv)
}
