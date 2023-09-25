// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"context"
	"flag"
	"sync"

	"github.com/qclaogui/golang-api-server/genproto/library/apiv1/librarypb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	defaultPageSize = 20
	maxPageSize     = 10000
	maxBatchSize    = 1000
)

type MemoryConfig struct {
	Enabled bool `yaml:"enabled"`
}

func (cfg *MemoryConfig) RegisterFlagsWithPrefix(prefix string, fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, prefix+"memory.enabled", false, "Enables memory Repository")
}

func (cfg *MemoryConfig) RegisterFlags(fs *flag.FlagSet) {
	cfg.RegisterFlagsWithPrefix("", fs)
}

func (cfg *MemoryConfig) Validate() error {
	return nil
}

// MemoryRepo fulfills the Repository interface
// All objects are managed in an in-memory non-persistent store.
//
// MemoryRepo is used to implement LibraryServiceServer.
type MemoryRepo struct {
	librarypb.UnimplementedLibraryServiceServer
	mu sync.Mutex // global mutex to synchronize service access

	// shelves are stored in a map keyed by shelf id
	// books are stored in a two level map, keyed first by shelf id and then by book id
	Shelves     map[int64]*librarypb.Shelf
	Books       map[int64]map[int64]*librarypb.Book
	LastShelfID int64 // the id of the last shelf that was added
	LastBookID  int64 // the id of the last book that was added
}

// NewMemoryRepo is a factory function to generate a new repository
func NewMemoryRepo(_ MemoryConfig) (*MemoryRepo, error) {
	mr := &MemoryRepo{
		Shelves: map[int64]*librarypb.Shelf{},
		Books:   map[int64]map[int64]*librarypb.Book{},
	}
	return mr, nil

}

func (mr *MemoryRepo) ListShelves(ctx context.Context, req *librarypb.ListShelvesRequest) (*librarypb.ListShelvesResponse, error) {

	mr.mu.Lock()
	defer mr.mu.Unlock()

	_ = ctx

	ps, err := validatePageSize(req.PageSize)
	if err != nil {
		return nil, err
	}
	_ = ps
	_ = maxBatchSize

	return nil, status.Errorf(codes.Unimplemented, "method ListShelves not implemented")
}

// validatePageSize returns the default page size if the specified page size is 0, otherwise it
// validates the specified page size.
func validatePageSize(ps int32) (int32, error) {
	switch {
	case ps == 0:
		return defaultPageSize, nil
	case ps > maxPageSize:
		return 0, status.Errorf(codes.InvalidArgument, "page size %d cannot be large than max page size %d", ps, maxPageSize)
	case ps < 0:
		return 0, status.Errorf(codes.InvalidArgument, "page size %d cannot be negative", ps)
	}

	return ps, nil
}
