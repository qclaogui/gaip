// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	"github.com/grafana/dskit/flagext"
	pb "github.com/qclaogui/golang-api-server/api/bookstore/v1alpha1/bookstorepb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MysqlRepo is used to implement BookstoreServiceServer.
type MysqlRepo struct {
	pb.UnimplementedBookstoreServiceServer
	db *sql.DB
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
	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	//param := "parseTime=true"

	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Schema, param)
	//toDoSrv, err := todov1.NewServiceServer(todov1.WithMysqlRepository(dsn))

	return nil
}

// NewMysqlRepo is a factory function to generate a new repository
func NewMysqlRepo(db *sql.DB) (*MysqlRepo, error) {
	repo := &MysqlRepo{db: db}
	return repo, nil
}

func (r *MysqlRepo) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := r.db.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database-> " + err.Error())
	}
	return c, nil
}

func (r *MysqlRepo) GetShelf(ctx context.Context, req *pb.GetShelfRequest) (*pb.Shelf, error) {
	// get SQL connection from pool
	c, err := r.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Close() }()

	sid := req.Shelf

	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Theme` FROM Shelf WHERE `ID`=?", sid)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Shelf-> "+err.Error())
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to select from Shelf-> "+err.Error())
		}
	}

	shelf := &pb.Shelf{}
	if err = rows.Scan(&shelf.Id, &shelf.Theme); err != nil {
		return nil, fmt.Errorf("failed to retrieve field values from Shelf row-> " + err.Error())
	}

	return shelf, nil
}
