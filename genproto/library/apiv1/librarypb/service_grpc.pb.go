// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: qclaogui/library/v1/service.proto

package librarypb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LibraryService_CreateShelf_FullMethodName  = "/qclaogui.library.v1.LibraryService/CreateShelf"
	LibraryService_GetShelf_FullMethodName     = "/qclaogui.library.v1.LibraryService/GetShelf"
	LibraryService_ListShelves_FullMethodName  = "/qclaogui.library.v1.LibraryService/ListShelves"
	LibraryService_DeleteShelf_FullMethodName  = "/qclaogui.library.v1.LibraryService/DeleteShelf"
	LibraryService_MergeShelves_FullMethodName = "/qclaogui.library.v1.LibraryService/MergeShelves"
	LibraryService_CreateBook_FullMethodName   = "/qclaogui.library.v1.LibraryService/CreateBook"
	LibraryService_GetBook_FullMethodName      = "/qclaogui.library.v1.LibraryService/GetBook"
	LibraryService_ListBooks_FullMethodName    = "/qclaogui.library.v1.LibraryService/ListBooks"
	LibraryService_DeleteBook_FullMethodName   = "/qclaogui.library.v1.LibraryService/DeleteBook"
	LibraryService_UpdateBook_FullMethodName   = "/qclaogui.library.v1.LibraryService/UpdateBook"
	LibraryService_MoveBook_FullMethodName     = "/qclaogui.library.v1.LibraryService/MoveBook"
)

// LibraryServiceClient is the client API for LibraryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// This API represents a simple digital library. It lets you manage Shelf
// resources and Book resources in the library. It defines the following
// resource model:
//
//   - The API has a collection of [Shelf][qclaogui.library.v1.Shelf]
//     resources, named `shelves/*`
//
//   - Each Shelf has a collection of [Book][qclaogui.library.v1.Book]
//     resources, named `shelves/*/books/*`
type LibraryServiceClient interface {
	// Creates a shelf, and returns the new Shelf.
	CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*Shelf, error)
	// Gets a shelf. Returns NOT_FOUND if the shelf does not exist.
	GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error)
	// Lists shelves. The order is unspecified but deterministic. Newly created
	// shelves will not necessarily be added to the end of this list.
	ListShelves(ctx context.Context, in *ListShelvesRequest, opts ...grpc.CallOption) (*ListShelvesResponse, error)
	// Deletes a shelf. Returns NOT_FOUND if the shelf does not exist.
	DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Merges two shelves by adding all books from the shelf named
	// `other_shelf_name` to shelf `name`, and deletes
	// `other_shelf_name`. Returns the updated shelf.
	// The book ids of the moved books may not be the same as the original books.
	//
	// Returns NOT_FOUND if either shelf does not exist.
	// This call is a no-op if the specified shelves are the same.
	MergeShelves(ctx context.Context, in *MergeShelvesRequest, opts ...grpc.CallOption) (*Shelf, error)
	// Creates a book, and returns the new Book.
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Gets a book. Returns NOT_FOUND if the book does not exist.
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Lists books in a shelf. The order is unspecified but deterministic. Newly
	// created books will not necessarily be added to the end of this list.
	// Returns NOT_FOUND if the shelf does not exist.
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	// Deletes a book. Returns NOT_FOUND if the book does not exist.
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Updates a book. Returns INVALID_ARGUMENT if the name of the book
	// is non-empty and does not equal the existing name.
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Moves a book to another shelf, and returns the new book. The book
	// id of the new book may not be the same as the original book.
	MoveBook(ctx context.Context, in *MoveBookRequest, opts ...grpc.CallOption) (*Book, error)
}

type libraryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLibraryServiceClient(cc grpc.ClientConnInterface) LibraryServiceClient {
	return &libraryServiceClient{cc}
}

