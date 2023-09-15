// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: library/v1/library.proto

package librarypb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A single book in the library.
type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the book.
	// Book names have the form `shelves/{shelf_id}/books/{book_id}`.
	// The name is ignored when creating a book.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The name of the book author.
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// The title of the book.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// Value indicating whether the book has been read.
	Read bool `protobuf:"varint,4,opt,name=read,proto3" json:"read,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_v1_library_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_library_v1_library_proto_msgTypes[0]
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
	return file_library_v1_library_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
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

func (x *Book) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

// A Shelf contains a collection of books with a theme.
type Shelf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the shelf.
	// Shelf names have the form `shelves/{shelf_id}`.
	// The name is ignored when creating a shelf.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The theme of the shelf
	Theme string `protobuf:"bytes,2,opt,name=theme,proto3" json:"theme,omitempty"`
}

func (x *Shelf) Reset() {
	*x = Shelf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_v1_library_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shelf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shelf) ProtoMessage() {}

func (x *Shelf) ProtoReflect() protoreflect.Message {
	mi := &file_library_v1_library_proto_msgTypes[1]
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
	return file_library_v1_library_proto_rawDescGZIP(), []int{1}
}

func (x *Shelf) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Shelf) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

// Request message for LibraryService.CreateShelf.
type CreateShelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The shelf to create.
	Shelf *Shelf `protobuf:"bytes,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *CreateShelfRequest) Reset() {
	*x = CreateShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_v1_library_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShelfRequest) ProtoMessage() {}

func (x *CreateShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_library_v1_library_proto_msgTypes[2]
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
	return file_library_v1_library_proto_rawDescGZIP(), []int{2}
}

func (x *CreateShelfRequest) GetShelf() *Shelf {
	if x != nil {
		return x.Shelf
	}
	return nil
}

// Request message for LibraryService.CreateBook.
type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the shelf in which the book is created.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The book to create.
	Book *Book `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_v1_library_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_library_v1_library_proto_msgTypes[3]
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
	return file_library_v1_library_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBookRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateBookRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

var File_library_v1_library_proto protoreflect.FileDescriptor

var file_library_v1_library_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x04, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x3a, 0x3f, 0xea, 0x41, 0x3c, 0x0a, 0x19, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1f, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65,
	0x73, 0x2f, 0x7b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x7d, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f,
	0x7b, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x22, 0x66, 0x0a, 0x05, 0x53, 0x68, 0x65,
	0x6c, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x3a, 0x33, 0xea, 0x41,
	0x30, 0x0a, 0x1a, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f,
	0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x12, 0x73,
	0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x5f, 0x69, 0x64,
	0x7d, 0x22, 0x42, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05,
	0x73, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x7a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xe0, 0x41, 0x02, 0xfa,
	0x41, 0x1c, 0x0a, 0x1a, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x62, 0x6f, 0x6f,
	0x6b, 0x32, 0x88, 0x02, 0x0a, 0x0e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x12, 0x1e, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x22, 0xda, 0x41, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x0b, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x12, 0x77, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1d, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x38, 0xda, 0x41, 0x0b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x2c, 0x62, 0x6f, 0x6f, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a,
	0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3d, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x1a, 0x17, 0xca, 0x41, 0x14, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x48, 0x5a, 0x46,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x6c, 0x61, 0x6f,
	0x67, 0x75, 0x69, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_library_v1_library_proto_rawDescOnce sync.Once
	file_library_v1_library_proto_rawDescData = file_library_v1_library_proto_rawDesc
)

func file_library_v1_library_proto_rawDescGZIP() []byte {
	file_library_v1_library_proto_rawDescOnce.Do(func() {
		file_library_v1_library_proto_rawDescData = protoimpl.X.CompressGZIP(file_library_v1_library_proto_rawDescData)
	})
	return file_library_v1_library_proto_rawDescData
}

var file_library_v1_library_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_library_v1_library_proto_goTypes = []interface{}{
	(*Book)(nil),               // 0: library.v1.Book
	(*Shelf)(nil),              // 1: library.v1.Shelf
	(*CreateShelfRequest)(nil), // 2: library.v1.CreateShelfRequest
	(*CreateBookRequest)(nil),  // 3: library.v1.CreateBookRequest
}
var file_library_v1_library_proto_depIdxs = []int32{
	1, // 0: library.v1.CreateShelfRequest.shelf:type_name -> library.v1.Shelf
	0, // 1: library.v1.CreateBookRequest.book:type_name -> library.v1.Book
	2, // 2: library.v1.LibraryService.CreateShelf:input_type -> library.v1.CreateShelfRequest
	3, // 3: library.v1.LibraryService.CreateBook:input_type -> library.v1.CreateBookRequest
	1, // 4: library.v1.LibraryService.CreateShelf:output_type -> library.v1.Shelf
	0, // 5: library.v1.LibraryService.CreateBook:output_type -> library.v1.Book
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_library_v1_library_proto_init() }
func file_library_v1_library_proto_init() {
	if File_library_v1_library_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_library_v1_library_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_library_v1_library_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_library_v1_library_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_library_v1_library_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_library_v1_library_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_library_v1_library_proto_goTypes,
		DependencyIndexes: file_library_v1_library_proto_depIdxs,
		MessageInfos:      file_library_v1_library_proto_msgTypes,
	}.Build()
	File_library_v1_library_proto = out.File
	file_library_v1_library_proto_rawDesc = nil
	file_library_v1_library_proto_goTypes = nil
	file_library_v1_library_proto_depIdxs = nil
}
