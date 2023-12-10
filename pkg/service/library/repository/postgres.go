// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/grafana/dskit/flagext"

	// Register stdlib is the compatibility layer from pgx to database/sql.
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/gaip/internal/ent"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostgresRepo struct {
	db *ent.Client
}

type PostgresConfig struct {
	Enabled bool `yaml:"enabled"`

	URL      string         `yaml:"url"`
	Host     string         `yaml:"host"`
	User     string         `yaml:"user"`
	Password flagext.Secret `yaml:"password"`
	Schema   string         `yaml:"schema"`
}

func (cfg *PostgresConfig) RegisterFlags(fs *flag.FlagSet) {
	cfg.RegisterFlagsWithPrefix("", fs)
}

func (cfg *PostgresConfig) RegisterFlagsWithPrefix(prefix string, fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, prefix+"postgres.enabled", false, "Enables Mysql for backend database")

	fs.StringVar(&cfg.URL, prefix+"postgres.url", "", "Use either URL or the other fields below to configure the database. Example: mysql://user:secret@host:port/database")
	fs.StringVar(&cfg.Host, prefix+"postgres.host", "127.0.0.1:5432", `IP or hostname and port or in case of Unix sockets the path to it.For example, for MySQL running on the same host: host = 127.0.0.1:3306 or with Unix sockets: host = /var/run/mysqld/mysqld.sock`)
	fs.StringVar(&cfg.User, prefix+"postgres.user", "root", "RepoCfg user")
	fs.Var(&cfg.Password, prefix+"postgres.password", "RepoCfg password")
	fs.StringVar(&cfg.Schema, prefix+"postgres.schema", "database", "RepoCfg schema")
}

func (cfg *PostgresConfig) Validate() error {
	return nil
}

func NewPostgresRepo(cfg PostgresConfig) (*PostgresRepo, error) {
	//"postgres://pgx_md5:secret@localhost:5432/pgx_test?sslmode=disable"
	dsn := fmt.Sprintf("postgres://%s:%s@tcp(%s)/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Schema)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	repo := &PostgresRepo{db: ent.NewClient(ent.Driver(drv))}
	return repo, nil
}

func (r *PostgresRepo) CreateShelf(_ context.Context, _ *librarypb.CreateShelfRequest) (*librarypb.Shelf, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepo) GetShelf(_ context.Context, _ *librarypb.GetShelfRequest) (*librarypb.Shelf, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepo) ListShelves(_ context.Context, _ *librarypb.ListShelvesRequest) (*librarypb.ListShelvesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepo) DeleteShelf(_ context.Context, _ *librarypb.DeleteShelfRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepo) MergeShelves(_ context.Context, _ *librarypb.MergeShelvesRequest) (*librarypb.Shelf, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepo) CreateBook(_ context.Context, _ *librarypb.CreateBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepo) GetBook(_ context.Context, _ *librarypb.GetBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepo) ListBooks(_ context.Context, _ *librarypb.ListBooksRequest) (*librarypb.ListBooksResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepo) DeleteBook(_ context.Context, _ *librarypb.DeleteBookRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepo) UpdateBook(_ context.Context, _ *librarypb.UpdateBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepo) MoveBook(_ context.Context, _ *librarypb.MoveBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}
