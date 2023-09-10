// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package v1alpha1

import (
	"context"

	pb "github.com/qclaogui/golang-api-server/api/gen/proto/bookstore/v1alpha1"
	"github.com/qclaogui/golang-api-server/pkg/service/bookstore/repository"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// The ServiceServer type implements a bookstore server.
// All objects are managed in an in-memory non-persistent store.
//
// ServiceServer is used to implement BookstoreServiceServer.
type ServiceServer struct {
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
	s.RegisterService(&pb.BookstoreService_ServiceDesc, srv)
}

func (srv *ServiceServer) ListShelves(ctx context.Context, req *emptypb.Empty) (*pb.ListShelvesResponse, error) {
	return srv.repo.ListShelves(ctx, req)
}

func (srv *ServiceServer) CreateShelf(ctx context.Context, req *pb.CreateShelfRequest) (*pb.Shelf, error) {
	return srv.repo.CreateShelf(ctx, req)
}
func (srv *ServiceServer) GetShelf(ctx context.Context, req *pb.GetShelfRequest) (*pb.Shelf, error) {
	return srv.repo.GetShelf(ctx, req)
}
func (srv *ServiceServer) DeleteShelf(ctx context.Context, req *pb.DeleteShelfRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteShelf(ctx, req)

}
func (srv *ServiceServer) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	return srv.repo.ListBooks(ctx, req)

}
func (srv *ServiceServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.Book, error) {
	return srv.repo.CreateBook(ctx, req)
}
func (srv *ServiceServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.Book, error) {
	return srv.repo.GetBook(ctx, req)
}
func (srv *ServiceServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteBook(ctx, req)
}
