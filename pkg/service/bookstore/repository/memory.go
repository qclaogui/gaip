// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"context"
	"flag"
	"fmt"
	"sync"

	pb "github.com/qclaogui/golang-api-server/api/bookstore/v1alpha1/bookstorepb"
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
	pb.UnimplementedBookstoreServiceServer

	// shelves are stored in a map keyed by shelf id
	// books are stored in a two level map, keyed first by shelf id and then by book id
	Shelves     map[int64]*pb.Shelf
	Books       map[int64]map[int64]*pb.Book
	LastShelfID int64      // the id of the last shelf that was added
	LastBookID  int64      // the id of the last book that was added
	Mutex       sync.Mutex // global mutex to synchronize service access
}

// NewMemoryRepo is a factory function to generate a new repository
func NewMemoryRepo() (*MemoryRepo, error) {
	mr := &MemoryRepo{
		Shelves: map[int64]*pb.Shelf{},
		Books:   map[int64]map[int64]*pb.Book{},
	}
	return mr, nil
}

// internal helpers
func (mr *MemoryRepo) getShelf(sid int64) (shelf *pb.Shelf, err error) {
	shelf, ok := mr.Shelves[sid]
	if !ok {
		return nil, fmt.Errorf("couldn't find shelf %d", sid)
	}
	return shelf, nil
}

func (mr *MemoryRepo) getBook(sid int64, bid int64) (book *pb.Book, err error) {
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

func (mr *MemoryRepo) ListShelves(context.Context, *emptypb.Empty) (*pb.ListShelvesResponse, error) {
	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()

	// copy shelf ids from Shelves map keys
	shelves := make([]*pb.Shelf, 0, len(mr.Shelves))
	for _, shelf := range mr.Shelves {
		shelves = append(shelves, shelf)
	}

	response := &pb.ListShelvesResponse{
		Shelves: shelves,
	}

	return response, nil
}

func (mr *MemoryRepo) CreateShelf(_ context.Context, req *pb.CreateShelfRequest) (*pb.Shelf, error) {
	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()

	// assign an id and name to a shelf and add it to the Shelves map.
	shelf := req.Shelf

	mr.LastShelfID++
	sid := mr.LastShelfID

	mr.Shelves[sid] = shelf

	return shelf, nil
}
func (mr *MemoryRepo) GetShelf(_ context.Context, req *pb.GetShelfRequest) (*pb.Shelf, error) {
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
func (mr *MemoryRepo) DeleteShelf(_ context.Context, req *pb.DeleteShelfRequest) (*emptypb.Empty, error) {
	sid := req.Shelf

	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()

	// delete a shelf by removing the shelf from the Shelves map and the associated books from the Books map.
	delete(mr.Shelves, sid)
	delete(mr.Books, sid)

	return nil, nil
}
func (mr *MemoryRepo) ListBooks(_ context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	sid := req.Shelf

	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()

	// list the books in a shelf
	_, err := mr.getShelf(sid)
	if err != nil {
		return nil, err
	}

	shelfBooks := mr.Books[sid]

	books := make([]*pb.Book, 0, len(shelfBooks))
	for _, book := range shelfBooks {
		books = append(books, book)
	}

	response := &pb.ListBooksResponse{
		Books: books,
	}

	return response, nil
}
func (mr *MemoryRepo) CreateBook(_ context.Context, req *pb.CreateBookRequest) (*pb.Book, error) {
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
		mr.Books[sid] = make(map[int64]*pb.Book)
	}

	mr.Books[sid][bid] = book

	return book, nil
}
func (mr *MemoryRepo) GetBook(_ context.Context, req *pb.GetBookRequest) (*pb.Book, error) {
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
func (mr *MemoryRepo) DeleteBook(_ context.Context, req *pb.DeleteBookRequest) (*emptypb.Empty, error) {
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
