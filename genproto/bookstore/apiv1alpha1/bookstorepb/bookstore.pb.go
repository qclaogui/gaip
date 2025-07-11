// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: qclaogui/bookstore/v1alpha1/bookstore.proto

package bookstorepb

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A shelf resource.
type Shelf struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A unique shelf id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// A theme of the shelf (fiction, poetry, etc).
	Theme         string `protobuf:"bytes,2,opt,name=theme,proto3" json:"theme,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Shelf) Reset() {
	*x = Shelf{}
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Shelf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shelf) ProtoMessage() {}

func (x *Shelf) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shelf.ProtoReflect.Descriptor instead.
func (*Shelf) Descriptor() ([]byte, []int) {
	return file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescGZIP(), []int{0}
}

func (x *Shelf) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Shelf) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

// A book resource.
type Book struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A unique book id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// An author of the book.
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// A book title.
	Title         string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// Response to ListShelves call.
type ListShelvesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Shelves in the bookstore.
	Shelves       []*Shelf `protobuf:"bytes,1,rep,name=shelves,proto3" json:"shelves,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListShelvesResponse) Reset() {
	*x = ListShelvesResponse{}
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListShelvesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShelvesResponse) ProtoMessage() {}

func (x *ListShelvesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShelvesResponse.ProtoReflect.Descriptor instead.
func (*ListShelvesResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescGZIP(), []int{2}
}

func (x *ListShelvesResponse) GetShelves() []*Shelf {
	if x != nil {
		return x.Shelves
	}
	return nil
}

// Request message for CreateShelf method.
type CreateShelfRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The shelf resource to create.
	Shelf         *Shelf `protobuf:"bytes,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateShelfRequest) Reset() {
	*x = CreateShelfRequest{}
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShelfRequest) ProtoMessage() {}

func (x *CreateShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShelfRequest.ProtoReflect.Descriptor instead.
func (*CreateShelfRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescGZIP(), []int{3}
}

func (x *CreateShelfRequest) GetShelf() *Shelf {
	if x != nil {
		return x.Shelf
	}
	return nil
}

// Request message for GetShelf method.
type GetShelfRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the shelf resource to retrieve.
	Shelf         int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetShelfRequest) Reset() {
	*x = GetShelfRequest{}
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShelfRequest) ProtoMessage() {}

func (x *GetShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShelfRequest.ProtoReflect.Descriptor instead.
func (*GetShelfRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescGZIP(), []int{4}
}

func (x *GetShelfRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

// Request message for DeleteShelf method.
type DeleteShelfRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the shelf to delete.
	Shelf         int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteShelfRequest) Reset() {
	*x = DeleteShelfRequest{}
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShelfRequest) ProtoMessage() {}

func (x *DeleteShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShelfRequest.ProtoReflect.Descriptor instead.
func (*DeleteShelfRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteShelfRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

// Request message for ListBooks method.
type ListBooksRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the shelf which books to list.
	Shelf         int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBooksRequest) Reset() {
	*x = ListBooksRequest{}
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksRequest) ProtoMessage() {}

func (x *ListBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksRequest.ProtoReflect.Descriptor instead.
func (*ListBooksRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescGZIP(), []int{6}
}

func (x *ListBooksRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

// Response message to ListBooks method.
type ListBooksResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The books on the shelf.
	Books         []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBooksResponse) Reset() {
	*x = ListBooksResponse{}
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksResponse) ProtoMessage() {}

func (x *ListBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksResponse.ProtoReflect.Descriptor instead.
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescGZIP(), []int{7}
}

func (x *ListBooksResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

// Request message for CreateBook method.
type CreateBookRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the shelf on which to create a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// A book resource to create on the shelf.
	Book          *Book `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescGZIP(), []int{8}
}

