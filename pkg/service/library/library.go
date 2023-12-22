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
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	Repo repository.Library `yaml:"-"`

	// CacheCfg holds the configuration used for the cache.
	//CacheCfg cache.BackendConfig `yaml:"cache"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "library.enabled", true, "Enables Library Service Server")
}

func (cfg *Config) Validate() error {
	//Validate RepoCfg Config
	//if err := cfg.CacheCfg.Validate(); err != nil {
	//	return err
	//}
	return nil
}

// The ServiceServer type implements a librarypb server.
type ServiceServer struct {
	librarypb.UnimplementedLibraryServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo repository.Library
}

func New(cfg Config, s *service.Server) (*ServiceServer, error) {
	srv := &ServiceServer{
		Cfg:        cfg,
		logger:     s.Log,
		Registerer: s.Registerer,
		repo:       cfg.Repo,
	}

	librarypb.RegisterLibraryServiceServer(s.GRPCServer, srv)
	return srv, nil
}

func (s *ServiceServer) CreateShelf(ctx context.Context, req *librarypb.CreateShelfRequest) (*librarypb.Shelf, error) {
	return s.repo.CreateShelf(ctx, req)
}

func (s *ServiceServer) ListShelves(ctx context.Context, req *librarypb.ListShelvesRequest) (*librarypb.ListShelvesResponse, error) {
	return s.repo.ListShelves(ctx, req)
}

func (s *ServiceServer) GetShelf(ctx context.Context, req *librarypb.GetShelfRequest) (*librarypb.Shelf, error) {
	return s.repo.GetShelf(ctx, req)
}

func (s *ServiceServer) DeleteShelf(ctx context.Context, req *librarypb.DeleteShelfRequest) (*emptypb.Empty, error) {
	return s.repo.DeleteShelf(ctx, req)
}

func (s *ServiceServer) MergeShelves(ctx context.Context, req *librarypb.MergeShelvesRequest) (*librarypb.Shelf, error) {
	return s.repo.MergeShelves(ctx, req)
}

func (s *ServiceServer) CreateBook(ctx context.Context, req *librarypb.CreateBookRequest) (*librarypb.Book, error) {
	return s.repo.CreateBook(ctx, req)
}

func (s *ServiceServer) GetBook(ctx context.Context, req *librarypb.GetBookRequest) (*librarypb.Book, error) {
	return s.repo.GetBook(ctx, req)
}

func (s *ServiceServer) ListBooks(ctx context.Context, req *librarypb.ListBooksRequest) (*librarypb.ListBooksResponse, error) {
	return s.repo.ListBooks(ctx, req)
}

func (s *ServiceServer) DeleteBook(ctx context.Context, req *librarypb.DeleteBookRequest) (*emptypb.Empty, error) {
	return s.repo.DeleteBook(ctx, req)
}

func (s *ServiceServer) UpdateBook(ctx context.Context, req *librarypb.UpdateBookRequest) (*librarypb.Book, error) {
	return s.repo.UpdateBook(ctx, req)

}

func (s *ServiceServer) MoveBook(ctx context.Context, req *librarypb.MoveBookRequest) (*librarypb.Book, error) {
	return s.repo.MoveBook(ctx, req)
}
