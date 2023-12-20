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
	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
)

const (
	DriverMemory   = "memory"
	DriverMysql    = "mysql"
	DriverPostgres = "postgres"
)

var supportedDatabaseDrivers = []string{DriverMemory, DriverMysql}

type Repository interface {
	librarypb.LibraryServiceServer
}

// Config RepoCfg Connections config
// Here are each of the database connections for application.
type Config struct {
	Driver string `yaml:"driver"`

	Memory   MemoryConfig   `yaml:"memory"`
	Mysql    MysqlConfig    `yaml:"mysql"`
	Postgres PostgresConfig `yaml:"postgres"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	prefix := "library.database."

	fs.StringVar(&cfg.Driver, prefix+"driver", DriverMemory, fmt.Sprintf("Driver storage to use. Supported drivers are: %s.", strings.Join(supportedDatabaseDrivers, ", ")))

	cfg.Memory.RegisterFlagsWithPrefix(prefix, fs)
	cfg.Mysql.RegisterFlagsWithPrefix(prefix, fs)
}

// Validate RepoCfg config.
func (cfg *Config) Validate() error {
	if cfg.Driver != "" && !slices.Contains(supportedDatabaseDrivers, cfg.Driver) {
		return fmt.Errorf("unsupported RepoCfg driver: %s", cfg.Driver)
	}

	switch cfg.Driver {
	case DriverMemory:
		return cfg.Memory.Validate()
	case DriverMysql:
		return cfg.Mysql.Validate()
	case DriverPostgres:
		return cfg.Postgres.Validate()
	}
	return nil
}

func NewRepository(cfg Config) (Repository, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database driver %s", cfg.Driver)
	case DriverMemory:
		return NewMemoryRepo(cfg.Memory)
	case DriverMysql:
		return NewMysqlRepo(cfg.Mysql)
	case DriverPostgres:
		return NewPostgresRepo(cfg.Postgres)
	default:
		return nil, errors.Errorf("unsupported driver for database %s", cfg.Driver)
	}
}
