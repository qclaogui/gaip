// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package routeguide

import (
	"flag"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
	lg "github.com/qclaogui/gaip/tools/log"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	Log        log.Logger            `yaml:"-"`
	Registerer prometheus.Registerer `yaml:"-"`

	Repo routeguidepb.RouteGuideServiceServer `yaml:"-"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "routeguide.enabled", true, "Enables RouteGuide Service Server")

	cfg.Log = lg.Logger
}

func (cfg *Config) Validate() error {
	return nil
}

// The Server type implements a routeguidepb server.
type Server struct {
	routeguidepb.UnimplementedRouteGuideServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo routeguidepb.RouteGuideServiceServer
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
