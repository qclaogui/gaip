// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"flag"

	"github.com/qclaogui/gaip/pkg/service/project/repository"
)

type Config struct {
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
