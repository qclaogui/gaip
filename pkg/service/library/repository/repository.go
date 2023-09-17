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
	"github.com/qclaogui/golang-api-server/genproto/library/apiv1/librarypb"
)

const (
	BackendMemory = "memory"
	BackendMysql  = "mysql"
)

var supportedDatabaseBackends = []string{BackendMemory, BackendMysql}

type Repository interface {
	librarypb.LibraryServiceServer
}

// Config RepoCfg Connections config
// Here are each of the database connections for application.
type Config struct {
	Backend string `yaml:"backend"`

	Memory MemoryConfig `yaml:"memory"`
	//Mysql MysqlConfig `yaml:"mysql"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	prefix := "library.database."

	fs.StringVar(&cfg.Backend, prefix+"backend", BackendMemory, fmt.Sprintf("Backend storage to use. Supported backends are: %s.", strings.Join(supportedDatabaseBackends, ", ")))

	cfg.Memory.RegisterFlagsWithPrefix(prefix, fs)
	//cfg.Mysql.RegisterFlagsWithPrefix(prefix, fs)
}

// Validate RepoCfg config.
func (cfg *Config) Validate() error {
	if cfg.Backend != "" && !slices.Contains(supportedDatabaseBackends, cfg.Backend) {
		return fmt.Errorf("unsupported RepoCfg backend: %s", cfg.Backend)
	}

	switch cfg.Backend {
	case BackendMemory:
		return cfg.Memory.Validate()
		//case BackendMysql:
		//	return cfg.Mysql.Validate()
	}
	return nil
}

func NewRepository(cfg Config) (Repository, error) {
	switch cfg.Backend {
	case "":
		return nil, errors.Errorf("empty database backend %s", cfg.Backend)
	case BackendMemory:
		return NewMemoryRepo()
	case BackendMysql:
		return nil, nil //TODO(qc)
	default:
		return nil, errors.Errorf("unsupported backend for database %s", cfg.Backend)
	}
}
