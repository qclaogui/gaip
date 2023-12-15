// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package library

import (
	"context"
	"flag"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/gaip/pkg/service/library/repository"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Service Library Service Server
type Service interface {
	librarypb.LibraryServiceServer
}

type Config struct {
	//RepoCfg holds the configuration used for the repository.
	RepoCfg repository.Config `yaml:"database"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	//Register RepoCfg Config
	cfg.RepoCfg.RegisterFlags(fs)
}

func (cfg *Config) Validate() error {
	//Validate RepoCfg Config
	if err := cfg.RepoCfg.Validate(); err != nil {
		return err
	}
	return nil
}

// The libraryServiceImpl type implements a library server.
type libraryServiceImpl struct {
	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo repository.Repository
}

func NewLibraryService(cfg Config, logger log.Logger, reg prometheus.Registerer) (Service, error) {
	// Create the libraryServiceImpl
	srv := &libraryServiceImpl{
		Cfg:        cfg,
		logger:     logger,
		Registerer: reg,
	}

	if err := srv.setupRepo(); err != nil {
		return nil, err
	}

	return srv, nil
}

func (srv *libraryServiceImpl) setupRepo() error {
	var err error
	if srv.repo, err = repository.NewRepository(srv.Cfg.RepoCfg); err != nil {
		return err
	}
	return nil
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
