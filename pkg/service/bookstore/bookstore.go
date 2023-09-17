// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package bookstore

import (
	"context"

	"github.com/qclaogui/golang-api-server/genproto/bookstore/apiv1alpha1/bookstorepb"
	"github.com/qclaogui/golang-api-server/pkg/service/bookstore/repository"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// The ServiceServer type implements a bookstore server.
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
	s.RegisterService(&bookstorepb.BookstoreService_ServiceDesc, srv)
}

func (srv *ServiceServer) ListShelves(ctx context.Context, req *emptypb.Empty) (*bookstorepb.ListShelvesResponse, error) {
	return srv.repo.ListShelves(ctx, req)
}

func (srv *ServiceServer) CreateShelf(ctx context.Context, req *bookstorepb.CreateShelfRequest) (*bookstorepb.Shelf, error) {
	return srv.repo.CreateShelf(ctx, req)
}
func (srv *ServiceServer) GetShelf(ctx context.Context, req *bookstorepb.GetShelfRequest) (*bookstorepb.Shelf, error) {
	return srv.repo.GetShelf(ctx, req)
}
func (srv *ServiceServer) DeleteShelf(ctx context.Context, req *bookstorepb.DeleteShelfRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteShelf(ctx, req)

}
func (srv *ServiceServer) ListBooks(ctx context.Context, req *bookstorepb.ListBooksRequest) (*bookstorepb.ListBooksResponse, error) {
	return srv.repo.ListBooks(ctx, req)

}
func (srv *ServiceServer) CreateBook(ctx context.Context, req *bookstorepb.CreateBookRequest) (*bookstorepb.Book, error) {
	return srv.repo.CreateBook(ctx, req)
}
func (srv *ServiceServer) GetBook(ctx context.Context, req *bookstorepb.GetBookRequest) (*bookstorepb.Book, error) {
	return srv.repo.GetBook(ctx, req)
}
func (srv *ServiceServer) DeleteBook(ctx context.Context, req *bookstorepb.DeleteBookRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteBook(ctx, req)
}
