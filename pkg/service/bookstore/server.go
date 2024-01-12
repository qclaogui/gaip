// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package bookstore

import (
	"flag"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	pb "github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	Log        log.Logger            `yaml:"-"`
	Registerer prometheus.Registerer `yaml:"-"`

	Repo pb.BookstoreServiceServer `yaml:"-"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "bookstore.enabled", false, "Enables Bookstore Service Server")
}

func (cfg *Config) Validate() error {
	return nil
}

// The Server type implements a pb server.
type Server struct {
	pb.UnimplementedBookstoreServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo pb.BookstoreServiceServer
}

func NewServer(cfg Config) (*Server, error) {
	srv := &Server{
		Cfg:        cfg,
		logger:     cfg.Log,
		Registerer: cfg.Registerer,
		repo:       cfg.Repo,
	}

	return srv, nil
}
