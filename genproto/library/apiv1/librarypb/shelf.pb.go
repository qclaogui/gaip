// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: qclaogui/library/v1/shelf.proto

package librarypb

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

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
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the shelf.
	// Shelf names have the form `shelves/{shelf_id}`.
	// The name is ignored when creating a shelf.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The theme of the shelf
	Theme         string `protobuf:"bytes,2,opt,name=theme,proto3" json:"theme,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// The shelf to create.
	Shelf         *Shelf `protobuf:"bytes,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the shelf to retrieve.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Requested page size. Server may return fewer shelves than requested.
	// If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A token identifying a page of results the server should return.
	// Typically, this is the value of
	// [ListShelvesResponse.next_page_token][qclaogui.library.v1.ListShelvesResponse.next_page_token]
	// returned from the previous call to `ListShelves` method.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// The list of shelves.
	Shelves []*Shelf `protobuf:"bytes,1,rep,name=shelves,proto3" json:"shelves,omitempty"`
	// A token to retrieve next page of results.
	// Pass this value in the
	// [ListShelvesRequest.page_token][qclaogui.library.v1.ListShelvesRequest.page_token]
	// field in the subsequent call to `ListShelves` method to retrieve the next
	// page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the shelf to delete.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the shelf we're adding books to.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OtherShelf    string `protobuf:"bytes,2,opt,name=other_shelf,json=otherShelf,proto3" json:"other_shelf,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

const file_qclaogui_library_v1_shelf_proto_rawDesc = "" +
	"\n" +
	"\x1fqclaogui/library/v1/shelf.proto\x12\x13qclaogui.library.v1\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"f\n" +
	"\x05Shelf\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05theme\x18\x02 \x01(\tR\x05theme:3\xeaA0\n" +
	"\x1alibrary.qclaogui.com/Shelf\x12\x12shelves/{shelf_id}\"L\n" +
	"\x12CreateShelfRequest\x126\n" +
	"\x05shelf\x18\x01 \x01(\v2\x1a.qclaogui.library.v1.ShelfB\x04\xe2A\x01\x02R\x05shelf\"J\n" +
	"\x0fGetShelfRequest\x127\n" +
	"\x04name\x18\x01 \x01(\tB#\xe2A\x01\x02\xfaA\x1c\n" +
	"\x1alibrary.qclaogui.com/ShelfR\x04name\"P\n" +
	"\x12ListShelvesRequest\x12\x1b\n" +
	"\tpage_size\x18\x01 \x01(\x05R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x02 \x01(\tR\tpageToken\"s\n" +
	"\x13ListShelvesResponse\x124\n" +
	"\ashelves\x18\x01 \x03(\v2\x1a.qclaogui.library.v1.ShelfR\ashelves\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"M\n" +
	"\x12DeleteShelfRequest\x127\n" +
	"\x04name\x18\x01 \x01(\tB#\xe2A\x01\x02\xfaA\x1c\n" +
	"\x1alibrary.qclaogui.com/ShelfR\x04name\"\x94\x01\n" +
	"\x13MergeShelvesRequest\x127\n" +
	"\x04name\x18\x01 \x01(\tB#\xe2A\x01\x02\xfaA\x1c\n" +
	"\x1alibrary.qclaogui.com/ShelfR\x04name\x12D\n" +
	"\vother_shelf\x18\x02 \x01(\tB#\xe2A\x01\x02\xfaA\x1c\n" +
	"\x1alibrary.qclaogui.com/ShelfR\n" +
	"otherShelfB;Z9github.com/qclaogui/gaip/genproto/library/apiv1/librarypbb\x06proto3"

var (
	file_qclaogui_library_v1_shelf_proto_rawDescOnce sync.Once
	file_qclaogui_library_v1_shelf_proto_rawDescData []byte
)

func file_qclaogui_library_v1_shelf_proto_rawDescGZIP() []byte {
	file_qclaogui_library_v1_shelf_proto_rawDescOnce.Do(func() {
		file_qclaogui_library_v1_shelf_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_qclaogui_library_v1_shelf_proto_rawDesc), len(file_qclaogui_library_v1_shelf_proto_rawDesc)))
	})
	return file_qclaogui_library_v1_shelf_proto_rawDescData
}

var (
	file_qclaogui_library_v1_shelf_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
	file_qclaogui_library_v1_shelf_proto_goTypes  = []any{
		(*Shelf)(nil),               // 0: qclaogui.library.v1.Shelf
		(*CreateShelfRequest)(nil),  // 1: qclaogui.library.v1.CreateShelfRequest
		(*GetShelfRequest)(nil),     // 2: qclaogui.library.v1.GetShelfRequest
		(*ListShelvesRequest)(nil),  // 3: qclaogui.library.v1.ListShelvesRequest
		(*ListShelvesResponse)(nil), // 4: qclaogui.library.v1.ListShelvesResponse
		(*DeleteShelfRequest)(nil),  // 5: qclaogui.library.v1.DeleteShelfRequest
		(*MergeShelvesRequest)(nil), // 6: qclaogui.library.v1.MergeShelvesRequest
	}
)

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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qclaogui_library_v1_shelf_proto_rawDesc), len(file_qclaogui_library_v1_shelf_proto_rawDesc)),
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
	file_qclaogui_library_v1_shelf_proto_goTypes = nil
	file_qclaogui_library_v1_shelf_proto_depIdxs = nil
}
