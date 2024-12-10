// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package routeguide

import (
	"context"

	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
)

// GetFeature returns the feature at the given point.
func (s *Server) GetFeature(ctx context.Context, req *routeguidepb.GetFeatureRequest) (*routeguidepb.GetFeatureResponse, error) {
	return s.repo.GetFeature(ctx, req)
}

func (s *Server) ListFeatures(req *routeguidepb.ListFeaturesRequest, stream routeguidepb.RouteGuideService_ListFeaturesServer) error {
	return s.repo.ListFeatures(req, stream)
}

func (s *Server) RecordRoute(req routeguidepb.RouteGuideService_RecordRouteServer) error {
	return s.repo.RecordRoute(req)
}

func (s *Server) RouteChat(req routeguidepb.RouteGuideService_RouteChatServer) error {
	return s.repo.RouteChat(req)
}
