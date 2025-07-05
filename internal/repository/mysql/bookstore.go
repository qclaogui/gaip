// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
	"github.com/qclaogui/gaip/internal/ent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// bookstoreImpl fulfills the bookstoreImpl Repository interface
// All data are managed by MysqlCfg.
//
// bookstoreImpl is used to implement BookstoreServiceServer.
type bookstoreImpl struct {
	bookstorepb.UnimplementedBookstoreServiceServer

	sqlDB     *sql.DB
	entClient *ent.Client
}

// NewBookstore is a factory function to generate a new repository
func NewBookstore(cfg Config) (bookstorepb.BookstoreServiceServer, error) {
	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Schema)
	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}
	repo := &bookstoreImpl{entClient: client}
	return repo, nil
}

// NewBookstoreWithSQLDB is a factory function to generate a new repository
func NewBookstoreWithSQLDB(db *sql.DB) (bookstorepb.BookstoreServiceServer, error) {
	repo := &bookstoreImpl{sqlDB: db}
	return repo, nil
}

func (r *bookstoreImpl) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := r.sqlDB.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database-> %w", err)
	}
	return c, nil
}

func (r *bookstoreImpl) GetShelf(ctx context.Context, req *bookstorepb.GetShelfRequest) (*bookstorepb.Shelf, error) {
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

	shelf := &bookstorepb.Shelf{}
	if err = rows.Scan(&shelf.Id, &shelf.Theme); err != nil {
		return nil, fmt.Errorf("failed to retrieve field values from Shelf row-> %w", err)
	}

	return shelf, nil
}
