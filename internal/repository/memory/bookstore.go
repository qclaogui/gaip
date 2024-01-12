// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// NewBookstore is a factory function to generate a new repository
func NewBookstore() (bookstorepb.BookstoreServiceServer, error) {
	m := &Bookstore{
		Shelves: map[int64]*bookstorepb.Shelf{},
		Books:   map[int64]map[int64]*bookstorepb.Book{},
	}
	return m, nil
}

// Bookstore fulfills the Bookstore Repository interface
// All objects are managed in an in-memory non-persistent store.
//
// Bookstore is used to implement bookstorepb.BookstoreServiceServer.
type Bookstore struct {
	bookstorepb.UnimplementedBookstoreServiceServer

	// shelves are stored in a map keyed by shelf id
	// books are stored in a two level map, keyed first by shelf id and then by book id
	Shelves     map[int64]*bookstorepb.Shelf
	Books       map[int64]map[int64]*bookstorepb.Book
	LastShelfID int64      // the id of the last shelf that was added
	LastBookID  int64      // the id of the last book that was added
	Mutex       sync.Mutex // global mutex to synchronize service access
}

func (m *Bookstore) getShelf(sid int64) (shelf *bookstorepb.Shelf, err error) {
	shelf, ok := m.Shelves[sid]
	if !ok {
		return nil, fmt.Errorf("couldn't find shelf %d", sid)
	}
	return shelf, nil
}

func (m *Bookstore) getBook(sid int64, bid int64) (book *bookstorepb.Book, err error) {
	_, err = m.getShelf(sid)
	if err != nil {
		return nil, err
	}
	book, ok := m.Books[sid][bid]
	if !ok {
		return nil, fmt.Errorf("couldn't find book %d on shelf %d", bid, sid)
	}

	return book, nil
}

func (m *Bookstore) ListShelves(context.Context, *emptypb.Empty) (*bookstorepb.ListShelvesResponse, error) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	// copy shelf ids from Shelves map keys
	shelves := make([]*bookstorepb.Shelf, 0, len(m.Shelves))
	for _, shelf := range m.Shelves {
		shelves = append(shelves, shelf)
	}

	response := &bookstorepb.ListShelvesResponse{
		Shelves: shelves,
	}

	return response, nil
}

func (m *Bookstore) CreateShelf(_ context.Context, req *bookstorepb.CreateShelfRequest) (*bookstorepb.Shelf, error) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	// assign an id and name to a shelf and add it to the Shelves map.
	shelf := req.Shelf

	m.LastShelfID++
	sid := m.LastShelfID

	m.Shelves[sid] = shelf

	return shelf, nil
}
func (m *Bookstore) GetShelf(_ context.Context, req *bookstorepb.GetShelfRequest) (*bookstorepb.Shelf, error) {
	sid := req.Shelf

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	// look up a shelf from the Shelves map.
	shelf, err := m.getShelf(sid)
	if err != nil {
		return nil, err
	}

	return shelf, nil
}
func (m *Bookstore) DeleteShelf(_ context.Context, req *bookstorepb.DeleteShelfRequest) (*emptypb.Empty, error) {
	sid := req.Shelf

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	// delete a shelf by removing the shelf from the Shelves map and the associated books from the Books map.
	delete(m.Shelves, sid)
	delete(m.Books, sid)

	return nil, nil
}
func (m *Bookstore) ListBooks(_ context.Context, req *bookstorepb.ListBooksRequest) (*bookstorepb.ListBooksResponse, error) {
	sid := req.Shelf

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	// list the books in a shelf
	_, err := m.getShelf(sid)
	if err != nil {
		return nil, err
	}

	shelfBooks := m.Books[sid]

	books := make([]*bookstorepb.Book, 0, len(shelfBooks))
	for _, book := range shelfBooks {
		books = append(books, book)
	}

	response := &bookstorepb.ListBooksResponse{
		Books: books,
	}

	return response, nil
}
func (m *Bookstore) CreateBook(_ context.Context, req *bookstorepb.CreateBookRequest) (*bookstorepb.Book, error) {
	sid := req.Shelf

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	_, err := m.getShelf(sid)
	if err != nil {
		return nil, err
	}

	// assign an id and name to a book and add it to the Books map.
	m.LastBookID++
	bid := m.LastBookID

	book := req.Book
	if m.Books[sid] == nil {
		m.Books[sid] = make(map[int64]*bookstorepb.Book)
	}

	m.Books[sid][bid] = book

	return book, nil
}
func (m *Bookstore) GetBook(_ context.Context, req *bookstorepb.GetBookRequest) (*bookstorepb.Book, error) {
	sid, bid := req.Shelf, req.Book

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	// get a book from the Books map
	book, err := m.getBook(sid, bid)
	if err != nil {
		return nil, err
	}

	return book, nil
}
func (m *Bookstore) DeleteBook(_ context.Context, req *bookstorepb.DeleteBookRequest) (*emptypb.Empty, error) {
	sid, bid := req.Shelf, req.Book

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	_, err := m.getShelf(sid)
	if err != nil {
		return nil, err
	}

	// delete a book by removing the book from the Books map.
	delete(m.Books[sid], bid)

	return nil, nil
}
