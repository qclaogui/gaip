// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package library

import (
	"github.com/go-kit/log"
	"github.com/qclaogui/gaip/pkg/service/library/repository"
)

// Option is an alias for a function that will take in a pointer to
// an libraryServiceImpl and modify it
type Option func(*libraryServiceImpl) error

// WithRepository applies a given repository to the libraryServiceImpl
func WithRepository(repo repository.Repository) Option {
	return func(srv *libraryServiceImpl) error {
		srv.repo = repo
		return nil
	}
}

// WithMemoryRepository applies a memory repository to the Option
func WithMemoryRepository() Option {
	return func(srv *libraryServiceImpl) error {
		repo, err := repository.NewMemoryRepo(repository.MemoryConfig{
			Enabled: true,
		})
		if err != nil {
			return err
		}
		srv.repo = repo
		return nil
	}
}

// WithLogger applies a given repository to the libraryServiceImpl
func WithLogger(logger log.Logger) Option {
	return func(srv *libraryServiceImpl) error {
		srv.logger = logger
		return nil
	}
}

func NewLibraryServiceWithOptions(cfg Config, opts ...Option) (Service, error) {
	// Create the libraryServiceImpl
	s := &libraryServiceImpl{Cfg: cfg}
	// Apply all Configurations passed in
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}
