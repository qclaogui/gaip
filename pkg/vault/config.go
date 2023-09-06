// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package vault

import (
	"errors"
	"flag"
)

type Config struct {
	Enabled bool `yaml:"enabled"`

	URL       string `yaml:"url"`
	Token     string `yaml:"token"`
	MountPath string `yaml:"mount_path"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, "vault.enabled", false, "Enables fetching of keys and certificates from Vault")
	fs.StringVar(&cfg.URL, "vault.url", "", "Location of the Vault server")
	fs.StringVar(&cfg.Token, "vault.token", "", "Token used to authenticate with Vault")
	fs.StringVar(&cfg.MountPath, "vault.mount-path", "", "Location of secrets engine within Vault")

}

func (cfg *Config) Validate() error {
	if !cfg.Enabled {
		return nil
	}

	if cfg.URL == "" {
		return errors.New("empty vault URL supplied")
	}

	if cfg.MountPath == "" {
		return errors.New("empty vault mount path supplied")
	}

	return nil
}
