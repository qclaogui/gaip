// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package task

import (
	"flag"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	pb "github.com/qclaogui/gaip/genproto/task/apiv1/taskpb"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	Log        log.Logger            `yaml:"-"`
	Registerer prometheus.Registerer `yaml:"-"`

	RepoTasksWriter pb.TasksWriterServiceServer `yaml:"-"`
	RepoTasksReader pb.TasksReaderServiceServer `yaml:"-"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "task.enabled", true, "Enables Task Service Server")
}

func (cfg *Config) Validate() error {
	return nil
}

// The Server type implements a pb server.
type Server struct {
	pb.UnimplementedTasksWriterServiceServer
	pb.UnimplementedTasksReaderServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	metrics taskMetrics

	repoTaskWriter pb.TasksWriterServiceServer `yaml:"-"`
	repoTaskReader pb.TasksReaderServiceServer `yaml:"-"`
}

func NewServer(cfg Config) (*Server, error) {
	srv := &Server{
		Cfg:        cfg,
		logger:     cfg.Log,
		Registerer: cfg.Registerer,

		metrics: newTaskMetrics(cfg.Registerer),

		repoTaskWriter: cfg.RepoTasksWriter,
		repoTaskReader: cfg.RepoTasksReader,
	}
	return srv, nil
}
