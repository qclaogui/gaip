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
	"github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
)

const (
	DriverMemory = "memory"
	DriverMysql  = "mysql"
)

var supportedDatabaseDrivers = []string{DriverMemory, DriverMysql}

type Repository interface {
	bookstorepb.BookstoreServiceServer
}

// Config RepoCfg Connections config
// Here are each of the database connections for application.
type Config struct {
	Driver string `yaml:"driver"`

	Memory MemoryConfig `yaml:"memory"`
	Mysql  MysqlConfig  `yaml:"mysql"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	prefix := "bookstore.database."

	fs.StringVar(&cfg.Driver, prefix+"driver", DriverMemory, fmt.Sprintf("Driver storage to use. Supported drivers are: %s.", strings.Join(supportedDatabaseDrivers, ", ")))

	cfg.Memory.RegisterFlagsWithPrefix(prefix, fs)
	cfg.Mysql.RegisterFlagsWithPrefix(prefix, fs)
}

// Validate RepoCfg config.
func (cfg *Config) Validate() error {
	if cfg.Driver != "" && !slices.Contains(supportedDatabaseDrivers, cfg.Driver) {
		return fmt.Errorf("unsupported drivers: %s", cfg.Driver)
	}

	switch cfg.Driver {
	case DriverMemory:
		return cfg.Memory.Validate()
	case DriverMysql:
		return cfg.Mysql.Validate()
	}
	return nil
}

func NewRepository(cfg Config) (Repository, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database drivers %s", cfg.Driver)
	case DriverMemory:
		return NewMemoryRepo()
	case DriverMysql:
		return nil, nil //TODO(qc)
	default:
		return nil, errors.Errorf("unsupported drivers for database %s", cfg.Driver)
	}
}
