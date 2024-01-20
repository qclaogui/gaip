// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package generativeai

import (
	"flag"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	pb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1/generativelanguagepb"
)

type Config struct {
	Enabled bool   `yaml:"enabled"`
	APIKey  string `yaml:"api_key"`

	Log        log.Logger            `yaml:"-"`
	Registerer prometheus.Registerer `yaml:"-"`

	RepoModel      pb.ModelServiceServer      `yaml:"-"`
	RepoGenerative pb.GenerativeServiceServer `yaml:"-"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "genai.enabled", false, "Enables GenerativeAi Service Server")
	fs.StringVar(&cfg.APIKey, "genai.api-key", "api-key", "API key.")
}

func (cfg *Config) Validate() error {
	return nil
}

// The Server type implements a pb server.
type Server struct {
	pb.UnimplementedModelServiceServer
	pb.UnimplementedGenerativeServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repoModel      pb.ModelServiceServer
	repoGenerative pb.GenerativeServiceServer
}

func NewServer(cfg Config) (*Server, error) {
	srv := &Server{
		Cfg:        cfg,
		logger:     cfg.Log,
		Registerer: cfg.Registerer,

		repoModel:      cfg.RepoModel,
		repoGenerative: cfg.RepoGenerative,
	}

	return srv, nil
}
