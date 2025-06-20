// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: qclaogui/generativelanguage/v1beta/cache_service.proto

package generativelanguagepb

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request to list CachedContents.
type ListCachedContentsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. The maximum number of cached contents to return. The service may
	// return fewer than this value. If unspecified, some default (under maximum)
	// number of items will be returned. The maximum value is 1000; values above
	// 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. A page token, received from a previous `ListCachedContents` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListCachedContents` must
	// match the call that provided the page token.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCachedContentsRequest) Reset() {
	*x = ListCachedContentsRequest{}
	mi := &file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCachedContentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCachedContentsRequest) ProtoMessage() {}

func (x *ListCachedContentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCachedContentsRequest.ProtoReflect.Descriptor instead.
func (*ListCachedContentsRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListCachedContentsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCachedContentsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response with CachedContents list.
type ListCachedContentsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of cached contents.
	CachedContents []*CachedContent `protobuf:"bytes,1,rep,name=cached_contents,json=cachedContents,proto3" json:"cached_contents,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCachedContentsResponse) Reset() {
	*x = ListCachedContentsResponse{}
	mi := &file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCachedContentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCachedContentsResponse) ProtoMessage() {}

func (x *ListCachedContentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCachedContentsResponse.ProtoReflect.Descriptor instead.
func (*ListCachedContentsResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListCachedContentsResponse) GetCachedContents() []*CachedContent {
	if x != nil {
		return x.CachedContents
	}
	return nil
}

func (x *ListCachedContentsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request to create CachedContent.
type CreateCachedContentRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The cached content to create.
	CachedContent *CachedContent `protobuf:"bytes,1,opt,name=cached_content,json=cachedContent,proto3" json:"cached_content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCachedContentRequest) Reset() {
	*x = CreateCachedContentRequest{}
	mi := &file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCachedContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCachedContentRequest) ProtoMessage() {}

func (x *CreateCachedContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCachedContentRequest.ProtoReflect.Descriptor instead.
func (*CreateCachedContentRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCachedContentRequest) GetCachedContent() *CachedContent {
	if x != nil {
		return x.CachedContent
	}
	return nil
}

// Request to read CachedContent.
type GetCachedContentRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The resource name referring to the content cache entry.
	// Format: `cachedContents/{id}`
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCachedContentRequest) Reset() {
	*x = GetCachedContentRequest{}
	mi := &file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCachedContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCachedContentRequest) ProtoMessage() {}

func (x *GetCachedContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCachedContentRequest.ProtoReflect.Descriptor instead.
func (*GetCachedContentRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCachedContentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request to update CachedContent.
type UpdateCachedContentRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The content cache entry to update
	CachedContent *CachedContent `protobuf:"bytes,1,opt,name=cached_content,json=cachedContent,proto3" json:"cached_content,omitempty"`
	// The list of fields to update.
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCachedContentRequest) Reset() {
	*x = UpdateCachedContentRequest{}
	mi := &file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCachedContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCachedContentRequest) ProtoMessage() {}

func (x *UpdateCachedContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCachedContentRequest.ProtoReflect.Descriptor instead.
func (*UpdateCachedContentRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCachedContentRequest) GetCachedContent() *CachedContent {
	if x != nil {
		return x.CachedContent
	}
	return nil
}

func (x *UpdateCachedContentRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Request to delete CachedContent.
type DeleteCachedContentRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The resource name referring to the content cache entry
	// Format: `cachedContents/{id}`
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCachedContentRequest) Reset() {
	*x = DeleteCachedContentRequest{}
	mi := &file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCachedContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCachedContentRequest) ProtoMessage() {}

func (x *DeleteCachedContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCachedContentRequest.ProtoReflect.Descriptor instead.
func (*DeleteCachedContentRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCachedContentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_qclaogui_generativelanguage_v1beta_cache_service_proto protoreflect.FileDescriptor

const file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDesc = "" +
	"\n" +
	"6qclaogui/generativelanguage/v1beta/cache_service.proto\x12\"qclaogui.generativelanguage.v1beta\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a google/protobuf/field_mask.proto\x1a7qclaogui/generativelanguage/v1beta/cached_content.proto\"c\n" +
	"\x19ListCachedContentsRequest\x12!\n" +
	"\tpage_size\x18\x01 \x01(\x05B\x04\xe2A\x01\x01R\bpageSize\x12#\n" +
	"\n" +
	"page_token\x18\x02 \x01(\tB\x04\xe2A\x01\x01R\tpageToken\"\xa0\x01\n" +
	"\x1aListCachedContentsResponse\x12Z\n" +
	"\x0fcached_contents\x18\x01 \x03(\v21.qclaogui.generativelanguage.v1beta.CachedContentR\x0ecachedContents\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"|\n" +
	"\x1aCreateCachedContentRequest\x12^\n" +
	"\x0ecached_content\x18\x01 \x01(\v21.qclaogui.generativelanguage.v1beta.CachedContentB\x04\xe2A\x01\x02R\rcachedContent\"e\n" +
	"\x17GetCachedContentRequest\x12J\n" +
	"\x04name\x18\x01 \x01(\tB6\xe2A\x01\x02\xfaA/\n" +
	"-generativelanguage.qclaogui.com/CachedContentR\x04name\"\xb9\x01\n" +
	"\x1aUpdateCachedContentRequest\x12^\n" +
	"\x0ecached_content\x18\x01 \x01(\v21.qclaogui.generativelanguage.v1beta.CachedContentB\x04\xe2A\x01\x02R\rcachedContent\x12;\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\"h\n" +
	"\x1aDeleteCachedContentRequest\x12J\n" +
	"\x04name\x18\x01 \x01(\tB6\xe2A\x01\x02\xfaA/\n" +
	"-generativelanguage.qclaogui.com/CachedContentR\x04name2\xfc\a\n" +
	"\fCacheService\x12\xb6\x01\n" +
	"\x12ListCachedContents\x12=.qclaogui.generativelanguage.v1beta.ListCachedContentsRequest\x1a>.qclaogui.generativelanguage.v1beta.ListCachedContentsResponse\"!\xdaA\x00\x82\xd3\xe4\x93\x02\x18\x12\x16/v1beta/cachedContents\x12\xc9\x01\n" +
	"\x13CreateCachedContent\x12>.qclaogui.generativelanguage.v1beta.CreateCachedContentRequest\x1a1.qclaogui.generativelanguage.v1beta.CachedContent\"?\xdaA\x0ecached_content\x82\xd3\xe4\x93\x02(:\x0ecached_content\"\x16/v1beta/cachedContents\x12\xb2\x01\n" +
	"\x10GetCachedContent\x12;.qclaogui.generativelanguage.v1beta.GetCachedContentRequest\x1a1.qclaogui.generativelanguage.v1beta.CachedContent\".\xdaA\x04name\x82\xd3\xe4\x93\x02!\x12\x1f/v1beta/{name=cachedContents/*}\x12\xed\x01\n" +
	"\x13UpdateCachedContent\x12>.qclaogui.generativelanguage.v1beta.UpdateCachedContentRequest\x1a1.qclaogui.generativelanguage.v1beta.CachedContent\"c\xdaA\x1acached_content,update_mask\x82\xd3\xe4\x93\x02@:\x0ecached_content2./v1beta/{cached_content.name=cachedContents/*}\x12\x9d\x01\n" +
	"\x13DeleteCachedContent\x12>.qclaogui.generativelanguage.v1beta.DeleteCachedContentRequest\x1a\x16.google.protobuf.Empty\".\xdaA\x04name\x82\xd3\xe4\x93\x02!*\x1f/v1beta/{name=cachedContents/*}\x1a\"\xcaA\x1fgenerativelanguage.qclaogui.comBUZSgithub.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta/generativelanguagepbb\x06proto3"

var (
	file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDescOnce sync.Once
	file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDescData []byte
)

func file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDescGZIP() []byte {
	file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDescOnce.Do(func() {
		file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDesc), len(file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDesc)))
	})
	return file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDescData
}

var (
	file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
	file_qclaogui_generativelanguage_v1beta_cache_service_proto_goTypes  = []any{
		(*ListCachedContentsRequest)(nil),  // 0: qclaogui.generativelanguage.v1beta.ListCachedContentsRequest
		(*ListCachedContentsResponse)(nil), // 1: qclaogui.generativelanguage.v1beta.ListCachedContentsResponse
		(*CreateCachedContentRequest)(nil), // 2: qclaogui.generativelanguage.v1beta.CreateCachedContentRequest
		(*GetCachedContentRequest)(nil),    // 3: qclaogui.generativelanguage.v1beta.GetCachedContentRequest
		(*UpdateCachedContentRequest)(nil), // 4: qclaogui.generativelanguage.v1beta.UpdateCachedContentRequest
		(*DeleteCachedContentRequest)(nil), // 5: qclaogui.generativelanguage.v1beta.DeleteCachedContentRequest
		(*CachedContent)(nil),              // 6: qclaogui.generativelanguage.v1beta.CachedContent
		(*fieldmaskpb.FieldMask)(nil),      // 7: google.protobuf.FieldMask
		(*emptypb.Empty)(nil),              // 8: google.protobuf.Empty
	}
)

var file_qclaogui_generativelanguage_v1beta_cache_service_proto_depIdxs = []int32{
	6, // 0: qclaogui.generativelanguage.v1beta.ListCachedContentsResponse.cached_contents:type_name -> qclaogui.generativelanguage.v1beta.CachedContent
	6, // 1: qclaogui.generativelanguage.v1beta.CreateCachedContentRequest.cached_content:type_name -> qclaogui.generativelanguage.v1beta.CachedContent
	6, // 2: qclaogui.generativelanguage.v1beta.UpdateCachedContentRequest.cached_content:type_name -> qclaogui.generativelanguage.v1beta.CachedContent
	7, // 3: qclaogui.generativelanguage.v1beta.UpdateCachedContentRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 4: qclaogui.generativelanguage.v1beta.CacheService.ListCachedContents:input_type -> qclaogui.generativelanguage.v1beta.ListCachedContentsRequest
	2, // 5: qclaogui.generativelanguage.v1beta.CacheService.CreateCachedContent:input_type -> qclaogui.generativelanguage.v1beta.CreateCachedContentRequest
	3, // 6: qclaogui.generativelanguage.v1beta.CacheService.GetCachedContent:input_type -> qclaogui.generativelanguage.v1beta.GetCachedContentRequest
	4, // 7: qclaogui.generativelanguage.v1beta.CacheService.UpdateCachedContent:input_type -> qclaogui.generativelanguage.v1beta.UpdateCachedContentRequest
	5, // 8: qclaogui.generativelanguage.v1beta.CacheService.DeleteCachedContent:input_type -> qclaogui.generativelanguage.v1beta.DeleteCachedContentRequest
	1, // 9: qclaogui.generativelanguage.v1beta.CacheService.ListCachedContents:output_type -> qclaogui.generativelanguage.v1beta.ListCachedContentsResponse
	6, // 10: qclaogui.generativelanguage.v1beta.CacheService.CreateCachedContent:output_type -> qclaogui.generativelanguage.v1beta.CachedContent
	6, // 11: qclaogui.generativelanguage.v1beta.CacheService.GetCachedContent:output_type -> qclaogui.generativelanguage.v1beta.CachedContent
	6, // 12: qclaogui.generativelanguage.v1beta.CacheService.UpdateCachedContent:output_type -> qclaogui.generativelanguage.v1beta.CachedContent
	8, // 13: qclaogui.generativelanguage.v1beta.CacheService.DeleteCachedContent:output_type -> google.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_qclaogui_generativelanguage_v1beta_cache_service_proto_init() }
func file_qclaogui_generativelanguage_v1beta_cache_service_proto_init() {
	if File_qclaogui_generativelanguage_v1beta_cache_service_proto != nil {
		return
	}
	file_qclaogui_generativelanguage_v1beta_cached_content_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDesc), len(file_qclaogui_generativelanguage_v1beta_cache_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qclaogui_generativelanguage_v1beta_cache_service_proto_goTypes,
		DependencyIndexes: file_qclaogui_generativelanguage_v1beta_cache_service_proto_depIdxs,
		MessageInfos:      file_qclaogui_generativelanguage_v1beta_cache_service_proto_msgTypes,
	}.Build()
	File_qclaogui_generativelanguage_v1beta_cache_service_proto = out.File
	file_qclaogui_generativelanguage_v1beta_cache_service_proto_goTypes = nil
	file_qclaogui_generativelanguage_v1beta_cache_service_proto_depIdxs = nil
}
