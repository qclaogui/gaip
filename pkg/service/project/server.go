// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"flag"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	pb "github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	Log        log.Logger            `yaml:"-"`
	Registerer prometheus.Registerer `yaml:"-"`

	RepoProject pb.ProjectServiceServer `yaml:"-"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "project.enabled", true, "Enables Project Service Server")
}

func (cfg *Config) Validate() error {
	return nil
}

// The Server type implements a pb server.
type Server struct {
	pb.UnimplementedProjectServiceServer

	longrunningpb.UnimplementedOperationsServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repoProject pb.ProjectServiceServer

	nowF func() time.Time
}

func NewServer(cfg Config) (*Server, error) {
	srv := &Server{
		Cfg:        cfg,
		logger:     cfg.Log,
		Registerer: cfg.Registerer,

		repoProject: cfg.RepoProject,

		nowF: time.Now,
	}

	return srv, nil
}
