// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package routeguide

import (
	"context"
	"flag"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
	"github.com/qclaogui/gaip/pkg/service"
	"github.com/qclaogui/gaip/pkg/service/routeguide/repository"
)

type Config struct {
	//RepoCfg holds the configuration used for the repository.
	RepoCfg repository.Config `yaml:"database"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	//Register RepoCfg Config
	cfg.RepoCfg.RegisterFlags(fs)
}

func (cfg *Config) Validate() error {
	//Validate RepoCfg Config
	if err := cfg.RepoCfg.Validate(); err != nil {
		return err
	}
	return nil
}

// routeGuideServiceImpl routeGuideServiceImpl
type routeGuideServiceImpl struct {
	routeguidepb.UnimplementedRouteGuideServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo repository.Repository
}

func New(cfg Config, s *service.Server) error {
	srv := &routeGuideServiceImpl{
		Cfg:        cfg,
		logger:     s.Log,
		Registerer: s.Registerer,
	}
	if err := srv.setupRepo(); err != nil {
		return err
	}

	routeguidepb.RegisterRouteGuideServiceServer(s.GRPCServer, srv)
	return nil
}

func (srv *routeGuideServiceImpl) setupRepo() error {
	var err error
	if srv.repo, err = repository.NewRepository(srv.Cfg.RepoCfg); err != nil {
		return err
	}
	return nil
}

// GetFeature returns the feature at the given point.
func (srv *routeGuideServiceImpl) GetFeature(ctx context.Context, req *routeguidepb.GetFeatureRequest) (*routeguidepb.GetFeatureResponse, error) {
	return srv.repo.GetFeature(ctx, req)
}

func (srv *routeGuideServiceImpl) ListFeatures(req *routeguidepb.ListFeaturesRequest, stream routeguidepb.RouteGuideService_ListFeaturesServer) error {
	return srv.repo.ListFeatures(req, stream)
}

func (srv *routeGuideServiceImpl) RecordRoute(req routeguidepb.RouteGuideService_RecordRouteServer) error {
	return srv.repo.RecordRoute(req)
}
func (srv *routeGuideServiceImpl) RouteChat(req routeguidepb.RouteGuideService_RouteChatServer) error {
	return srv.repo.RouteChat(req)
}
