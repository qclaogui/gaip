// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package v1

import (
	"context"

	"github.com/go-kit/log"
	pb "github.com/qclaogui/golang-api-server/api/routeguide/v1/routeguidepb"
	routeguide "github.com/qclaogui/golang-api-server/pkg/service/routeguide"
	"google.golang.org/grpc"
)

type Option func(*ServiceServer) error

// WithRepository applies a given repository to the ServiceServer
func WithRepository(repo routeguide.Repository) Option {
	return func(srv *ServiceServer) error {
		srv.repo = repo
		return nil
	}
}

// WithMemoryRepository applies a memory repository to the ServiceServer
func WithMemoryRepository() Option {
	return func(srv *ServiceServer) error {
		repo, err := routeguide.NewMemoryRepository("")
		if err != nil {
			return err
		}
		srv.repo = repo
		return nil
	}
}

// ServiceServer ServiceServer
type ServiceServer struct {
	pb.UnimplementedRouteGuideServiceServer

	repo   routeguide.Repository
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
	s.RegisterService(&pb.RouteGuideService_ServiceDesc, srv)
}

// GetFeature returns the feature at the given point.
func (srv *ServiceServer) GetFeature(ctx context.Context, req *pb.GetFeatureRequest) (*pb.GetFeatureResponse, error) {
	return srv.repo.GetFeature(ctx, req)
}

func (srv *ServiceServer) ListFeatures(req *pb.ListFeaturesRequest, stream pb.RouteGuideService_ListFeaturesServer) error {
	return srv.repo.ListFeatures(req, stream)
}

func (srv *ServiceServer) RecordRoute(req pb.RouteGuideService_RecordRouteServer) error {
	return srv.repo.RecordRoute(req)
}
func (srv *ServiceServer) RouteChat(req pb.RouteGuideService_RouteChatServer) error {
	return srv.repo.RouteChat(req)
}
