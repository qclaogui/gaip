// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package showcase

import (
	"flag"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	pb "github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	Log        log.Logger            `yaml:"-"`
	Registerer prometheus.Registerer `yaml:"-"`

	RepoIdentity  pb.IdentityServiceServer  `yaml:"-"`
	RepoMessaging pb.MessagingServiceServer `yaml:"-"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "showcase.enabled", true, "Enables Showcase Service Server")
}

func (cfg *Config) Validate() error {
	return nil
}

// The Server type implements a pb server.
type Server struct {
	pb.UnimplementedEchoServiceServer
	pb.UnimplementedIdentityServiceServer
	pb.UnimplementedMessagingServiceServer

	longrunningpb.UnimplementedOperationsServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repoIdentity  pb.IdentityServiceServer
	repoMessaging pb.MessagingServiceServer

	nowF func() time.Time
}

func NewServer(cfg Config) (*Server, error) {
	srv := &Server{
		Cfg:        cfg,
		logger:     cfg.Log,
		Registerer: cfg.Registerer,

		repoIdentity:  cfg.RepoIdentity,
		repoMessaging: cfg.RepoMessaging,

		nowF: time.Now,
	}
	return srv, nil
}
