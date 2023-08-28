package routeguide

import (
	"context"
	"errors"

	pb "github.com/qclaogui/golang-api-server/pkg/api/routeguidepb"
)

var (
	// ErrNotFound is returned when a item is not found.
	ErrNotFound = errors.New("the item was not found in the repository")
)

// func (UnimplementedRouteGuideServer) GetFeature(context.Context, *Point) (*Feature, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
// }

type Repository interface {
	GetFeature(context.Context, *pb.Point) (*pb.Feature, error)
}