func (x *CreateBookRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

func (x *CreateBookRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

// Request message for GetBook method.
type GetBookRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the shelf from which to retrieve a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// The ID of the book to retrieve.
	Book          int64 `protobuf:"varint,2,opt,name=book,proto3" json:"book,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookRequest) Reset() {
	*x = GetBookRequest{}
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRequest) ProtoMessage() {}

func (x *GetBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookRequest.ProtoReflect.Descriptor instead.
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescGZIP(), []int{9}
}

func (x *GetBookRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

func (x *GetBookRequest) GetBook() int64 {
	if x != nil {
		return x.Book
	}
	return 0
}

// Request message for DeleteBook method.
type DeleteBookRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the shelf from which to delete a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// The ID of the book to delete.
	Book          int64 `protobuf:"varint,2,opt,name=book,proto3" json:"book,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteBookRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

func (x *DeleteBookRequest) GetBook() int64 {
	if x != nil {
		return x.Book
	}
	return 0
}

var File_qclaogui_bookstore_v1alpha1_bookstore_proto protoreflect.FileDescriptor

const file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDesc = "" +
	"\n" +
	"+qclaogui/bookstore/v1alpha1/bookstore.proto\x12\x1bqclaogui.bookstore.v1alpha1\x1a\x1bbuf/validate/validate.proto\x1a\x17google/api/client.proto\x1a\x1bgoogle/protobuf/empty.proto\"6\n" +
	"\x05Shelf\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1d\n" +
	"\x05theme\x18\x02 \x01(\tB\a\xbaH\x04r\x02\x10\x03R\x05theme\"D\n" +
	"\x04Book\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x16\n" +
	"\x06author\x18\x02 \x01(\tR\x06author\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\"S\n" +
	"\x13ListShelvesResponse\x12<\n" +
	"\ashelves\x18\x01 \x03(\v2\".qclaogui.bookstore.v1alpha1.ShelfR\ashelves\"N\n" +
	"\x12CreateShelfRequest\x128\n" +
	"\x05shelf\x18\x01 \x01(\v2\".qclaogui.bookstore.v1alpha1.ShelfR\x05shelf\"'\n" +
	"\x0fGetShelfRequest\x12\x14\n" +
	"\x05shelf\x18\x01 \x01(\x03R\x05shelf\"*\n" +
	"\x12DeleteShelfRequest\x12\x14\n" +
	"\x05shelf\x18\x01 \x01(\x03R\x05shelf\"(\n" +
	"\x10ListBooksRequest\x12\x14\n" +
	"\x05shelf\x18\x01 \x01(\x03R\x05shelf\"L\n" +
	"\x11ListBooksResponse\x127\n" +
	"\x05books\x18\x01 \x03(\v2!.qclaogui.bookstore.v1alpha1.BookR\x05books\"`\n" +
	"\x11CreateBookRequest\x12\x14\n" +
	"\x05shelf\x18\x01 \x01(\x03R\x05shelf\x125\n" +
	"\x04book\x18\x02 \x01(\v2!.qclaogui.bookstore.v1alpha1.BookR\x04book\":\n" +
	"\x0eGetBookRequest\x12\x14\n" +
	"\x05shelf\x18\x01 \x01(\x03R\x05shelf\x12\x12\n" +
	"\x04book\x18\x02 \x01(\x03R\x04book\"=\n" +
	"\x11DeleteBookRequest\x12\x14\n" +
	"\x05shelf\x18\x01 \x01(\x03R\x05shelf\x12\x12\n" +
	"\x04book\x18\x02 \x01(\x03R\x04book2\xae\x06\n" +
	"\x10BookstoreService\x12Y\n" +
	"\vListShelves\x12\x16.google.protobuf.Empty\x1a0.qclaogui.bookstore.v1alpha1.ListShelvesResponse\"\x00\x12d\n" +
	"\vCreateShelf\x12/.qclaogui.bookstore.v1alpha1.CreateShelfRequest\x1a\".qclaogui.bookstore.v1alpha1.Shelf\"\x00\x12^\n" +
	"\bGetShelf\x12,.qclaogui.bookstore.v1alpha1.GetShelfRequest\x1a\".qclaogui.bookstore.v1alpha1.Shelf\"\x00\x12X\n" +
	"\vDeleteShelf\x12/.qclaogui.bookstore.v1alpha1.DeleteShelfRequest\x1a\x16.google.protobuf.Empty\"\x00\x12l\n" +
	"\tListBooks\x12-.qclaogui.bookstore.v1alpha1.ListBooksRequest\x1a..qclaogui.bookstore.v1alpha1.ListBooksResponse\"\x00\x12a\n" +
	"\n" +
	"CreateBook\x12..qclaogui.bookstore.v1alpha1.CreateBookRequest\x1a!.qclaogui.bookstore.v1alpha1.Book\"\x00\x12[\n" +
	"\aGetBook\x12+.qclaogui.bookstore.v1alpha1.GetBookRequest\x1a!.qclaogui.bookstore.v1alpha1.Book\"\x00\x12V\n" +
	"\n" +
	"DeleteBook\x12..qclaogui.bookstore.v1alpha1.DeleteBookRequest\x1a\x16.google.protobuf.Empty\"\x00\x1a\x19\xcaA\x16bookstore.qclaogui.comBEZCgithub.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepbb\x06proto3"

var (
	file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescOnce sync.Once
	file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescData []byte
)

func file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescGZIP() []byte {
	file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescOnce.Do(func() {
		file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDesc), len(file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDesc)))
	})
	return file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescData
}

