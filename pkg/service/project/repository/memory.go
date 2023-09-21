// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"context"
	"flag"
	"sync"

	"github.com/qclaogui/golang-api-server/genproto/project/apiv1/projectpb"
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
// MemoryRepo is used to implement ProjectServiceServer.
type MemoryRepo struct {
	projectpb.UnimplementedProjectServiceServer
	mu sync.Mutex // global mutex to synchronize service access

	projects map[string]*projectpb.Project
}

// NewMemoryRepo is a factory function to generate a new repository
func NewMemoryRepo() (*MemoryRepo, error) {
	mr := &MemoryRepo{
		projects: map[string]*projectpb.Project{},
	}
	return mr, nil
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

func (mr *MemoryRepo) CreateProject(ctx context.Context, req *projectpb.CreateProjectRequest) (*projectpb.Project, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	_ = ctx
	_ = req

	_ = maxBatchSize
	_, _ = validatePageSize(1)
	//TODO implement me
	panic("implement me")
}
