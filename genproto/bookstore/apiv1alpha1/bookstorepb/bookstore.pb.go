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
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: qclaogui/bookstore/v1alpha1/bookstore.proto

package bookstorepb

import (
	reflect "reflect"
	sync "sync"

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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique shelf id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// A theme of the shelf (fiction, poetry, etc).
	Theme string `protobuf:"bytes,2,opt,name=theme,proto3" json:"theme,omitempty"`
}

func (x *Shelf) Reset() {
	*x = Shelf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shelf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shelf) ProtoMessage() {}

func (x *Shelf) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique book id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// An author of the book.
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// A book title.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Shelves in the bookstore.
	Shelves []*Shelf `protobuf:"bytes,1,rep,name=shelves,proto3" json:"shelves,omitempty"`
}

func (x *ListShelvesResponse) Reset() {
	*x = ListShelvesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShelvesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShelvesResponse) ProtoMessage() {}

func (x *ListShelvesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The shelf resource to create.
	Shelf *Shelf `protobuf:"bytes,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *CreateShelfRequest) Reset() {
	*x = CreateShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShelfRequest) ProtoMessage() {}

func (x *CreateShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the shelf resource to retrieve.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *GetShelfRequest) Reset() {
	*x = GetShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShelfRequest) ProtoMessage() {}

func (x *GetShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the shelf to delete.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *DeleteShelfRequest) Reset() {
	*x = DeleteShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShelfRequest) ProtoMessage() {}

func (x *DeleteShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the shelf which books to list.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *ListBooksRequest) Reset() {
	*x = ListBooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksRequest) ProtoMessage() {}

func (x *ListBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The books on the shelf.
	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *ListBooksResponse) Reset() {
	*x = ListBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksResponse) ProtoMessage() {}

func (x *ListBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the shelf on which to create a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// A book resource to create on the shelf.
	Book *Book `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the shelf from which to retrieve a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// The ID of the book to retrieve.
	Book int64 `protobuf:"varint,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *GetBookRequest) Reset() {
	*x = GetBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRequest) ProtoMessage() {}

func (x *GetBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the shelf from which to delete a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// The ID of the book to delete.
	Book int64 `protobuf:"varint,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x71,
	0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a,
	0x05, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x03, 0x52, 0x05,
	0x74, 0x68, 0x65, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x53, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x07, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73,
	0x22, 0x4e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66,
	0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x2a, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x28, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65,
	0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x22,
	0x4c, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x60, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x35, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22,
	0x3a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x3d, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x32, 0xae, 0x06, 0x0a, 0x10, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x59, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x30, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x2f, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x00,
	0x12, 0x5e, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x2c, 0x2e, 0x71,
	0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x00,
	0x12, 0x58, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12,
	0x2f, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x2d, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67,
	0x75, 0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x2e, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x2b, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x2e, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x1a, 0x19, 0xca, 0x41, 0x16, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x71,
	0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x45, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67,
	0x75, 0x69, 0x2f, 0x67, 0x61, 0x69, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescOnce sync.Once
	file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescData = file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDesc
)

func file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescGZIP() []byte {
	file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescOnce.Do(func() {
		file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescData)
	})
	return file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDescData
}

var file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_qclaogui_bookstore_v1alpha1_bookstore_proto_goTypes = []interface{}{
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
	if !protoimpl.UnsafeEnabled {
		file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shelf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShelvesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShelfRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShelfRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteShelfRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBooksRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBooksResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qclaogui_bookstore_v1alpha1_bookstore_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDesc,
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
	file_qclaogui_bookstore_v1alpha1_bookstore_proto_rawDesc = nil
	file_qclaogui_bookstore_v1alpha1_bookstore_proto_goTypes = nil
	file_qclaogui_bookstore_v1alpha1_bookstore_proto_depIdxs = nil
}
