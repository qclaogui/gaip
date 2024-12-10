// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package mysql

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
	fs.StringVar(&cfg.URL, prefix+"mysql.url", "", "Use either URL or the other fields below to configure the database. Example: mysql://user:secret@host:port/database")
	fs.StringVar(&cfg.Host, prefix+"mysql.host", "127.0.0.1:3306", `IP or hostname and port or in case of Unix sockets the path to it.For example, for MySQL running on the same host: host = 127.0.0.1:3306 or with Unix sockets: host = /var/run/mysqld/mysqld.sock`)
	fs.StringVar(&cfg.User, prefix+"mysql.user", "root", "mysql user")
	fs.Var(&cfg.Password, prefix+"mysql.password", "password")
	fs.StringVar(&cfg.Schema, prefix+"mysql.schema", "database", "schema")
}

func (cfg *Config) Validate() error {
	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	// param := "parseTime=true"

	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Schema, param)
	// toDoSrv, err := todov1.NewServiceServer(todov1.WithMysqlRepository(dsn))

	return nil
}
