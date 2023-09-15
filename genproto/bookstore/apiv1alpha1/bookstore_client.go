// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.


package bookstore

import (
	"context"
	"math"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	bookstorepb "github.com/qclaogui/golang-api-server/genproto/bookstore/apiv1alpha1/bookstorepb"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ListShelves []gax.CallOption
	CreateShelf []gax.CallOption
	GetShelf []gax.CallOption
	DeleteShelf []gax.CallOption
	ListBooks []gax.CallOption
	CreateBook []gax.CallOption
	GetBook []gax.CallOption
	DeleteBook []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("api.qclaogui.com:443"),
		internaloption.WithDefaultMTLSEndpoint("api.qclaogui.com:443"),
		internaloption.WithDefaultAudience("https://api.qclaogui.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ListShelves: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateShelf: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetShelf: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteShelf: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListBooks: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateBook: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetBook: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteBook: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalClient is an interface that defines the methods available from .
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListShelves(context.Context, *emptypb.Empty, ...gax.CallOption) (*bookstorepb.ListShelvesResponse, error)
	CreateShelf(context.Context, *bookstorepb.CreateShelfRequest, ...gax.CallOption) (*bookstorepb.Shelf, error)
	GetShelf(context.Context, *bookstorepb.GetShelfRequest, ...gax.CallOption) (*bookstorepb.Shelf, error)
	DeleteShelf(context.Context, *bookstorepb.DeleteShelfRequest, ...gax.CallOption) error
	ListBooks(context.Context, *bookstorepb.ListBooksRequest, ...gax.CallOption) (*bookstorepb.ListBooksResponse, error)
	CreateBook(context.Context, *bookstorepb.CreateBookRequest, ...gax.CallOption) (*bookstorepb.Book, error)
	GetBook(context.Context, *bookstorepb.GetBookRequest, ...gax.CallOption) (*bookstorepb.Book, error)
	DeleteBook(context.Context, *bookstorepb.DeleteBookRequest, ...gax.CallOption) error
}

// Client is a client for interacting with .
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A simple Bookstore API.
//
// The API manages shelves and books resources. Shelves contain books.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListShelves returns a list of all shelves in the bookstore.
func (c *Client) ListShelves(ctx context.Context, req *emptypb.Empty, opts ...gax.CallOption) (*bookstorepb.ListShelvesResponse, error) {
	return c.internalClient.ListShelves(ctx, req, opts...)
}

// CreateShelf creates a new shelf in the bookstore.
func (c *Client) CreateShelf(ctx context.Context, req *bookstorepb.CreateShelfRequest, opts ...gax.CallOption) (*bookstorepb.Shelf, error) {
	return c.internalClient.CreateShelf(ctx, req, opts...)
}

// GetShelf returns a specific bookstore shelf.
func (c *Client) GetShelf(ctx context.Context, req *bookstorepb.GetShelfRequest, opts ...gax.CallOption) (*bookstorepb.Shelf, error) {
	return c.internalClient.GetShelf(ctx, req, opts...)
}

// DeleteShelf deletes a shelf, including all books that are stored on the shelf.
func (c *Client) DeleteShelf(ctx context.Context, req *bookstorepb.DeleteShelfRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteShelf(ctx, req, opts...)
}

// ListBooks returns a list of books on a shelf.
func (c *Client) ListBooks(ctx context.Context, req *bookstorepb.ListBooksRequest, opts ...gax.CallOption) (*bookstorepb.ListBooksResponse, error) {
	return c.internalClient.ListBooks(ctx, req, opts...)
}

// CreateBook creates a new book.
func (c *Client) CreateBook(ctx context.Context, req *bookstorepb.CreateBookRequest, opts ...gax.CallOption) (*bookstorepb.Book, error) {
	return c.internalClient.CreateBook(ctx, req, opts...)
}

// GetBook returns a specific book.
func (c *Client) GetBook(ctx context.Context, req *bookstorepb.GetBookRequest, opts ...gax.CallOption) (*bookstorepb.Book, error) {
	return c.internalClient.GetBook(ctx, req, opts...)
}

// DeleteBook deletes a book from a shelf.
func (c *Client) DeleteBook(ctx context.Context, req *bookstorepb.DeleteBookRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteBook(ctx, req, opts...)
}

// gRPCClient is a client for interacting with  over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client bookstorepb.BookstoreServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewClient creates a new bookstore service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A simple Bookstore API.
//
// The API manages shelves and books resources. Shelves contain books.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:    connPool,
		client: bookstorepb.NewBookstoreServiceClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) ListShelves(ctx context.Context, req *emptypb.Empty, opts ...gax.CallOption) (*bookstorepb.ListShelvesResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListShelves[0:len((*c.CallOptions).ListShelves):len((*c.CallOptions).ListShelves)], opts...)
	var resp *bookstorepb.ListShelvesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ListShelves(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateShelf(ctx context.Context, req *bookstorepb.CreateShelfRequest, opts ...gax.CallOption) (*bookstorepb.Shelf, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).CreateShelf[0:len((*c.CallOptions).CreateShelf):len((*c.CallOptions).CreateShelf)], opts...)
	var resp *bookstorepb.Shelf
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateShelf(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetShelf(ctx context.Context, req *bookstorepb.GetShelfRequest, opts ...gax.CallOption) (*bookstorepb.Shelf, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).GetShelf[0:len((*c.CallOptions).GetShelf):len((*c.CallOptions).GetShelf)], opts...)
	var resp *bookstorepb.Shelf
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetShelf(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteShelf(ctx context.Context, req *bookstorepb.DeleteShelfRequest, opts ...gax.CallOption) error {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).DeleteShelf[0:len((*c.CallOptions).DeleteShelf):len((*c.CallOptions).DeleteShelf)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteShelf(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) ListBooks(ctx context.Context, req *bookstorepb.ListBooksRequest, opts ...gax.CallOption) (*bookstorepb.ListBooksResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListBooks[0:len((*c.CallOptions).ListBooks):len((*c.CallOptions).ListBooks)], opts...)
	var resp *bookstorepb.ListBooksResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ListBooks(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateBook(ctx context.Context, req *bookstorepb.CreateBookRequest, opts ...gax.CallOption) (*bookstorepb.Book, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).CreateBook[0:len((*c.CallOptions).CreateBook):len((*c.CallOptions).CreateBook)], opts...)
	var resp *bookstorepb.Book
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateBook(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetBook(ctx context.Context, req *bookstorepb.GetBookRequest, opts ...gax.CallOption) (*bookstorepb.Book, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).GetBook[0:len((*c.CallOptions).GetBook):len((*c.CallOptions).GetBook)], opts...)
	var resp *bookstorepb.Book
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetBook(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteBook(ctx context.Context, req *bookstorepb.DeleteBookRequest, opts ...gax.CallOption) error {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).DeleteBook[0:len((*c.CallOptions).DeleteBook):len((*c.CallOptions).DeleteBook)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteBook(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}
