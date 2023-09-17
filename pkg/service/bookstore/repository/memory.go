// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"context"
	"flag"
	"fmt"
	"sync"

	"github.com/qclaogui/golang-api-server/genproto/bookstore/apiv1alpha1/bookstorepb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MemoryConfig struct {
	Enabled bool `yaml:"enabled"`
}

func (cfg *MemoryConfig) RegisterFlags(fs *flag.FlagSet) {
	cfg.RegisterFlagsWithPrefix("", fs)
}

func (cfg *MemoryConfig) RegisterFlagsWithPrefix(prefix string, fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, prefix+"memory.enabled", false, "Enables memory Repository")
}

func (cfg *MemoryConfig) Validate() error {
	return nil
}

// MemoryRepo fulfills the Repository interface
// All objects are managed in an in-memory non-persistent store.
//
// MemoryRepo is used to implement BookstoreServiceServer.
type MemoryRepo struct {
	bookstorepb.UnimplementedBookstoreServiceServer

	// shelves are stored in a map keyed by shelf id
	// books are stored in a two level map, keyed first by shelf id and then by book id
	Shelves     map[int64]*bookstorepb.Shelf
	Books       map[int64]map[int64]*bookstorepb.Book
	LastShelfID int64      // the id of the last shelf that was added
	LastBookID  int64      // the id of the last book that was added
	Mutex       sync.Mutex // global mutex to synchronize service access
}

// NewMemoryRepo is a factory function to generate a new repository
func NewMemoryRepo() (*MemoryRepo, error) {
	mr := &MemoryRepo{
		Shelves: map[int64]*bookstorepb.Shelf{},
		Books:   map[int64]map[int64]*bookstorepb.Book{},
	}
	return mr, nil
}

// internal helpers
func (mr *MemoryRepo) getShelf(sid int64) (shelf *bookstorepb.Shelf, err error) {
	shelf, ok := mr.Shelves[sid]
	if !ok {
		return nil, fmt.Errorf("couldn't find shelf %d", sid)
	}
	return shelf, nil
}

func (mr *MemoryRepo) getBook(sid int64, bid int64) (book *bookstorepb.Book, err error) {
	_, err = mr.getShelf(sid)
	if err != nil {
		return nil, err
	}
	book, ok := mr.Books[sid][bid]
	if !ok {
		return nil, fmt.Errorf("couldn't find book %d on shelf %d", bid, sid)
	}

	return book, nil
}

func (mr *MemoryRepo) ListShelves(context.Context, *emptypb.Empty) (*bookstorepb.ListShelvesResponse, error) {
	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()

	// copy shelf ids from Shelves map keys
	shelves := make([]*bookstorepb.Shelf, 0, len(mr.Shelves))
	for _, shelf := range mr.Shelves {
		shelves = append(shelves, shelf)
	}

	response := &bookstorepb.ListShelvesResponse{
		Shelves: shelves,
	}

	return response, nil
}

func (mr *MemoryRepo) CreateShelf(_ context.Context, req *bookstorepb.CreateShelfRequest) (*bookstorepb.Shelf, error) {
	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()

	// assign an id and name to a shelf and add it to the Shelves map.
	shelf := req.Shelf

	mr.LastShelfID++
	sid := mr.LastShelfID

	mr.Shelves[sid] = shelf

	return shelf, nil
}
func (mr *MemoryRepo) GetShelf(_ context.Context, req *bookstorepb.GetShelfRequest) (*bookstorepb.Shelf, error) {
	sid := req.Shelf

	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()

	// look up a shelf from the Shelves map.
	shelf, err := mr.getShelf(sid)
	if err != nil {
		return nil, err
	}

	return shelf, nil
}
func (mr *MemoryRepo) DeleteShelf(_ context.Context, req *bookstorepb.DeleteShelfRequest) (*emptypb.Empty, error) {
	sid := req.Shelf

	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()

	// delete a shelf by removing the shelf from the Shelves map and the associated books from the Books map.
	delete(mr.Shelves, sid)
	delete(mr.Books, sid)

	return nil, nil
}
func (mr *MemoryRepo) ListBooks(_ context.Context, req *bookstorepb.ListBooksRequest) (*bookstorepb.ListBooksResponse, error) {
	sid := req.Shelf

	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()

	// list the books in a shelf
	_, err := mr.getShelf(sid)
	if err != nil {
		return nil, err
	}

	shelfBooks := mr.Books[sid]

	books := make([]*bookstorepb.Book, 0, len(shelfBooks))
	for _, book := range shelfBooks {
		books = append(books, book)
	}

	response := &bookstorepb.ListBooksResponse{
		Books: books,
	}

	return response, nil
}
func (mr *MemoryRepo) CreateBook(_ context.Context, req *bookstorepb.CreateBookRequest) (*bookstorepb.Book, error) {
	sid := req.Shelf

	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()

	_, err := mr.getShelf(sid)
	if err != nil {
		return nil, err
	}

	// assign an id and name to a book and add it to the Books map.
	mr.LastBookID++
	bid := mr.LastBookID

	book := req.Book
	if mr.Books[sid] == nil {
		mr.Books[sid] = make(map[int64]*bookstorepb.Book)
	}

	mr.Books[sid][bid] = book

	return book, nil
}
func (mr *MemoryRepo) GetBook(_ context.Context, req *bookstorepb.GetBookRequest) (*bookstorepb.Book, error) {
	sid, bid := req.Shelf, req.Book

	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()

	// get a book from the Books map
	book, err := mr.getBook(sid, bid)
	if err != nil {
		return nil, err
	}

	return book, nil
}
func (mr *MemoryRepo) DeleteBook(_ context.Context, req *bookstorepb.DeleteBookRequest) (*emptypb.Empty, error) {
	sid, bid := req.Shelf, req.Book

	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()

	_, err := mr.getShelf(sid)
	if err != nil {
		return nil, err
	}

	// delete a book by removing the book from the Books map.
	delete(mr.Books[sid], bid)

	return nil, nil
}
