package routeguide

import (
	"context"
	pb "github.com/qclaogui/golang-api-server/pkg/api/routeguidepb"
)

type Repository interface {
	GetFeature(context.Context, *pb.Point) (*pb.Feature, error)

	ListFeatures(*pb.Rectangle, pb.RouteGuide_ListFeaturesServer) error
}
