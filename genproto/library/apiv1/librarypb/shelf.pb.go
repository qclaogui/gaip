// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: qclaogui/library/v1/shelf.proto

package librarypb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

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
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Shelf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shelf) ProtoMessage() {}

func (x *Shelf) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[0]
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
	return file_qclaogui_library_v1_shelf_proto_rawDescGZIP(), []int{0}
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
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShelfRequest) ProtoMessage() {}

func (x *CreateShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[1]
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
	return file_qclaogui_library_v1_shelf_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShelfRequest) GetShelf() *Shelf {
	if x != nil {
		return x.Shelf
	}
	return nil
}

// Request message for LibraryService.GetShelf.
type GetShelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the shelf to retrieve.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetShelfRequest) Reset() {
	*x = GetShelfRequest{}
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShelfRequest) ProtoMessage() {}

func (x *GetShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[2]
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
	return file_qclaogui_library_v1_shelf_proto_rawDescGZIP(), []int{2}
}

func (x *GetShelfRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for LibraryService.ListShelves.
type ListShelvesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested page size. Server may return fewer shelves than requested.
	// If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A token identifying a page of results the server should return.
	// Typically, this is the value of
	// [ListShelvesResponse.next_page_token][qclaogui.library.v1.ListShelvesResponse.next_page_token]
	// returned from the previous call to `ListShelves` method.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListShelvesRequest) Reset() {
	*x = ListShelvesRequest{}
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListShelvesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShelvesRequest) ProtoMessage() {}

func (x *ListShelvesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShelvesRequest.ProtoReflect.Descriptor instead.
func (*ListShelvesRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_library_v1_shelf_proto_rawDescGZIP(), []int{3}
}

func (x *ListShelvesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListShelvesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response message for LibraryService.ListShelves.
type ListShelvesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of shelves.
	Shelves []*Shelf `protobuf:"bytes,1,rep,name=shelves,proto3" json:"shelves,omitempty"`
	// A token to retrieve next page of results.
	// Pass this value in the
	// [ListShelvesRequest.page_token][qclaogui.library.v1.ListShelvesRequest.page_token]
	// field in the subsequent call to `ListShelves` method to retrieve the next
	// page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListShelvesResponse) Reset() {
	*x = ListShelvesResponse{}
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListShelvesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShelvesResponse) ProtoMessage() {}

func (x *ListShelvesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[4]
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
	return file_qclaogui_library_v1_shelf_proto_rawDescGZIP(), []int{4}
}

func (x *ListShelvesResponse) GetShelves() []*Shelf {
	if x != nil {
		return x.Shelves
	}
	return nil
}

func (x *ListShelvesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request message for LibraryService.DeleteShelf.
type DeleteShelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the shelf to delete.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteShelfRequest) Reset() {
	*x = DeleteShelfRequest{}
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShelfRequest) ProtoMessage() {}

func (x *DeleteShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[5]
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
	return file_qclaogui_library_v1_shelf_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteShelfRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Describes the shelf being removed (other_shelf_name) and updated
// (name) in this merge.
type MergeShelvesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the shelf we're adding books to.
	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OtherShelf string `protobuf:"bytes,2,opt,name=other_shelf,json=otherShelf,proto3" json:"other_shelf,omitempty"`
}

func (x *MergeShelvesRequest) Reset() {
	*x = MergeShelvesRequest{}
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MergeShelvesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeShelvesRequest) ProtoMessage() {}

func (x *MergeShelvesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_library_v1_shelf_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeShelvesRequest.ProtoReflect.Descriptor instead.
func (*MergeShelvesRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_library_v1_shelf_proto_rawDescGZIP(), []int{6}
}

func (x *MergeShelvesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MergeShelvesRequest) GetOtherShelf() string {
	if x != nil {
		return x.OtherShelf
	}
	return ""
}

var File_qclaogui_library_v1_shelf_proto protoreflect.FileDescriptor

var file_qclaogui_library_v1_shelf_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x66, 0x0a, 0x05, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x68, 0x65, 0x6d, 0x65, 0x3a, 0x33, 0xea, 0x41, 0x30, 0x0a, 0x1a, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x12, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x2f,
	0x7b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x5f, 0x69, 0x64, 0x7d, 0x22, 0x4c, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x42, 0x04, 0xe2, 0x41, 0x01,
	0x02, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x4a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xe2, 0x41, 0x01, 0x02, 0xfa,
	0x41, 0x1c, 0x0a, 0x1a, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c,
	0x76, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x73, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68,
	0x65, 0x6c, 0x76, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x07, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x07, 0x73, 0x68, 0x65, 0x6c,
	0x76, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4d, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x23, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x1c, 0x0a, 0x1a, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x68, 0x65, 0x6c, 0x66, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x13, 0x4d,
	0x65, 0x72, 0x67, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x23, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x1c, 0x0a, 0x1a, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x23, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x1c, 0x0a, 0x1a, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x68, 0x65, 0x6c,
	0x66, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67, 0x61, 0x69, 0x70, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qclaogui_library_v1_shelf_proto_rawDescOnce sync.Once
	file_qclaogui_library_v1_shelf_proto_rawDescData = file_qclaogui_library_v1_shelf_proto_rawDesc
)

func file_qclaogui_library_v1_shelf_proto_rawDescGZIP() []byte {
	file_qclaogui_library_v1_shelf_proto_rawDescOnce.Do(func() {
		file_qclaogui_library_v1_shelf_proto_rawDescData = protoimpl.X.CompressGZIP(file_qclaogui_library_v1_shelf_proto_rawDescData)
	})
	return file_qclaogui_library_v1_shelf_proto_rawDescData
}

var file_qclaogui_library_v1_shelf_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_qclaogui_library_v1_shelf_proto_goTypes = []any{
	(*Shelf)(nil),               // 0: qclaogui.library.v1.Shelf
	(*CreateShelfRequest)(nil),  // 1: qclaogui.library.v1.CreateShelfRequest
	(*GetShelfRequest)(nil),     // 2: qclaogui.library.v1.GetShelfRequest
	(*ListShelvesRequest)(nil),  // 3: qclaogui.library.v1.ListShelvesRequest
	(*ListShelvesResponse)(nil), // 4: qclaogui.library.v1.ListShelvesResponse
	(*DeleteShelfRequest)(nil),  // 5: qclaogui.library.v1.DeleteShelfRequest
	(*MergeShelvesRequest)(nil), // 6: qclaogui.library.v1.MergeShelvesRequest
}
var file_qclaogui_library_v1_shelf_proto_depIdxs = []int32{
	0, // 0: qclaogui.library.v1.CreateShelfRequest.shelf:type_name -> qclaogui.library.v1.Shelf
	0, // 1: qclaogui.library.v1.ListShelvesResponse.shelves:type_name -> qclaogui.library.v1.Shelf
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_qclaogui_library_v1_shelf_proto_init() }
func file_qclaogui_library_v1_shelf_proto_init() {
	if File_qclaogui_library_v1_shelf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_qclaogui_library_v1_shelf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qclaogui_library_v1_shelf_proto_goTypes,
		DependencyIndexes: file_qclaogui_library_v1_shelf_proto_depIdxs,
		MessageInfos:      file_qclaogui_library_v1_shelf_proto_msgTypes,
	}.Build()
	File_qclaogui_library_v1_shelf_proto = out.File
	file_qclaogui_library_v1_shelf_proto_rawDesc = nil
	file_qclaogui_library_v1_shelf_proto_goTypes = nil
	file_qclaogui_library_v1_shelf_proto_depIdxs = nil
}
