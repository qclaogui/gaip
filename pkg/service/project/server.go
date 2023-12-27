// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"flag"
	"time"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	lg "github.com/qclaogui/gaip/tools/log"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	Log        log.Logger            `yaml:"-"`
	Registerer prometheus.Registerer `yaml:"-"`

	RepoProject  projectpb.ProjectServiceServer  `yaml:"-"`
	RepoIdentity projectpb.IdentityServiceServer `yaml:"-"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "project.enabled", true, "Enables Server Service Server")

	cfg.Log = lg.Logger
}

func (cfg *Config) Validate() error {

	return nil
}

// The Server type implements a projectpb server.
type Server struct {
	projectpb.UnimplementedProjectServiceServer
	projectpb.UnimplementedIdentityServiceServer
	projectpb.UnimplementedEchoServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repoProject  projectpb.ProjectServiceServer
	repoIdentity projectpb.IdentityServiceServer

	nowF func() time.Time
}

func NewServer(cfg Config) (*Server, error) {
	srv := &Server{
		Cfg:          cfg,
		logger:       cfg.Log,
		Registerer:   cfg.Registerer,
		repoProject:  cfg.RepoProject,
		repoIdentity: cfg.RepoIdentity,
	}

	return srv, nil
}
