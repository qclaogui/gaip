package routeguide

import (
	"context"

	pb "github.com/qclaogui/golang-api-server/api/gen/proto/routeguide/v1"
)

type Repository interface {
	GetFeature(context.Context, *pb.GetFeatureRequest) (*pb.GetFeatureResponse, error)

	ListFeatures(*pb.ListFeaturesRequest, pb.RouteGuideService_ListFeaturesServer) error
}
