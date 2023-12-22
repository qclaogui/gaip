// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package todo

import (
	"context"
	"flag"
	"fmt"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	APIVersion = "v1"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	Repo repository.Todo `yaml:"-"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "todo.enabled", true, "Enables Todo Service Server")
}

func (cfg *Config) Validate() error {
	return nil
}

// The ServiceServer type implements a todopb service server.
type ServiceServer struct {
	todopb.UnimplementedToDoServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo repository.Todo
}

func New(cfg Config, srv *service.Server) (*ServiceServer, error) {
	s := &ServiceServer{
		Cfg:        cfg,
		logger:     srv.Log,
		Registerer: srv.Registerer,
		repo:       cfg.Repo,
	}

	todopb.RegisterToDoServiceServer(srv.GRPCServer, s)
	return s, nil
}

func (s *ServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if APIVersion != api {
			return fmt.Errorf("unsupported API version: service implements API version '%s', but asked for '%s'", APIVersion, api)
		}
	}
	return nil
}

func (s *ServiceServer) Create(ctx context.Context, req *todopb.CreateRequest) (*todopb.CreateResponse, error) {
	_ = level.Info(s.logger).Log("msg", "[Create] received request")
	defer func() { _ = level.Info(s.logger).Log("msg", "[Create] completed request") }()

	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if err := req.GetItem().GetCreatedAt().CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	return s.repo.Create(ctx, req)
}

func (s *ServiceServer) Update(ctx context.Context, req *todopb.UpdateRequest) (*todopb.UpdateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return s.repo.Update(ctx, req)
}

func (s *ServiceServer) Get(ctx context.Context, req *todopb.GetRequest) (*todopb.GetResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return s.repo.Get(ctx, req)
}

func (s *ServiceServer) Delete(ctx context.Context, req *todopb.DeleteRequest) (*todopb.DeleteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return s.repo.Delete(ctx, req)
}

func (s *ServiceServer) List(ctx context.Context, req *todopb.ListRequest) (*todopb.ListResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return s.repo.List(ctx, req)
}
