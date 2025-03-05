// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: qclaogui/generativelanguage/v1/content.proto

package generativelanguagepb

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

// The base structured datatype containing multi-part content of a message.
//
// A `Content` includes a `role` field designating the producer of the `Content`
// and a `parts` field containing multi-part data that contains the content of
// the message turn.
type Content struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Ordered `Parts` that constitute a single message. Parts may have different
	// MIME types.
	Parts []*Part `protobuf:"bytes,1,rep,name=parts,proto3" json:"parts,omitempty"`
	// Optional. The producer of the content. Must be either 'user' or 'model'.
	//
	// Useful to set for multi-turn conversations, otherwise can be left blank
	// or unset.
	Role          string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Content) Reset() {
	*x = Content{}
	mi := &file_qclaogui_generativelanguage_v1_content_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1_content_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1_content_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetParts() []*Part {
	if x != nil {
		return x.Parts
	}
	return nil
}

func (x *Content) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

// A datatype containing media that is part of a multi-part `Content` message.
//
// A `Part` consists of data which has an associated datatype. A `Part` can only
// contain one of the accepted types in `Part.data`.
//
// A `Part` must have a fixed IANA MIME type identifying the type and subtype
// of the media if the `inline_data` field is filled with raw bytes.
type Part struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*Part_Text
	//	*Part_InlineData
	Data          isPart_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Part) Reset() {
	*x = Part{}
	mi := &file_qclaogui_generativelanguage_v1_content_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Part) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Part) ProtoMessage() {}

func (x *Part) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1_content_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Part.ProtoReflect.Descriptor instead.
func (*Part) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1_content_proto_rawDescGZIP(), []int{1}
}

func (x *Part) GetData() isPart_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Part) GetText() string {
	if x != nil {
		if x, ok := x.Data.(*Part_Text); ok {
			return x.Text
		}
	}
	return ""
}

func (x *Part) GetInlineData() *Blob {
	if x != nil {
		if x, ok := x.Data.(*Part_InlineData); ok {
			return x.InlineData
		}
	}
	return nil
}

type isPart_Data interface {
	isPart_Data()
}

type Part_Text struct {
	// Inline text.
	Text string `protobuf:"bytes,2,opt,name=text,proto3,oneof"`
}

type Part_InlineData struct {
	// Inline media bytes.
	InlineData *Blob `protobuf:"bytes,3,opt,name=inline_data,json=inlineData,proto3,oneof"`
}

func (*Part_Text) isPart_Data() {}

func (*Part_InlineData) isPart_Data() {}

// Raw media bytes.
//
// Text should not be sent as raw bytes, use the 'text' field.
type Blob struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The IANA standard MIME type of the source data.
	// Examples:
	//   - image/png
	//   - image/jpeg
	//
	// If an unsupported MIME type is provided, an error will be returned. For a
	// complete list of supported types, see [Supported file
	// formats](https://ai.google.dev/gemini-api/docs/prompting_with_media#supported_file_formats).
	MimeType string `protobuf:"bytes,1,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Raw bytes for media formats.
	Data          []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Blob) Reset() {
	*x = Blob{}
	mi := &file_qclaogui_generativelanguage_v1_content_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Blob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blob) ProtoMessage() {}

func (x *Blob) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1_content_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blob.ProtoReflect.Descriptor instead.
func (*Blob) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1_content_proto_rawDescGZIP(), []int{2}
}

func (x *Blob) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *Blob) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_qclaogui_generativelanguage_v1_content_proto protoreflect.FileDescriptor

var file_qclaogui_generativelanguage_v1_content_proto_rawDesc = string([]byte{
	0x0a, 0x2c, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5f, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x05, 0x70, 0x61,
	0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x52,
	0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x22, 0x6d, 0x0a, 0x04, 0x50, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x47,
	0x0a, 0x0b, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x37, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f,
	0x67, 0x61, 0x69, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_qclaogui_generativelanguage_v1_content_proto_rawDescOnce sync.Once
	file_qclaogui_generativelanguage_v1_content_proto_rawDescData []byte
)

func file_qclaogui_generativelanguage_v1_content_proto_rawDescGZIP() []byte {
	file_qclaogui_generativelanguage_v1_content_proto_rawDescOnce.Do(func() {
		file_qclaogui_generativelanguage_v1_content_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_qclaogui_generativelanguage_v1_content_proto_rawDesc), len(file_qclaogui_generativelanguage_v1_content_proto_rawDesc)))
	})
	return file_qclaogui_generativelanguage_v1_content_proto_rawDescData
}

var (
	file_qclaogui_generativelanguage_v1_content_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_qclaogui_generativelanguage_v1_content_proto_goTypes  = []any{
		(*Content)(nil), // 0: qclaogui.generativelanguage.v1.Content
		(*Part)(nil),    // 1: qclaogui.generativelanguage.v1.Part
		(*Blob)(nil),    // 2: qclaogui.generativelanguage.v1.Blob
	}
)

var file_qclaogui_generativelanguage_v1_content_proto_depIdxs = []int32{
	1, // 0: qclaogui.generativelanguage.v1.Content.parts:type_name -> qclaogui.generativelanguage.v1.Part
	2, // 1: qclaogui.generativelanguage.v1.Part.inline_data:type_name -> qclaogui.generativelanguage.v1.Blob
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_qclaogui_generativelanguage_v1_content_proto_init() }
func file_qclaogui_generativelanguage_v1_content_proto_init() {
	if File_qclaogui_generativelanguage_v1_content_proto != nil {
		return
	}
	file_qclaogui_generativelanguage_v1_content_proto_msgTypes[1].OneofWrappers = []any{
		(*Part_Text)(nil),
		(*Part_InlineData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qclaogui_generativelanguage_v1_content_proto_rawDesc), len(file_qclaogui_generativelanguage_v1_content_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qclaogui_generativelanguage_v1_content_proto_goTypes,
		DependencyIndexes: file_qclaogui_generativelanguage_v1_content_proto_depIdxs,
		MessageInfos:      file_qclaogui_generativelanguage_v1_content_proto_msgTypes,
	}.Build()
	File_qclaogui_generativelanguage_v1_content_proto = out.File
	file_qclaogui_generativelanguage_v1_content_proto_goTypes = nil
	file_qclaogui_generativelanguage_v1_content_proto_depIdxs = nil
}
