// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package bookstore

import (
	"context"
	"flag"

	"github.com/go-kit/log"
	"github.com/grafana/dskit/cache"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
	"github.com/qclaogui/gaip/pkg/service"
	"github.com/qclaogui/gaip/pkg/service/bookstore/repository"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Config struct {
	//RepoCfg holds the configuration used for the repository.
	RepoCfg repository.Config `yaml:"database"`

	// CacheCfg holds the configuration used for the cache.
	CacheCfg cache.BackendConfig `yaml:"cache"`
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

// The bookstoreServerImpl type implements a bookstore server.
type bookstoreServerImpl struct {
	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo repository.Repository
}

func New(cfg Config, s *service.Server) error {
	srv := &bookstoreServerImpl{
		Cfg:        cfg,
		logger:     s.Log,
		Registerer: s.Registerer,
	}

	if err := srv.setupRepo(); err != nil {
		return err
	}

	bookstorepb.RegisterBookstoreServiceServer(s.GRPCServer, srv)
	return nil
}

func (srv *bookstoreServerImpl) setupRepo() error {
	var err error
	if srv.repo, err = repository.NewRepository(srv.Cfg.RepoCfg); err != nil {
		return err
	}
	return nil
}

func (srv *bookstoreServerImpl) ListShelves(ctx context.Context, req *emptypb.Empty) (*bookstorepb.ListShelvesResponse, error) {
	return srv.repo.ListShelves(ctx, req)
}

func (srv *bookstoreServerImpl) CreateShelf(ctx context.Context, req *bookstorepb.CreateShelfRequest) (*bookstorepb.Shelf, error) {
	return srv.repo.CreateShelf(ctx, req)
}
func (srv *bookstoreServerImpl) GetShelf(ctx context.Context, req *bookstorepb.GetShelfRequest) (*bookstorepb.Shelf, error) {
	return srv.repo.GetShelf(ctx, req)
}
func (srv *bookstoreServerImpl) DeleteShelf(ctx context.Context, req *bookstorepb.DeleteShelfRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteShelf(ctx, req)

}
func (srv *bookstoreServerImpl) ListBooks(ctx context.Context, req *bookstorepb.ListBooksRequest) (*bookstorepb.ListBooksResponse, error) {
	return srv.repo.ListBooks(ctx, req)

}
func (srv *bookstoreServerImpl) CreateBook(ctx context.Context, req *bookstorepb.CreateBookRequest) (*bookstorepb.Book, error) {
	return srv.repo.CreateBook(ctx, req)
}
func (srv *bookstoreServerImpl) GetBook(ctx context.Context, req *bookstorepb.GetBookRequest) (*bookstorepb.Book, error) {
	return srv.repo.GetBook(ctx, req)
}
func (srv *bookstoreServerImpl) DeleteBook(ctx context.Context, req *bookstorepb.DeleteBookRequest) (*emptypb.Empty, error) {
	return srv.repo.DeleteBook(ctx, req)
}
