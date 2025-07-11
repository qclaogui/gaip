// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package a2a

import (
	"flag"

	"github.com/go-kit/log"

	"github.com/prometheus/client_golang/prometheus"
	pb "github.com/qclaogui/gaip/genproto/a2a/apiv1/a2apb"
)

type Config struct {
	Enabled    bool                  `yaml:"enabled"`
	Log        log.Logger            `yaml:"-"`
	Registerer prometheus.Registerer `yaml:"-"`
	Repo       pb.A2AServiceServer   `yaml:"-"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "a2a.enabled", true, "Enables A2A Service Server")
}

func (cfg *Config) Validate() error {
	return nil
}

// The Server type implements a pb server.
type Server struct {
	pb.UnimplementedA2AServiceServer
	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer
	repo       pb.A2AServiceServer
}

func NewServer(cfg Config) (*Server, error) {
	s := &Server{
		Cfg:        cfg,
		logger:     cfg.Log,
		Registerer: cfg.Registerer,
		repo:       cfg.Repo,
	}
	return s, nil
}
