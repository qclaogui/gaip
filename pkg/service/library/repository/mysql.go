// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"context"
	"flag"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/grafana/dskit/flagext"
	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/gaip/internal/ent"
	"google.golang.org/protobuf/types/known/emptypb"
)

// MysqlRepo is used to implement LibraryServiceServer.
type MysqlRepo struct {
	db *ent.Client
}

type MysqlConfig struct {
	Enabled bool `yaml:"enabled"`

	URL      string         `yaml:"url"`
	Host     string         `yaml:"host"`
	User     string         `yaml:"user"`
	Password flagext.Secret `yaml:"password"`
	Schema   string         `yaml:"schema"`
}

func (cfg *MysqlConfig) RegisterFlags(fs *flag.FlagSet) {
	cfg.RegisterFlagsWithPrefix("", fs)
}

func (cfg *MysqlConfig) RegisterFlagsWithPrefix(prefix string, fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, prefix+"mysql.enabled", false, "Enables Mysql for backend database")

	fs.StringVar(&cfg.URL, prefix+"mysql.url", "", "Use either URL or the other fields below to configure the database. Example: mysql://user:secret@host:port/database")
	fs.StringVar(&cfg.Host, prefix+"mysql.host", "127.0.0.1:3306", `IP or hostname and port or in case of Unix sockets the path to it.For example, for MySQL running on the same host: host = 127.0.0.1:3306 or with Unix sockets: host = /var/run/mysqld/mysqld.sock`)
	fs.StringVar(&cfg.User, prefix+"mysql.user", "root", "RepoCfg user")
	fs.Var(&cfg.Password, prefix+"mysql.password", "RepoCfg password")
	fs.StringVar(&cfg.Schema, prefix+"mysql.schema", "database", "RepoCfg schema")
}

func (cfg *MysqlConfig) Validate() error {
	return nil
}

// NewMysqlRepo is a factory function to generate a new repository
func NewMysqlRepo(cfg MysqlConfig) (*MysqlRepo, error) {
	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Schema, param)

	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}
	repo := &MysqlRepo{db: client}
	return repo, nil
}

func (r *MysqlRepo) CreateShelf(_ context.Context, _ *librarypb.CreateShelfRequest) (*librarypb.Shelf, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MysqlRepo) GetShelf(_ context.Context, _ *librarypb.GetShelfRequest) (*librarypb.Shelf, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MysqlRepo) ListShelves(_ context.Context, _ *librarypb.ListShelvesRequest) (*librarypb.ListShelvesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MysqlRepo) DeleteShelf(_ context.Context, _ *librarypb.DeleteShelfRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MysqlRepo) MergeShelves(_ context.Context, _ *librarypb.MergeShelvesRequest) (*librarypb.Shelf, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MysqlRepo) CreateBook(_ context.Context, _ *librarypb.CreateBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MysqlRepo) GetBook(_ context.Context, _ *librarypb.GetBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MysqlRepo) ListBooks(_ context.Context, _ *librarypb.ListBooksRequest) (*librarypb.ListBooksResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MysqlRepo) DeleteBook(_ context.Context, _ *librarypb.DeleteBookRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MysqlRepo) UpdateBook(_ context.Context, _ *librarypb.UpdateBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MysqlRepo) MoveBook(_ context.Context, _ *librarypb.MoveBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}
