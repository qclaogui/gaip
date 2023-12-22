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
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	Repo repository.RouteGuide `yaml:"-"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "routeguide.enabled", true, "Enables RouteGuide Service Server")
}

func (cfg *Config) Validate() error {
	return nil
}

// The ServiceServer type implements a routeguidepb server.
type ServiceServer struct {
	routeguidepb.UnimplementedRouteGuideServiceServer

	Cfg        Config
	logger     log.Logger
	Registerer prometheus.Registerer

	repo repository.RouteGuide
}

func New(cfg Config, s *service.Server) (*ServiceServer, error) {
	srv := &ServiceServer{
		Cfg:        cfg,
		logger:     s.Log,
		Registerer: s.Registerer,
		repo:       cfg.Repo,
	}

	routeguidepb.RegisterRouteGuideServiceServer(s.GRPCServer, srv)
	return srv, nil
}

// GetFeature returns the feature at the given point.
func (s *ServiceServer) GetFeature(ctx context.Context, req *routeguidepb.GetFeatureRequest) (*routeguidepb.GetFeatureResponse, error) {
	return s.repo.GetFeature(ctx, req)
}

func (s *ServiceServer) ListFeatures(req *routeguidepb.ListFeaturesRequest, stream routeguidepb.RouteGuideService_ListFeaturesServer) error {
	return s.repo.ListFeatures(req, stream)
}

func (s *ServiceServer) RecordRoute(req routeguidepb.RouteGuideService_RecordRouteServer) error {
	return s.repo.RecordRoute(req)
}
func (s *ServiceServer) RouteChat(req routeguidepb.RouteGuideService_RouteChatServer) error {
	return s.repo.RouteChat(req)
}
