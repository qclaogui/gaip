// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package library

import (
	"context"

	"github.com/qclaogui/golang-api-server/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/golang-api-server/pkg/service"
	"github.com/qclaogui/golang-api-server/pkg/service/library/repository"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Server Project Server
type Server interface {
	service.Backend

	librarypb.LibraryServiceServer
}

// The libraryServerImpl type implements a library server.
type libraryServerImpl struct {
	Cfg Config

	repo repository.Repository
}

func NewLibraryServer(cfg Config) (Server, error) {
	// Create the libraryServerImpl
	s := &libraryServerImpl{Cfg: cfg}

	if err := s.setupRepo(); err != nil {
		return nil, err
	}

	return s, nil
}

func (srv *libraryServerImpl) setupRepo() error {
	repo, err := repository.NewRepository(srv.Cfg.RepoCfg)
	if err != nil {
		return err
	}

	srv.repo = repo
	return nil
}

func (srv *libraryServerImpl) RegisterGRPC(s *grpc.Server) {
	s.RegisterService(&librarypb.LibraryService_ServiceDesc, srv)
}

func (srv *libraryServerImpl) CreateShelf(ctx context.Context, req *librarypb.CreateShelfRequest) (*librarypb.Shelf, error) {
	return srv.repo.CreateShelf(ctx, req)
}

func (srv *libraryServerImpl) ListShelves(ctx context.Context, req *librarypb.ListShelvesRequest) (*librarypb.ListShelvesResponse, error) {
	return srv.repo.ListShelves(ctx, req)
}

func (srv *libraryServerImpl) GetShelf(ctx context.Context, req *librarypb.GetShelfRequest) (*librarypb.Shelf, error) {
	return srv.repo.GetShelf(ctx, req)
}

func (srv *libraryServerImpl) DeleteShelf(ctx context.Context, req *librarypb.DeleteShelfRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteShelf(ctx, req)
}

func (srv *libraryServerImpl) MergeShelves(ctx context.Context, req *librarypb.MergeShelvesRequest) (*librarypb.Shelf, error) {
	return srv.repo.MergeShelves(ctx, req)
}

func (srv *libraryServerImpl) CreateBook(ctx context.Context, req *librarypb.CreateBookRequest) (*librarypb.Book, error) {
	return srv.repo.CreateBook(ctx, req)
}

func (srv *libraryServerImpl) GetBook(ctx context.Context, req *librarypb.GetBookRequest) (*librarypb.Book, error) {
	return srv.repo.GetBook(ctx, req)
}

func (srv *libraryServerImpl) ListBooks(ctx context.Context, req *librarypb.ListBooksRequest) (*librarypb.ListBooksResponse, error) {
	return srv.repo.ListBooks(ctx, req)
}

func (srv *libraryServerImpl) DeleteBook(ctx context.Context, req *librarypb.DeleteBookRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteBook(ctx, req)
}

func (srv *libraryServerImpl) UpdateBook(ctx context.Context, req *librarypb.UpdateBookRequest) (*librarypb.Book, error) {
	return srv.repo.UpdateBook(ctx, req)

}

func (srv *libraryServerImpl) MoveBook(ctx context.Context, req *librarypb.MoveBookRequest) (*librarypb.Book, error) {
	return srv.repo.MoveBook(ctx, req)
}
