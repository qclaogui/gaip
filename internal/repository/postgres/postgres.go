// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	// Register stdlib is the compatibility layer from pgx to database/sql.
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/gaip/internal/ent"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Library fulfills the Library Repository interface
// All data are managed by Postgres.
//
// Library is used to implement LibraryServiceServer.
type Library struct {
	entClient *ent.Client
}

func NewLibrary(cfg Config) (librarypb.LibraryServiceServer, error) {
	//"postgres://pgx_md5:secret@localhost:5432/pgx_test?sslmode=disable"
	dsn := fmt.Sprintf("postgres://%s:%s@tcp(%s)/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Schema)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}

	// Create an ent.Driver from `entClient`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	repo := &Library{entClient: ent.NewClient(ent.Driver(drv))}
	return repo, nil
}

func (r *Library) CreateShelf(_ context.Context, _ *librarypb.CreateShelfRequest) (*librarypb.Shelf, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Library) GetShelf(_ context.Context, _ *librarypb.GetShelfRequest) (*librarypb.Shelf, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Library) ListShelves(_ context.Context, _ *librarypb.ListShelvesRequest) (*librarypb.ListShelvesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Library) DeleteShelf(_ context.Context, _ *librarypb.DeleteShelfRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Library) MergeShelves(_ context.Context, _ *librarypb.MergeShelvesRequest) (*librarypb.Shelf, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Library) CreateBook(_ context.Context, _ *librarypb.CreateBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Library) GetBook(_ context.Context, _ *librarypb.GetBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Library) ListBooks(_ context.Context, _ *librarypb.ListBooksRequest) (*librarypb.ListBooksResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Library) DeleteBook(_ context.Context, _ *librarypb.DeleteBookRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Library) UpdateBook(_ context.Context, _ *librarypb.UpdateBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Library) MoveBook(_ context.Context, _ *librarypb.MoveBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}
