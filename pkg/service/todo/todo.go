// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package todo

import (
	"context"
	"flag"
	"fmt"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"github.com/qclaogui/gaip/pkg/service/todo/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	APIVersion = "v1"
)

// Service Library Service Server
type Service interface {
	todopb.ToDoServiceServer
}

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

type todoServiceImpl struct {
	todopb.UnimplementedToDoServiceServer
	repo repository.Repository

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer
}

func NewServiceServer(cfg Config, logger log.Logger, reg prometheus.Registerer) (Service, error) {
	// Create the todoServiceImpl
	s := &todoServiceImpl{
		Cfg:        cfg,
		logger:     logger,
		Registerer: reg,
	}

	if err := s.setupRepo(); err != nil {
		return nil, err
	}

	return s, nil
}

func (srv *todoServiceImpl) setupRepo() error {
	var err error
	if srv.repo, err = repository.NewRepository(srv.Cfg.RepoCfg); err != nil {
		return err
	}

	return nil
}

func (srv *todoServiceImpl) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if APIVersion != api {
			return fmt.Errorf("unsupported API version: service implements API version '%s', but asked for '%s'", APIVersion, api)
		}
	}
	return nil
}

func (srv *todoServiceImpl) Create(ctx context.Context, req *todopb.CreateRequest) (*todopb.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if err := req.GetItem().GetCreatedAt().CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	return srv.repo.Create(ctx, req)
}

func (srv *todoServiceImpl) Update(ctx context.Context, req *todopb.UpdateRequest) (*todopb.UpdateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return srv.repo.Update(ctx, req)
}

func (srv *todoServiceImpl) Get(ctx context.Context, req *todopb.GetRequest) (*todopb.GetResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return srv.repo.Get(ctx, req)
}

func (srv *todoServiceImpl) Delete(ctx context.Context, req *todopb.DeleteRequest) (*todopb.DeleteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return srv.repo.Delete(ctx, req)
}

func (srv *todoServiceImpl) List(ctx context.Context, req *todopb.ListRequest) (*todopb.ListResponse, error) {
	// check if the API version requested by client is supported by server
	if err := srv.checkAPI(req.GetApi()); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return srv.repo.List(ctx, req)
}
