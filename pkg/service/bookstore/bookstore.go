// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package bookstore

import (
	"context"

	"github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListShelves(ctx context.Context, req *emptypb.Empty) (*bookstorepb.ListShelvesResponse, error) {
	return s.repo.ListShelves(ctx, req)
}

func (s *Server) CreateShelf(ctx context.Context, req *bookstorepb.CreateShelfRequest) (*bookstorepb.Shelf, error) {
	return s.repo.CreateShelf(ctx, req)
}
func (s *Server) GetShelf(ctx context.Context, req *bookstorepb.GetShelfRequest) (*bookstorepb.Shelf, error) {
	return s.repo.GetShelf(ctx, req)
}
func (s *Server) DeleteShelf(ctx context.Context, req *bookstorepb.DeleteShelfRequest) (*emptypb.Empty, error) {
	return s.repo.DeleteShelf(ctx, req)

}
func (s *Server) ListBooks(ctx context.Context, req *bookstorepb.ListBooksRequest) (*bookstorepb.ListBooksResponse, error) {
	return s.repo.ListBooks(ctx, req)

}
func (s *Server) CreateBook(ctx context.Context, req *bookstorepb.CreateBookRequest) (*bookstorepb.Book, error) {
	return s.repo.CreateBook(ctx, req)
}
func (s *Server) GetBook(ctx context.Context, req *bookstorepb.GetBookRequest) (*bookstorepb.Book, error) {
	return s.repo.GetBook(ctx, req)
}
func (s *Server) DeleteBook(ctx context.Context, req *bookstorepb.DeleteBookRequest) (*emptypb.Empty, error) {
	return s.repo.DeleteBook(ctx, req)
}
