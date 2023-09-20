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

// The Server type implements a bookstore server.
//
// Server is used to implement BookstoreServiceServer.
type Server struct {
	Cfg Config

	repo repository.Repository
}

func NewServiceServer(cfg Config) (*Server, error) {
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
	s.RegisterService(&bookstorepb.BookstoreService_ServiceDesc, srv)
}

func (srv *Server) ListShelves(ctx context.Context, req *emptypb.Empty) (*bookstorepb.ListShelvesResponse, error) {
	return srv.repo.ListShelves(ctx, req)
}

func (srv *Server) CreateShelf(ctx context.Context, req *bookstorepb.CreateShelfRequest) (*bookstorepb.Shelf, error) {
	return srv.repo.CreateShelf(ctx, req)
}
func (srv *Server) GetShelf(ctx context.Context, req *bookstorepb.GetShelfRequest) (*bookstorepb.Shelf, error) {
	return srv.repo.GetShelf(ctx, req)
}
func (srv *Server) DeleteShelf(ctx context.Context, req *bookstorepb.DeleteShelfRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteShelf(ctx, req)

}
func (srv *Server) ListBooks(ctx context.Context, req *bookstorepb.ListBooksRequest) (*bookstorepb.ListBooksResponse, error) {
	return srv.repo.ListBooks(ctx, req)

}
func (srv *Server) CreateBook(ctx context.Context, req *bookstorepb.CreateBookRequest) (*bookstorepb.Book, error) {
	return srv.repo.CreateBook(ctx, req)
}
func (srv *Server) GetBook(ctx context.Context, req *bookstorepb.GetBookRequest) (*bookstorepb.Book, error) {
	return srv.repo.GetBook(ctx, req)
}
func (srv *Server) DeleteBook(ctx context.Context, req *bookstorepb.DeleteBookRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteBook(ctx, req)
}
