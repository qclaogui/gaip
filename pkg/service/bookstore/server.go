// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package bookstore

import (
	"flag"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
	lg "github.com/qclaogui/gaip/tools/log"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	Log        log.Logger            `yaml:"-"`
	Registerer prometheus.Registerer `yaml:"-"`

	Repo bookstorepb.BookstoreServiceServer `yaml:"-"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "bookstore.enabled", false, "Enables Bookstore Service Server")

	cfg.Log = lg.Logger

}

func (cfg *Config) Validate() error {
	return nil
}

// The Server type implements a bookstorepb server.
type Server struct {
	bookstorepb.UnimplementedBookstoreServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo bookstorepb.BookstoreServiceServer
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
