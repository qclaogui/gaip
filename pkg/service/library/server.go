// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package library

import (
	"flag"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	pb "github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	Log        log.Logger            `yaml:"-"`
	Registerer prometheus.Registerer `yaml:"-"`

	Repo pb.LibraryServiceServer `yaml:"-"`

	// CacheCfg holds the configuration used for the cache.
	// CacheCfg cache.BackendConfig `yaml:"cache"`
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

// The Server type implements a pb server.
type Server struct {
	pb.UnimplementedLibraryServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo pb.LibraryServiceServer
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
