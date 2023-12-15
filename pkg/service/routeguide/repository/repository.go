// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/pkg/errors"
	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
)

const (
	BackendMemory = "memory"
)

var supportedDatabaseBackends = []string{BackendMemory}

type Repository interface {
	routeguidepb.RouteGuideServiceServer
}

// Config RepoCfg Connections config
// Here are each of the database connections for application.
type Config struct {
	Backend string `yaml:"backend"`

	Memory MemoryConfig `yaml:"memory"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	prefix := "routeguide.database."

	fs.StringVar(&cfg.Backend, prefix+"backend", BackendMemory, fmt.Sprintf("Backend storage to use. Supported backends are: %s.", strings.Join(supportedDatabaseBackends, ", ")))

	cfg.Memory.RegisterFlagsWithPrefix(prefix, fs)
}

// Validate RepoCfg config.
func (cfg *Config) Validate() error {
	if cfg.Backend != "" && !slices.Contains(supportedDatabaseBackends, cfg.Backend) {
		return fmt.Errorf("unsupported RepoCfg backend: %s", cfg.Backend)
	}

	switch cfg.Backend {
	case BackendMemory:
		return cfg.Memory.Validate()
	}
	return nil
}

func NewRepository(cfg Config) (Repository, error) {
	switch cfg.Backend {
	case "":
		return nil, errors.Errorf("empty database backend %s", cfg.Backend)
	case BackendMemory:
		return NewMemoryRepo(cfg.Memory)
	default:
		return nil, errors.Errorf("unsupported backend for database %s", cfg.Backend)
	}
}
