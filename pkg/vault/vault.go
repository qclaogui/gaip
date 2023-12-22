// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package vault

import (
	"context"
	"errors"
	"flag"
	"fmt"

	hashivault "github.com/hashicorp/vault/api"
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

// SecretsEngine  Vault Secrets Engine
type SecretsEngine interface {
	Get(ctx context.Context, path string) (*hashivault.KVSecret, error)
}

type Vault struct {
	KVStore SecretsEngine
}

func New(cfg Config) (*Vault, error) {
	config := hashivault.DefaultConfig()
	config.Address = cfg.URL

	client, err := hashivault.NewClient(config)
	if err != nil {
		return nil, err
	}
	client.SetToken(cfg.Token)

	vault := &Vault{
		KVStore: client.KVv2(cfg.MountPath),
	}

	return vault, nil
}

func (v *Vault) ReadSecret(path string) ([]byte, error) {
	secret, err := v.KVStore.Get(context.Background(), path)
	if err != nil {
		return nil, fmt.Errorf("unable to read secret from vault: %v", err)
	}

	if secret == nil || secret.Data == nil {
		return nil, errors.New("secret data is nil")
	}

	data, ok := secret.Data["value"].(string)
	if !ok {
		return nil, fmt.Errorf("secret data type is not string, found %T value: %#v", secret.Data["value"], secret.Data["value"])
	}
	return []byte(data), nil
}
