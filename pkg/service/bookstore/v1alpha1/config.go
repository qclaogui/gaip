// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package v1alpha1

import (
	"flag"

	"github.com/grafana/dskit/cache"
	"github.com/qclaogui/golang-api-server/pkg/service/bookstore/repository"
)

type Config struct {
	//RepoCfg holds the configuration used for the repository.
	RepoCfg repository.Config `yaml:"database"`

	// CacheCfg holds the configuration used for the cache.
	CacheCfg cache.BackendConfig `yaml:"cache"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	//Register RepoCfg Config
	cfg.RepoCfg.RegisterFlags(fs)
}

func (cfg *Config) Validate() error {
	//Register RepoCfg Config
	if err := cfg.RepoCfg.Validate(); err != nil {
		return err
	}
	return nil
}
