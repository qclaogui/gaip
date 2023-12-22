// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package memory

import (
	"context"
	"sync"

	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Library fulfills the Repository Library interface
// All objects are managed in an in-memory non-persistent store.
//
// Library is used to implement LibraryServiceServer.
type Library struct {
	librarypb.UnimplementedLibraryServiceServer

	// shelves are stored in a map keyed by shelf id
	// books are stored in a two level map, keyed first by shelf id and then by book id
	Shelves     map[int64]*librarypb.Shelf
	Books       map[int64]map[int64]*librarypb.Book
	LastShelfID int64      // the id of the last shelf that was added
	LastBookID  int64      // the id of the last book that was added
	mu          sync.Mutex // global mutex to synchronize service access
}

// NewLibrary is a factory function to generate a new repository
func NewLibrary() (*Library, error) {
	m := &Library{
		Shelves: map[int64]*librarypb.Shelf{},
		Books:   map[int64]map[int64]*librarypb.Book{},
	}
	return m, nil

}

func (m *Library) ListShelves(ctx context.Context, req *librarypb.ListShelvesRequest) (*librarypb.ListShelvesResponse, error) {

	m.mu.Lock()
	defer m.mu.Unlock()

	_ = ctx

	ps, err := validatePageSize(req.PageSize)
	if err != nil {
		return nil, err
	}
	_ = ps
	_ = maxBatchSize

	return nil, status.Errorf(codes.Unimplemented, "method ListShelves not implemented")
}
