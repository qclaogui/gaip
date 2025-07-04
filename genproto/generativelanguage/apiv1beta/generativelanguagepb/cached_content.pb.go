// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: qclaogui/generativelanguage/v1beta/cached_content.proto

package generativelanguagepb

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Content that has been preprocessed and can be used in subsequent request
// to GenerativeService.
//
// Cached content can be only used with model it was created for.
type CachedContent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Specifies when this resource will expire.
	//
	// Types that are valid to be assigned to Expiration:
	//
	//	*CachedContent_ExpireTime
	//	*CachedContent_Ttl
	Expiration isCachedContent_Expiration `protobuf_oneof:"expiration"`
	// Optional. Identifier. The resource name referring to the cached content.
	// Format: `cachedContents/{id}`
	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Optional. Immutable. The user-generated meaningful display name of the
	// cached content. Maximum 128 Unicode characters.
	DisplayName *string `protobuf:"bytes,11,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
	// Required. Immutable. The name of the `Model` to use for cached content
	// Format: `models/{model}`
	Model *string `protobuf:"bytes,2,opt,name=model,proto3,oneof" json:"model,omitempty"`
	// Optional. Input only. Immutable. Developer set system instruction.
	// Currently text only.
	SystemInstruction *Content `protobuf:"bytes,3,opt,name=system_instruction,json=systemInstruction,proto3,oneof" json:"system_instruction,omitempty"`
	// Optional. Input only. Immutable. The content to cache.
	Contents []*Content `protobuf:"bytes,4,rep,name=contents,proto3" json:"contents,omitempty"`
	// Optional. Input only. Immutable. A list of `Tools` the model may use to
	// generate the next response
	Tools []*Tool `protobuf:"bytes,5,rep,name=tools,proto3" json:"tools,omitempty"`
	// Optional. Input only. Immutable. Tool config. This config is shared for all
	// tools.
	ToolConfig *ToolConfig `protobuf:"bytes,6,opt,name=tool_config,json=toolConfig,proto3,oneof" json:"tool_config,omitempty"`
	// Output only. Creation time of the cache entry.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. When the cache entry was last updated in UTC time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Metadata on the usage of the cached content.
	UsageMetadata *CachedContent_UsageMetadata `protobuf:"bytes,12,opt,name=usage_metadata,json=usageMetadata,proto3" json:"usage_metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CachedContent) Reset() {
	*x = CachedContent{}
	mi := &file_qclaogui_generativelanguage_v1beta_cached_content_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CachedContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CachedContent) ProtoMessage() {}

func (x *CachedContent) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta_cached_content_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CachedContent.ProtoReflect.Descriptor instead.
func (*CachedContent) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta_cached_content_proto_rawDescGZIP(), []int{0}
}

func (x *CachedContent) GetExpiration() isCachedContent_Expiration {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *CachedContent) GetExpireTime() *timestamppb.Timestamp {
	if x != nil {
		if x, ok := x.Expiration.(*CachedContent_ExpireTime); ok {
			return x.ExpireTime
		}
	}
	return nil
}

func (x *CachedContent) GetTtl() *durationpb.Duration {
	if x != nil {
		if x, ok := x.Expiration.(*CachedContent_Ttl); ok {
			return x.Ttl
		}
	}
	return nil
}

func (x *CachedContent) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CachedContent) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *CachedContent) GetModel() string {
	if x != nil && x.Model != nil {
		return *x.Model
	}
	return ""
}

func (x *CachedContent) GetSystemInstruction() *Content {
	if x != nil {
		return x.SystemInstruction
	}
	return nil
}

func (x *CachedContent) GetContents() []*Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *CachedContent) GetTools() []*Tool {
	if x != nil {
		return x.Tools
	}
	return nil
}

func (x *CachedContent) GetToolConfig() *ToolConfig {
	if x != nil {
		return x.ToolConfig
	}
	return nil
}

func (x *CachedContent) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *CachedContent) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *CachedContent) GetUsageMetadata() *CachedContent_UsageMetadata {
	if x != nil {
		return x.UsageMetadata
	}
	return nil
}

type isCachedContent_Expiration interface {
	isCachedContent_Expiration()
}

type CachedContent_ExpireTime struct {
	// Timestamp in UTC of when this resource is considered expired.
	// This is *always* provided on output, regardless of what was sent
	// on input.
	ExpireTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=expire_time,json=expireTime,proto3,oneof"`
}

type CachedContent_Ttl struct {
	// Input only. New TTL for this resource, input only.
	Ttl *durationpb.Duration `protobuf:"bytes,10,opt,name=ttl,proto3,oneof"`
}

