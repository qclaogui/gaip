// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package mysql

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/gaip/internal/ent"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Library fulfills the Library Repository interface
// All data are managed by MysqlCfg.
//
// Library is used to implement LibraryServiceServer.
type Library struct {
	entClient *ent.Client
}

// NewLibrary is a factory function to generate a new repository
func NewLibrary(cfg Config) (*Library, error) {
	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Schema, param)

	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}
	repo := &Library{entClient: client}
	return repo, nil
}

func (l *Library) CreateShelf(_ context.Context, _ *librarypb.CreateShelfRequest) (*librarypb.Shelf, error) {
	//TODO implement me
	panic("implement me")
}

func (l *Library) GetShelf(_ context.Context, _ *librarypb.GetShelfRequest) (*librarypb.Shelf, error) {
	//TODO implement me
	panic("implement me")
}

func (l *Library) ListShelves(_ context.Context, _ *librarypb.ListShelvesRequest) (*librarypb.ListShelvesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l *Library) DeleteShelf(_ context.Context, _ *librarypb.DeleteShelfRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (l *Library) MergeShelves(_ context.Context, _ *librarypb.MergeShelvesRequest) (*librarypb.Shelf, error) {
	//TODO implement me
	panic("implement me")
}

func (l *Library) CreateBook(_ context.Context, _ *librarypb.CreateBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (l *Library) GetBook(_ context.Context, _ *librarypb.GetBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (l *Library) ListBooks(_ context.Context, _ *librarypb.ListBooksRequest) (*librarypb.ListBooksResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l *Library) DeleteBook(_ context.Context, _ *librarypb.DeleteBookRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (l *Library) UpdateBook(_ context.Context, _ *librarypb.UpdateBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (l *Library) MoveBook(_ context.Context, _ *librarypb.MoveBookRequest) (*librarypb.Book, error) {
	//TODO implement me
	panic("implement me")
}
