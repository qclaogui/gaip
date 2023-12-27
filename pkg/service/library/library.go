// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package library

import (
	"context"

	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreateShelf(ctx context.Context, req *librarypb.CreateShelfRequest) (*librarypb.Shelf, error) {
	return s.repo.CreateShelf(ctx, req)
}

func (s *Server) ListShelves(ctx context.Context, req *librarypb.ListShelvesRequest) (*librarypb.ListShelvesResponse, error) {
	return s.repo.ListShelves(ctx, req)
}

func (s *Server) GetShelf(ctx context.Context, req *librarypb.GetShelfRequest) (*librarypb.Shelf, error) {
	return s.repo.GetShelf(ctx, req)
}

func (s *Server) DeleteShelf(ctx context.Context, req *librarypb.DeleteShelfRequest) (*emptypb.Empty, error) {
	return s.repo.DeleteShelf(ctx, req)
}

func (s *Server) MergeShelves(ctx context.Context, req *librarypb.MergeShelvesRequest) (*librarypb.Shelf, error) {
	return s.repo.MergeShelves(ctx, req)
}

func (s *Server) CreateBook(ctx context.Context, req *librarypb.CreateBookRequest) (*librarypb.Book, error) {
	return s.repo.CreateBook(ctx, req)
}

func (s *Server) GetBook(ctx context.Context, req *librarypb.GetBookRequest) (*librarypb.Book, error) {
	return s.repo.GetBook(ctx, req)
}

func (s *Server) ListBooks(ctx context.Context, req *librarypb.ListBooksRequest) (*librarypb.ListBooksResponse, error) {
	return s.repo.ListBooks(ctx, req)
}

func (s *Server) DeleteBook(ctx context.Context, req *librarypb.DeleteBookRequest) (*emptypb.Empty, error) {
	return s.repo.DeleteBook(ctx, req)
}

func (s *Server) UpdateBook(ctx context.Context, req *librarypb.UpdateBookRequest) (*librarypb.Book, error) {
	return s.repo.UpdateBook(ctx, req)

}

func (s *Server) MoveBook(ctx context.Context, req *librarypb.MoveBookRequest) (*librarypb.Book, error) {
	return s.repo.MoveBook(ctx, req)
}