func (*CachedContent_ExpireTime) isCachedContent_Expiration() {}

func (*CachedContent_Ttl) isCachedContent_Expiration() {}

// Metadata on the usage of the cached content.
type CachedContent_UsageMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Total number of tokens that the cached content consumes.
	TotalTokenCount int32 `protobuf:"varint,1,opt,name=total_token_count,json=totalTokenCount,proto3" json:"total_token_count,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CachedContent_UsageMetadata) Reset() {
	*x = CachedContent_UsageMetadata{}
	mi := &file_qclaogui_generativelanguage_v1beta_cached_content_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CachedContent_UsageMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CachedContent_UsageMetadata) ProtoMessage() {}

func (x *CachedContent_UsageMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta_cached_content_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CachedContent_UsageMetadata.ProtoReflect.Descriptor instead.
func (*CachedContent_UsageMetadata) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta_cached_content_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CachedContent_UsageMetadata) GetTotalTokenCount() int32 {
	if x != nil {
		return x.TotalTokenCount
	}
	return 0
}

var File_qclaogui_generativelanguage_v1beta_cached_content_proto protoreflect.FileDescriptor

const file_qclaogui_generativelanguage_v1beta_cached_content_proto_rawDesc = "" +
	"\n" +
	"7qclaogui/generativelanguage/v1beta/cached_content.proto\x12\"qclaogui.generativelanguage.v1beta\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a0qclaogui/generativelanguage/v1beta/content.proto\"\xf0\b\n" +
	"\rCachedContent\x12=\n" +
	"\vexpire_time\x18\t \x01(\v2\x1a.google.protobuf.TimestampH\x00R\n" +
	"expireTime\x123\n" +
	"\x03ttl\x18\n" +
	" \x01(\v2\x19.google.protobuf.DurationB\x04\xe2A\x01\x04H\x00R\x03ttl\x12\x1e\n" +
	"\x04name\x18\x01 \x01(\tB\x05\xe2A\x02\b\x01H\x01R\x04name\x88\x01\x01\x12-\n" +
	"\fdisplay_name\x18\v \x01(\tB\x05\xe2A\x02\x01\x05H\x02R\vdisplayName\x88\x01\x01\x12J\n" +
	"\x05model\x18\x02 \x01(\tB/\xe2A\x02\x05\x02\xfaA'\n" +
	"%generativelanguage.qclaogui.com/ModelH\x03R\x05model\x88\x01\x01\x12g\n" +
	"\x12system_instruction\x18\x03 \x01(\v2+.qclaogui.generativelanguage.v1beta.ContentB\x06\xe2A\x03\x01\x05\x04H\x04R\x11systemInstruction\x88\x01\x01\x12O\n" +
	"\bcontents\x18\x04 \x03(\v2+.qclaogui.generativelanguage.v1beta.ContentB\x06\xe2A\x03\x01\x05\x04R\bcontents\x12F\n" +
	"\x05tools\x18\x05 \x03(\v2(.qclaogui.generativelanguage.v1beta.ToolB\x06\xe2A\x03\x01\x05\x04R\x05tools\x12\\\n" +
	"\vtool_config\x18\x06 \x01(\v2..qclaogui.generativelanguage.v1beta.ToolConfigB\x06\xe2A\x03\x01\x05\x04H\x05R\n" +
	"toolConfig\x88\x01\x01\x12A\n" +
	"\vcreate_time\x18\a \x01(\v2\x1a.google.protobuf.TimestampB\x04\xe2A\x01\x03R\n" +
	"createTime\x12A\n" +
	"\vupdate_time\x18\b \x01(\v2\x1a.google.protobuf.TimestampB\x04\xe2A\x01\x03R\n" +
	"updateTime\x12l\n" +
	"\x0eusage_metadata\x18\f \x01(\v2?.qclaogui.generativelanguage.v1beta.CachedContent.UsageMetadataB\x04\xe2A\x01\x03R\rusageMetadata\x1a;\n" +
	"\rUsageMetadata\x12*\n" +
	"\x11total_token_count\x18\x01 \x01(\x05R\x0ftotalTokenCount:f\xeaAc\n" +
	"-generativelanguage.qclaogui.com/CachedContent\x12\x13cachedContents/{id}*\x0ecachedContents2\rcachedContentB\f\n" +
	"\n" +
	"expirationB\a\n" +
	"\x05_nameB\x0f\n" +
	"\r_display_nameB\b\n" +
	"\x06_modelB\x15\n" +
	"\x13_system_instructionB\x0e\n" +
	"\f_tool_configBUZSgithub.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta/generativelanguagepbb\x06proto3"

var (
	file_qclaogui_generativelanguage_v1beta_cached_content_proto_rawDescOnce sync.Once
	file_qclaogui_generativelanguage_v1beta_cached_content_proto_rawDescData []byte
)

func file_qclaogui_generativelanguage_v1beta_cached_content_proto_rawDescGZIP() []byte {
	file_qclaogui_generativelanguage_v1beta_cached_content_proto_rawDescOnce.Do(func() {
		file_qclaogui_generativelanguage_v1beta_cached_content_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_qclaogui_generativelanguage_v1beta_cached_content_proto_rawDesc), len(file_qclaogui_generativelanguage_v1beta_cached_content_proto_rawDesc)))
	})
	return file_qclaogui_generativelanguage_v1beta_cached_content_proto_rawDescData
}

var (
	file_qclaogui_generativelanguage_v1beta_cached_content_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_qclaogui_generativelanguage_v1beta_cached_content_proto_goTypes  = []any{
		(*CachedContent)(nil),               // 0: qclaogui.generativelanguage.v1beta.CachedContent
		(*CachedContent_UsageMetadata)(nil), // 1: qclaogui.generativelanguage.v1beta.CachedContent.UsageMetadata
		(*timestamppb.Timestamp)(nil),       // 2: google.protobuf.Timestamp
		(*durationpb.Duration)(nil),         // 3: google.protobuf.Duration
		(*Content)(nil),                     // 4: qclaogui.generativelanguage.v1beta.Content
		(*Tool)(nil),                        // 5: qclaogui.generativelanguage.v1beta.Tool
		(*ToolConfig)(nil),                  // 6: qclaogui.generativelanguage.v1beta.ToolConfig
	}
)

var file_qclaogui_generativelanguage_v1beta_cached_content_proto_depIdxs = []int32{
	2, // 0: qclaogui.generativelanguage.v1beta.CachedContent.expire_time:type_name -> google.protobuf.Timestamp
	3, // 1: qclaogui.generativelanguage.v1beta.CachedContent.ttl:type_name -> google.protobuf.Duration
	4, // 2: qclaogui.generativelanguage.v1beta.CachedContent.system_instruction:type_name -> qclaogui.generativelanguage.v1beta.Content
	4, // 3: qclaogui.generativelanguage.v1beta.CachedContent.contents:type_name -> qclaogui.generativelanguage.v1beta.Content
	5, // 4: qclaogui.generativelanguage.v1beta.CachedContent.tools:type_name -> qclaogui.generativelanguage.v1beta.Tool
	6, // 5: qclaogui.generativelanguage.v1beta.CachedContent.tool_config:type_name -> qclaogui.generativelanguage.v1beta.ToolConfig
	2, // 6: qclaogui.generativelanguage.v1beta.CachedContent.create_time:type_name -> google.protobuf.Timestamp
	2, // 7: qclaogui.generativelanguage.v1beta.CachedContent.update_time:type_name -> google.protobuf.Timestamp
	1, // 8: qclaogui.generativelanguage.v1beta.CachedContent.usage_metadata:type_name -> qclaogui.generativelanguage.v1beta.CachedContent.UsageMetadata
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_qclaogui_generativelanguage_v1beta_cached_content_proto_init() }
func file_qclaogui_generativelanguage_v1beta_cached_content_proto_init() {
	if File_qclaogui_generativelanguage_v1beta_cached_content_proto != nil {
		return
	}
	file_qclaogui_generativelanguage_v1beta_content_proto_init()
	file_qclaogui_generativelanguage_v1beta_cached_content_proto_msgTypes[0].OneofWrappers = []any{
		(*CachedContent_ExpireTime)(nil),
		(*CachedContent_Ttl)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qclaogui_generativelanguage_v1beta_cached_content_proto_rawDesc), len(file_qclaogui_generativelanguage_v1beta_cached_content_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qclaogui_generativelanguage_v1beta_cached_content_proto_goTypes,
		DependencyIndexes: file_qclaogui_generativelanguage_v1beta_cached_content_proto_depIdxs,
		MessageInfos:      file_qclaogui_generativelanguage_v1beta_cached_content_proto_msgTypes,
	}.Build()
	File_qclaogui_generativelanguage_v1beta_cached_content_proto = out.File
	file_qclaogui_generativelanguage_v1beta_cached_content_proto_goTypes = nil
	file_qclaogui_generativelanguage_v1beta_cached_content_proto_depIdxs = nil
}
