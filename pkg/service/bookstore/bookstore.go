// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package bookstore

import (
	"context"
	"flag"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	Repo repository.Bookstore `yaml:"-"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "bookstore.enabled", false, "Enables Bookstore Service Server")
}

func (cfg *Config) Validate() error {
	return nil
}

// The ServiceServer type implements a bookstorepb server.
type ServiceServer struct {
	bookstorepb.UnimplementedBookstoreServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo repository.Bookstore
}

func New(cfg Config, s *service.Server) (*ServiceServer, error) {
	srv := &ServiceServer{
		Cfg:        cfg,
		logger:     s.Log,
		Registerer: s.Registerer,
		repo:       cfg.Repo,
	}

	bookstorepb.RegisterBookstoreServiceServer(s.GRPCServer, srv)
	return srv, nil
}

func (s *ServiceServer) ListShelves(ctx context.Context, req *emptypb.Empty) (*bookstorepb.ListShelvesResponse, error) {
	return s.repo.ListShelves(ctx, req)
}

func (s *ServiceServer) CreateShelf(ctx context.Context, req *bookstorepb.CreateShelfRequest) (*bookstorepb.Shelf, error) {
	return s.repo.CreateShelf(ctx, req)
}
func (s *ServiceServer) GetShelf(ctx context.Context, req *bookstorepb.GetShelfRequest) (*bookstorepb.Shelf, error) {
	return s.repo.GetShelf(ctx, req)
}
func (s *ServiceServer) DeleteShelf(ctx context.Context, req *bookstorepb.DeleteShelfRequest) (*emptypb.Empty, error) {
	return s.repo.DeleteShelf(ctx, req)

}
func (s *ServiceServer) ListBooks(ctx context.Context, req *bookstorepb.ListBooksRequest) (*bookstorepb.ListBooksResponse, error) {
	return s.repo.ListBooks(ctx, req)

}
func (s *ServiceServer) CreateBook(ctx context.Context, req *bookstorepb.CreateBookRequest) (*bookstorepb.Book, error) {
	return s.repo.CreateBook(ctx, req)
}
func (s *ServiceServer) GetBook(ctx context.Context, req *bookstorepb.GetBookRequest) (*bookstorepb.Book, error) {
	return s.repo.GetBook(ctx, req)
}
func (s *ServiceServer) DeleteBook(ctx context.Context, req *bookstorepb.DeleteBookRequest) (*emptypb.Empty, error) {
	return s.repo.DeleteBook(ctx, req)
}
