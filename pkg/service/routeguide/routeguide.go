// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package routeguide

import (
	"context"

	"github.com/go-kit/log"
	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
	"google.golang.org/grpc"
)

type Option func(*ServiceServer) error

// WithRepository applies a given repository to the ServiceServer
func WithRepository(repo Repository) Option {
	return func(srv *ServiceServer) error {
		srv.repo = repo
		return nil
	}
}

// WithMemoryRepository applies a memory repository to the ServiceServer
func WithMemoryRepository() Option {
	return func(srv *ServiceServer) error {
		repo, err := NewMemoryRepository("")
		if err != nil {
			return err
		}
		srv.repo = repo
		return nil
	}
}

// ServiceServer ServiceServer
type ServiceServer struct {
	routeguidepb.UnimplementedRouteGuideServiceServer

	repo   Repository
	logger log.Logger
}

func NewServiceServer(logger log.Logger, opts ...Option) (*ServiceServer, error) {
	// Create the Server
	srv := &ServiceServer{logger: logger}
	// Apply all Configurations passed in
	for _, opt := range opts {
		// Pass the service into the configuration function
		if err := opt(srv); err != nil {
			return nil, err
		}
	}
	return srv, nil
}

func (srv *ServiceServer) RegisterGRPC(s *grpc.Server) {
	s.RegisterService(&routeguidepb.RouteGuideService_ServiceDesc, srv)
}

// GetFeature returns the feature at the given point.
func (srv *ServiceServer) GetFeature(ctx context.Context, req *routeguidepb.GetFeatureRequest) (*routeguidepb.GetFeatureResponse, error) {
	return srv.repo.GetFeature(ctx, req)
}

func (srv *ServiceServer) ListFeatures(req *routeguidepb.ListFeaturesRequest, stream routeguidepb.RouteGuideService_ListFeaturesServer) error {
	return srv.repo.ListFeatures(req, stream)
}

func (srv *ServiceServer) RecordRoute(req routeguidepb.RouteGuideService_RecordRouteServer) error {
	return srv.repo.RecordRoute(req)
}
func (srv *ServiceServer) RouteChat(req routeguidepb.RouteGuideService_RouteChatServer) error {
	return srv.repo.RouteChat(req)
}