func (c *libraryServiceClient) CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*Shelf, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Shelf)
	err := c.cc.Invoke(ctx, LibraryService_CreateShelf_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Shelf)
	err := c.cc.Invoke(ctx, LibraryService_GetShelf_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListShelves(ctx context.Context, in *ListShelvesRequest, opts ...grpc.CallOption) (*ListShelvesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListShelvesResponse)
	err := c.cc.Invoke(ctx, LibraryService_ListShelves_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LibraryService_DeleteShelf_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) MergeShelves(ctx context.Context, in *MergeShelvesRequest, opts ...grpc.CallOption) (*Shelf, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Shelf)
	err := c.cc.Invoke(ctx, LibraryService_MergeShelves_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Book)
	err := c.cc.Invoke(ctx, LibraryService_CreateBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Book)
	err := c.cc.Invoke(ctx, LibraryService_GetBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBooksResponse)
	err := c.cc.Invoke(ctx, LibraryService_ListBooks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LibraryService_DeleteBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Book)
	err := c.cc.Invoke(ctx, LibraryService_UpdateBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) MoveBook(ctx context.Context, in *MoveBookRequest, opts ...grpc.CallOption) (*Book, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Book)
	err := c.cc.Invoke(ctx, LibraryService_MoveBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibraryServiceServer is the server API for LibraryService service.
// All implementations should embed UnimplementedLibraryServiceServer
// for forward compatibility.
//
// This API represents a simple digital library. It lets you manage Shelf
// resources and Book resources in the library. It defines the following
// resource model:
//
//   - The API has a collection of [Shelf][qclaogui.library.v1.Shelf]
//     resources, named `shelves/*`
//
//   - Each Shelf has a collection of [Book][qclaogui.library.v1.Book]
//     resources, named `shelves/*/books/*`
type LibraryServiceServer interface {
	// Creates a shelf, and returns the new Shelf.
	CreateShelf(context.Context, *CreateShelfRequest) (*Shelf, error)
	// Gets a shelf. Returns NOT_FOUND if the shelf does not exist.
	GetShelf(context.Context, *GetShelfRequest) (*Shelf, error)
	// Lists shelves. The order is unspecified but deterministic. Newly created
	// shelves will not necessarily be added to the end of this list.
	ListShelves(context.Context, *ListShelvesRequest) (*ListShelvesResponse, error)
	// Deletes a shelf. Returns NOT_FOUND if the shelf does not exist.
	DeleteShelf(context.Context, *DeleteShelfRequest) (*emptypb.Empty, error)
	// Merges two shelves by adding all books from the shelf named
	// `other_shelf_name` to shelf `name`, and deletes
	// `other_shelf_name`. Returns the updated shelf.
	// The book ids of the moved books may not be the same as the original books.
	//
	// Returns NOT_FOUND if either shelf does not exist.
	// This call is a no-op if the specified shelves are the same.
	MergeShelves(context.Context, *MergeShelvesRequest) (*Shelf, error)
	// Creates a book, and returns the new Book.
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	// Gets a book. Returns NOT_FOUND if the book does not exist.
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	// Lists books in a shelf. The order is unspecified but deterministic. Newly
	// created books will not necessarily be added to the end of this list.
	// Returns NOT_FOUND if the shelf does not exist.
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	// Deletes a book. Returns NOT_FOUND if the book does not exist.
	DeleteBook(context.Context, *DeleteBookRequest) (*emptypb.Empty, error)
	// Updates a book. Returns INVALID_ARGUMENT if the name of the book
	// is non-empty and does not equal the existing name.
	UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
	// Moves a book to another shelf, and returns the new book. The book
	// id of the new book may not be the same as the original book.
	MoveBook(context.Context, *MoveBookRequest) (*Book, error)
}

// UnimplementedLibraryServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLibraryServiceServer struct{}

func (UnimplementedLibraryServiceServer) CreateShelf(context.Context, *CreateShelfRequest) (*Shelf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShelf not implemented")
}

func (UnimplementedLibraryServiceServer) GetShelf(context.Context, *GetShelfRequest) (*Shelf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShelf not implemented")
}

func (UnimplementedLibraryServiceServer) ListShelves(context.Context, *ListShelvesRequest) (*ListShelvesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShelves not implemented")
}

func (UnimplementedLibraryServiceServer) DeleteShelf(context.Context, *DeleteShelfRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShelf not implemented")
}

func (UnimplementedLibraryServiceServer) MergeShelves(context.Context, *MergeShelvesRequest) (*Shelf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MergeShelves not implemented")
}

func (UnimplementedLibraryServiceServer) CreateBook(context.Context, *CreateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}

func (UnimplementedLibraryServiceServer) GetBook(context.Context, *GetBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}

func (UnimplementedLibraryServiceServer) ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}

func (UnimplementedLibraryServiceServer) DeleteBook(context.Context, *DeleteBookRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}

func (UnimplementedLibraryServiceServer) UpdateBook(context.Context, *UpdateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}

func (UnimplementedLibraryServiceServer) MoveBook(context.Context, *MoveBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveBook not implemented")
}
func (UnimplementedLibraryServiceServer) testEmbeddedByValue() {}

// UnsafeLibraryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibraryServiceServer will
// result in compilation errors.
type UnsafeLibraryServiceServer interface {
	mustEmbedUnimplementedLibraryServiceServer()
}

func RegisterLibraryServiceServer(s grpc.ServiceRegistrar, srv LibraryServiceServer) {
	// If the following call pancis, it indicates UnimplementedLibraryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LibraryService_ServiceDesc, srv)
}

func _LibraryService_CreateShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).CreateShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_CreateShelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).CreateShelf(ctx, req.(*CreateShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_GetShelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetShelf(ctx, req.(*GetShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListShelves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShelvesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListShelves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_ListShelves_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListShelves(ctx, req.(*ListShelvesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_DeleteShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).DeleteShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_DeleteShelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).DeleteShelf(ctx, req.(*DeleteShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_MergeShelves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeShelvesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).MergeShelves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_MergeShelves_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).MergeShelves(ctx, req.(*MergeShelvesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_CreateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_GetBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_ListBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_DeleteBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_UpdateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_MoveBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).MoveBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_MoveBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).MoveBook(ctx, req.(*MoveBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LibraryService_ServiceDesc is the grpc.ServiceDesc for LibraryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LibraryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qclaogui.library.v1.LibraryService",
	HandlerType: (*LibraryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShelf",
			Handler:    _LibraryService_CreateShelf_Handler,
		},
		{
			MethodName: "GetShelf",
			Handler:    _LibraryService_GetShelf_Handler,
		},
		{
			MethodName: "ListShelves",
			Handler:    _LibraryService_ListShelves_Handler,
		},
		{
			MethodName: "DeleteShelf",
			Handler:    _LibraryService_DeleteShelf_Handler,
		},
		{
			MethodName: "MergeShelves",
			Handler:    _LibraryService_MergeShelves_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _LibraryService_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _LibraryService_GetBook_Handler,
		},
		{
			MethodName: "ListBooks",
			Handler:    _LibraryService_ListBooks_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _LibraryService_DeleteBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _LibraryService_UpdateBook_Handler,
		},
		{
			MethodName: "MoveBook",
			Handler:    _LibraryService_MoveBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qclaogui/library/v1/service.proto",
}
