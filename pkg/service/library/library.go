// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package library

import (
	"context"

	"github.com/go-kit/log"
	"github.com/qclaogui/golang-api-server/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/golang-api-server/pkg/service"
	"github.com/qclaogui/golang-api-server/pkg/service/library/repository"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Service Project Service
type Service interface {
	service.Backend

	librarypb.LibraryServiceServer
}

// The libraryServiceImpl type implements a library server.
type libraryServiceImpl struct {
	Cfg Config

	repo repository.Repository

	logger log.Logger
}

func NewLibraryService(cfg Config) (Service, error) {
	// Create the libraryServiceImpl
	s := &libraryServiceImpl{Cfg: cfg}

	if err := s.setupRepo(); err != nil {
		return nil, err
	}

	return s, nil
}

func (srv *libraryServiceImpl) setupRepo() error {
	repo, err := repository.NewRepository(srv.Cfg.RepoCfg)
	if err != nil {
		return err
	}

	srv.repo = repo
	return nil
}

func (srv *libraryServiceImpl) RegisterGRPC(s *grpc.Server) {
	s.RegisterService(&librarypb.LibraryService_ServiceDesc, srv)
}

func (srv *libraryServiceImpl) CreateShelf(ctx context.Context, req *librarypb.CreateShelfRequest) (*librarypb.Shelf, error) {
	return srv.repo.CreateShelf(ctx, req)
}

func (srv *libraryServiceImpl) ListShelves(ctx context.Context, req *librarypb.ListShelvesRequest) (*librarypb.ListShelvesResponse, error) {
	return srv.repo.ListShelves(ctx, req)
}

func (srv *libraryServiceImpl) GetShelf(ctx context.Context, req *librarypb.GetShelfRequest) (*librarypb.Shelf, error) {
	return srv.repo.GetShelf(ctx, req)
}

func (srv *libraryServiceImpl) DeleteShelf(ctx context.Context, req *librarypb.DeleteShelfRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteShelf(ctx, req)
}

func (srv *libraryServiceImpl) MergeShelves(ctx context.Context, req *librarypb.MergeShelvesRequest) (*librarypb.Shelf, error) {
	return srv.repo.MergeShelves(ctx, req)
}

func (srv *libraryServiceImpl) CreateBook(ctx context.Context, req *librarypb.CreateBookRequest) (*librarypb.Book, error) {
	return srv.repo.CreateBook(ctx, req)
}

func (srv *libraryServiceImpl) GetBook(ctx context.Context, req *librarypb.GetBookRequest) (*librarypb.Book, error) {
	return srv.repo.GetBook(ctx, req)
}

func (srv *libraryServiceImpl) ListBooks(ctx context.Context, req *librarypb.ListBooksRequest) (*librarypb.ListBooksResponse, error) {
	return srv.repo.ListBooks(ctx, req)
}

func (srv *libraryServiceImpl) DeleteBook(ctx context.Context, req *librarypb.DeleteBookRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteBook(ctx, req)
}

func (srv *libraryServiceImpl) UpdateBook(ctx context.Context, req *librarypb.UpdateBookRequest) (*librarypb.Book, error) {
	return srv.repo.UpdateBook(ctx, req)

}

func (srv *libraryServiceImpl) MoveBook(ctx context.Context, req *librarypb.MoveBookRequest) (*librarypb.Book, error) {
	return srv.repo.MoveBook(ctx, req)
}
