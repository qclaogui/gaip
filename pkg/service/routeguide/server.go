package routeguide

import (
	"context"
	"sync"

	pb "github.com/qclaogui/golang-api-server/pkg/api/routeguidepb"
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
	pb.UnimplementedRouteGuideServer
	repo Repository

	mu sync.Mutex // protects routeNotes
}

func NewServiceServer(opts ...Option) (*ServiceServer, error) {
	// Create the Server
	srv := &ServiceServer{}
	// Apply all Configurations passed in
	for _, opt := range opts {
		// Pass the service into the configuration function
		if err := opt(srv); err != nil {
			return nil, err
		}
	}
	return srv, nil
}

// GetFeature returns the feature at the given point.
func (srv *ServiceServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	return srv.repo.GetFeature(ctx, point)
}

func (srv *ServiceServer) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
	return srv.repo.ListFeatures(rect, stream)
}
