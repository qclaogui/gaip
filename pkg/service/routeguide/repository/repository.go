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

type Repository interface {
	routeguidepb.RouteGuideServiceServer
}

const (
	DriverMemory = "memory"
)

var supportedDatabaseDrivers = []string{DriverMemory}

// Config RepoCfg Connections config
// Here are each of the database connections for application.
type Config struct {
	Driver string `yaml:"driver"`

	Memory MemoryConfig `yaml:"memory"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	prefix := "routeguide.database."

	fs.StringVar(&cfg.Driver, prefix+"driver", DriverMemory, fmt.Sprintf("Driver storage to use. Supported drivers are: %s.", strings.Join(supportedDatabaseDrivers, ", ")))

	cfg.Memory.RegisterFlagsWithPrefix(prefix, fs)
}

// Validate RepoCfg config.
func (cfg *Config) Validate() error {
	if cfg.Driver != "" && !slices.Contains(supportedDatabaseDrivers, cfg.Driver) {
		return fmt.Errorf("unsupported RepoCfg driver: %s", cfg.Driver)
	}

	switch cfg.Driver {
	case DriverMemory:
		return cfg.Memory.Validate()
	}
	return nil
}

func NewRepository(cfg Config) (Repository, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database driver %s", cfg.Driver)
	case DriverMemory:
		return NewMemoryRepo(cfg.Memory)
	default:
		return nil, errors.Errorf("unsupported driver for database %s", cfg.Driver)
	}
}
