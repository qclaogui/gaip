// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package routeguide

import (
	"context"

	pb "github.com/qclaogui/golang-api-server/api/gen/proto/routeguide/v1"
)

type Repository interface {
	GetFeature(context.Context, *pb.GetFeatureRequest) (*pb.GetFeatureResponse, error)

	ListFeatures(*pb.ListFeaturesRequest, pb.RouteGuideService_ListFeaturesServer) error

	RecordRoute(pb.RouteGuideService_RecordRouteServer) error

	RouteChat(pb.RouteGuideService_RouteChatServer) error
}
