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
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
)

const (
	DriverMemory = "memory"
	DriverMysql  = "mysql"
)

var (
	supportedDatabaseBackends = []string{DriverMemory, DriverMysql}

	// ErrNotFound is returned when a item is not found.
	ErrNotFound = errors.New("the item was not found in the repository")

	// ErrFailedToCreate is returned when a item is create Failed
	ErrFailedToCreate = errors.New("failed to add the todo to the repository")
)

type Repository interface {
	todopb.ToDoServiceServer
}

type Config struct {
	Driver string `yaml:"driver"`

	Memory MemoryConfig `yaml:"memory"`
	Mysql  MysqlConfig  `yaml:"mysql"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	prefix := "todo.database."
	fs.StringVar(&cfg.Driver, prefix+"driver", DriverMemory, fmt.Sprintf("Driver storage to use. Supported drivers are: %s.", strings.Join(supportedDatabaseBackends, ", ")))

	cfg.Memory.RegisterFlagsWithPrefix(prefix, fs)
	cfg.Mysql.RegisterFlagsWithPrefix(prefix, fs)
}

func (cfg *Config) Validate() error {
	if cfg.Driver != "" && !slices.Contains(supportedDatabaseBackends, cfg.Driver) {
		return fmt.Errorf("unsupported RepoCfg driver: %s", cfg.Driver)
	}

	return nil
}

func NewRepository(cfg Config) (Repository, error) {
	switch cfg.Driver {
	case "":
		return nil, errors.Errorf("empty database driver %s", cfg.Driver)
	case DriverMemory:
		return NewMemoryRepo(), nil
	case DriverMysql:
		return NewMysqlRepo(cfg.Mysql)
	default:
		return nil, errors.Errorf("unsupported driver for database %s", cfg.Driver)
	}
}