var (
	file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
	file_qclaogui_bookstore_v1alpha1_bookstore_proto_goTypes  = []any{
		(*Shelf)(nil),               // 0: qclaogui.bookstore.v1alpha1.Shelf
		(*Book)(nil),                // 1: qclaogui.bookstore.v1alpha1.Book
		(*ListShelvesResponse)(nil), // 2: qclaogui.bookstore.v1alpha1.ListShelvesResponse
		(*CreateShelfRequest)(nil),  // 3: qclaogui.bookstore.v1alpha1.CreateShelfRequest
		(*GetShelfRequest)(nil),     // 4: qclaogui.bookstore.v1alpha1.GetShelfRequest
		(*DeleteShelfRequest)(nil),  // 5: qclaogui.bookstore.v1alpha1.DeleteShelfRequest
		(*ListBooksRequest)(nil),    // 6: qclaogui.bookstore.v1alpha1.ListBooksRequest
		(*ListBooksResponse)(nil),   // 7: qclaogui.bookstore.v1alpha1.ListBooksResponse
		(*CreateBookRequest)(nil),   // 8: qclaogui.bookstore.v1alpha1.CreateBookRequest
		(*GetBookRequest)(nil),      // 9: qclaogui.bookstore.v1alpha1.GetBookRequest
		(*DeleteBookRequest)(nil),   // 10: qclaogui.bookstore.v1alpha1.DeleteBookRequest
		(*emptypb.Empty)(nil),       // 11: google.protobuf.Empty
	}
)

var file_qclaogui_bookstore_v1alpha1_bookstore_proto_depIdxs = []int32{
	0,  // 0: qclaogui.bookstore.v1alpha1.ListShelvesResponse.shelves:type_name -> qclaogui.bookstore.v1alpha1.Shelf
	0,  // 1: qclaogui.bookstore.v1alpha1.CreateShelfRequest.shelf:type_name -> qclaogui.bookstore.v1alpha1.Shelf
	1,  // 2: qclaogui.bookstore.v1alpha1.ListBooksResponse.books:type_name -> qclaogui.bookstore.v1alpha1.Book
	1,  // 3: qclaogui.bookstore.v1alpha1.CreateBookRequest.book:type_name -> qclaogui.bookstore.v1alpha1.Book
	11, // 4: qclaogui.bookstore.v1alpha1.BookstoreService.ListShelves:input_type -> google.protobuf.Empty
	3,  // 5: qclaogui.bookstore.v1alpha1.BookstoreService.CreateShelf:input_type -> qclaogui.bookstore.v1alpha1.CreateShelfRequest
	4,  // 6: qclaogui.bookstore.v1alpha1.BookstoreService.GetShelf:input_type -> qclaogui.bookstore.v1alpha1.GetShelfRequest
	5,  // 7: qclaogui.bookstore.v1alpha1.BookstoreService.DeleteShelf:input_type -> qclaogui.bookstore.v1alpha1.DeleteShelfRequest
	6,  // 8: qclaogui.bookstore.v1alpha1.BookstoreService.ListBooks:input_type -> qclaogui.bookstore.v1alpha1.ListBooksRequest
	8,  // 9: qclaogui.bookstore.v1alpha1.BookstoreService.CreateBook:input_type -> qclaogui.bookstore.v1alpha1.CreateBookRequest
	9,  // 10: qclaogui.bookstore.v1alpha1.BookstoreService.GetBook:input_type -> qclaogui.bookstore.v1alpha1.GetBookRequest
	10, // 11: qclaogui.bookstore.v1alpha1.BookstoreService.DeleteBook:input_type -> qclaogui.bookstore.v1alpha1.DeleteBookRequest
	2,  // 12: qclaogui.bookstore.v1alpha1.BookstoreService.ListShelves:output_type -> qclaogui.bookstore.v1alpha1.ListShelvesResponse
	0,  // 13: qclaogui.bookstore.v1alpha1.BookstoreService.CreateShelf:output_type -> qclaogui.bookstore.v1alpha1.Shelf
	0,  // 14: qclaogui.bookstore.v1alpha1.BookstoreService.GetShelf:output_type -> qclaogui.bookstore.v1alpha1.Shelf
	11, // 15: qclaogui.bookstore.v1alpha1.BookstoreService.DeleteShelf:output_type -> google.protobuf.Empty
	7,  // 16: qclaogui.bookstore.v1alpha1.BookstoreService.ListBooks:output_type -> qclaogui.bookstore.v1alpha1.ListBooksResponse
	1,  // 17: qclaogui.bookstore.v1alpha1.BookstoreService.CreateBook:output_type -> qclaogui.bookstore.v1alpha1.Book
	1,  // 18: qclaogui.bookstore.v1alpha1.BookstoreService.GetBook:output_type -> qclaogui.bookstore.v1alpha1.Book
	11, // 19: qclaogui.bookstore.v1alpha1.BookstoreService.DeleteBook:output_type -> google.protobuf.Empty
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_qclaogui_bookstore_v1alpha1_bookstore_proto_init() }
func file_qclaogui_bookstore_v1alpha1_bookstore_proto_init() {
	if File_qclaogui_bookstore_v1alpha1_bookstore_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDesc), len(file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qclaogui_bookstore_v1alpha1_bookstore_proto_goTypes,
		DependencyIndexes: file_qclaogui_bookstore_v1alpha1_bookstore_proto_depIdxs,
		MessageInfos:      file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes,
	}.Build()
	File_qclaogui_bookstore_v1alpha1_bookstore_proto = out.File
	file_qclaogui_bookstore_v1alpha1_bookstore_proto_goTypes = nil
	file_qclaogui_bookstore_v1alpha1_bookstore_proto_depIdxs = nil
}
