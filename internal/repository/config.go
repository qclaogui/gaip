// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/qclaogui/gaip/internal/repository/memory"
	"github.com/qclaogui/gaip/internal/repository/mysql"
	"github.com/qclaogui/gaip/internal/repository/postgres"
)

const (
	DriverMemory   = "memory"
	DriverMysql    = "mysql"
	DriverPostgres = "postgres"
)

var (
	supportedDatabaseBackends = []string{DriverMemory, DriverMysql, DriverPostgres}
)

type Config struct {
	Driver string `yaml:"driver"`

	MemoryCfg   memory.Config   `yaml:"memory"`
	MysqlCfg    mysql.Config    `yaml:"mysql"`
	PostgresCfg postgres.Config `yaml:"postgres"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	prefix := "database."
	fs.StringVar(&cfg.Driver, prefix+"driver", DriverMemory, fmt.Sprintf("Driver storage to use. Supported drivers are: %s.", strings.Join(supportedDatabaseBackends, ", ")))

	cfg.MemoryCfg.RegisterFlagsWithPrefix(prefix, fs)
	cfg.MysqlCfg.RegisterFlagsWithPrefix(prefix, fs)
	cfg.PostgresCfg.RegisterFlagsWithPrefix(prefix, fs)
}

func (cfg *Config) Validate() error {
	if cfg.Driver != "" && !slices.Contains(supportedDatabaseBackends, cfg.Driver) {
		return fmt.Errorf("unsupported database driver: %s", cfg.Driver)
	}

	return nil
}
