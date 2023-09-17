// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"flag"
	"sync"

	"github.com/qclaogui/golang-api-server/genproto/library/apiv1/librarypb"
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

	// shelves are stored in a map keyed by shelf id
	// books are stored in a two level map, keyed first by shelf id and then by book id
	Shelves     map[int64]*librarypb.Shelf
	Books       map[int64]map[int64]*librarypb.Book
	LastShelfID int64      // the id of the last shelf that was added
	LastBookID  int64      // the id of the last book that was added
	Mutex       sync.Mutex // global mutex to synchronize service access
}

// NewMemoryRepo is a factory function to generate a new repository
func NewMemoryRepo() (*MemoryRepo, error) {
	mr := &MemoryRepo{
		Shelves: map[int64]*librarypb.Shelf{},
		Books:   map[int64]map[int64]*librarypb.Book{},
	}
	return mr, nil

}
