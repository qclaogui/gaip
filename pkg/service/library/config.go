// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package library

import (
	"flag"

	"github.com/qclaogui/golang-api-server/pkg/service/library/repository"
)

type Config struct {
	//RepoCfg holds the configuration used for the repository.
	RepoCfg repository.Config `yaml:"database"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	//Register RepoCfg Config
	cfg.RepoCfg.RegisterFlags(fs)
}

func (cfg *Config) Validate() error {
	//Validate RepoCfg Config
	if err := cfg.RepoCfg.Validate(); err != nil {
		return err
	}
	return nil
}
