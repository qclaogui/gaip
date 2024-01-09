// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package todo

import (
	"flag"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	Log        log.Logger            `yaml:"-"`
	Registerer prometheus.Registerer `yaml:"-"`

	Repo todopb.ToDoServiceServer `yaml:"-"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "todo.enabled", true, "Enables Todo Service Server")
}

func (cfg *Config) Validate() error {
	return nil
}

// The Server type implements a todopb service server.
type Server struct {
	todopb.UnimplementedToDoServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo todopb.ToDoServiceServer
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
