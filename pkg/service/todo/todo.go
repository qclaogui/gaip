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
	"github.com/qclaogui/gaip/pkg/service"
	"github.com/qclaogui/gaip/pkg/service/todo/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	APIVersion = "v1"
)

type Config struct {
	RepoCfg repository.Config `yaml:"database"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	cfg.RepoCfg.RegisterFlags(fs)
}

func (cfg *Config) Validate() error {
	if err := cfg.RepoCfg.Validate(); err != nil {
		return err
	}
	return nil
}

// The Todo type implements a todopb server.
type Todo struct {
	todopb.UnimplementedToDoServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo repository.Repository
}

func New(cfg Config, s *service.Server) (*Todo, error) {
	srv := &Todo{
		Cfg:        cfg,
		logger:     s.Log,
		Registerer: s.Registerer,
	}

	if err := srv.setupRepo(); err != nil {
		return nil, err
	}

	todopb.RegisterToDoServiceServer(s.GRPCServer, srv)
	return srv, nil
}

func (srv *Todo) setupRepo() error {
	var err error
	if srv.repo, err = repository.NewRepository(srv.Cfg.RepoCfg); err != nil {
		return err
	}

	return nil
}

func (srv *Todo) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if APIVersion != api {
			return fmt.Errorf("unsupported API version: service implements API version '%s', but asked for '%s'", APIVersion, api)
		}
	}
	return nil
}

func (srv *Todo) Create(ctx context.Context, req *todopb.CreateRequest) (*todopb.CreateResponse, error) {
	_ = level.Info(srv.logger).Log("msg", "[Create] received request")
	defer func() { _ = level.Info(srv.logger).Log("msg", "[Create] completed request") }()

	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if err := req.GetItem().GetCreatedAt().CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	return srv.repo.Create(ctx, req)
}

func (srv *Todo) Update(ctx context.Context, req *todopb.UpdateRequest) (*todopb.UpdateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return srv.repo.Update(ctx, req)
}

func (srv *Todo) Get(ctx context.Context, req *todopb.GetRequest) (*todopb.GetResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return srv.repo.Get(ctx, req)
}

func (srv *Todo) Delete(ctx context.Context, req *todopb.DeleteRequest) (*todopb.DeleteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return srv.repo.Delete(ctx, req)
}

func (srv *Todo) List(ctx context.Context, req *todopb.ListRequest) (*todopb.ListResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return srv.repo.List(ctx, req)
}
