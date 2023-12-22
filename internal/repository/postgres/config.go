// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package postgres

import (
	"flag"

	"github.com/grafana/dskit/flagext"
)

type Config struct {
	URL      string         `yaml:"url"`
	Host     string         `yaml:"host"`
	User     string         `yaml:"user"`
	Password flagext.Secret `yaml:"password"`
	Schema   string         `yaml:"schema"`
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	cfg.RegisterFlagsWithPrefix("", fs)
}

func (cfg *Config) RegisterFlagsWithPrefix(prefix string, fs *flag.FlagSet) {
	fs.StringVar(&cfg.URL, prefix+"postgres.url", "", "Use either URL or the other fields below to configure the database. Example: postgres://pgx_md5:secret@localhost:5432/pgx_test?sslmode=disable")
	fs.StringVar(&cfg.Host, prefix+"postgres.host", "127.0.0.1:5432", `IP or hostname and port or in case of Unix sockets the path to it.For example, for MySQL running on the same host: host = 127.0.0.1:3306 or with Unix sockets: host = /var/run/mysqld/mysqld.sock`)
	fs.StringVar(&cfg.User, prefix+"postgres.user", "root", "user")
	fs.Var(&cfg.Password, prefix+"postgres.password", "password")
	fs.StringVar(&cfg.Schema, prefix+"postgres.schema", "database", "schema")
}

func (cfg *Config) Validate() error {
	return nil
}
