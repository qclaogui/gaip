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
	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
)

type Project interface {
	projectpb.ProjectServiceServer
}

const (
	DriverMemory = "memory"
)

var supportedDatabaseDrivers = []string{DriverMemory}

// Config database connections config
// Here are each of the database connections for application.
type Config struct {
	Driver string `yaml:"driver"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	prefix := "project.database."

	fs.StringVar(&cfg.Driver, prefix+"driver", DriverMemory, fmt.Sprintf("Driver storage to use. Supported drivers are: %s.", strings.Join(supportedDatabaseDrivers, ", ")))

}

// Validate RepoCfg config.
func (cfg *Config) Validate() error {
	if cfg.Driver != "" && !slices.Contains(supportedDatabaseDrivers, cfg.Driver) {
		return fmt.Errorf("unsupported database driver: %s", cfg.Driver)
	}

	return nil
}

func NewProject(cfg Config) (Project, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database driver %s", cfg.Driver)
	case DriverMemory:
		return NewMemoryRepo()
	default:
		return nil, errors.Errorf("unsupported driver for database %s", cfg.Driver)
	}
}
