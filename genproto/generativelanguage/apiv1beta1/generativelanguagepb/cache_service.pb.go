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
// 	protoc        v6.30.1
// source: qclaogui/generativelanguage/v1beta1/cache_service.proto

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
	mi := &file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCachedContentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCachedContentsRequest) ProtoMessage() {}

func (x *ListCachedContentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes[0]
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
	return file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDescGZIP(), []int{0}
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
	mi := &file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCachedContentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCachedContentsResponse) ProtoMessage() {}

func (x *ListCachedContentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes[1]
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
	return file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDescGZIP(), []int{1}
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
	mi := &file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCachedContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCachedContentRequest) ProtoMessage() {}

func (x *CreateCachedContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes[2]
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
	return file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDescGZIP(), []int{2}
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
	mi := &file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCachedContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCachedContentRequest) ProtoMessage() {}

func (x *GetCachedContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes[3]
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
	return file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDescGZIP(), []int{3}
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
	mi := &file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCachedContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCachedContentRequest) ProtoMessage() {}

func (x *UpdateCachedContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes[4]
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
	return file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDescGZIP(), []int{4}
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
	mi := &file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCachedContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCachedContentRequest) ProtoMessage() {}

func (x *DeleteCachedContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes[5]
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
	return file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCachedContentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_qclaogui_generativelanguage_v1beta1_cache_service_proto protoreflect.FileDescriptor

var file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDesc = string([]byte{
	0x0a, 0x37, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x71, 0x63, 0x6c, 0x61, 0x6f,
	0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x38, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x19, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0xe2, 0x41, 0x01, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xa1, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b,
	0x0a, 0x0f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67,
	0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x7d, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x5f, 0x0a, 0x0e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x04, 0xe2,
	0x41, 0x01, 0x02, 0x52, 0x0d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x65, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xe2, 0x41, 0x01,
	0x02, 0xfa, 0x41, 0x2f, 0x0a, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x1a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5f, 0x0a, 0x0e, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x0d, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x68, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x36, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x2f, 0x0a, 0x2d, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x32, 0x85, 0x08, 0x0a, 0x0c, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xb8, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3e, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f,
	0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f,
	0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0xda, 0x41, 0x00, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0xcb, 0x01, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3f, 0xda, 0x41, 0x0e, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x28, 0x3a, 0x0e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0xb4, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x3c, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x2e, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21,
	0x12, 0x1f, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a,
	0x7d, 0x12, 0xef, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x63,
	0xda, 0x41, 0x1a, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x40, 0x3a, 0x0e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x32, 0x2e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x2a, 0x7d, 0x12, 0x9e, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x2e, 0x71, 0x63,
	0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x2e, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x21, 0x2a, 0x1f, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x2a, 0x7d, 0x1a, 0x22, 0xca, 0x41, 0x1f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f,
	0x67, 0x61, 0x69, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDescOnce sync.Once
	file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDescData []byte
)

func file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDescGZIP() []byte {
	file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDescOnce.Do(func() {
		file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDesc), len(file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDesc)))
	})
	return file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDescData
}

var (
	file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
	file_qclaogui_generativelanguage_v1beta1_cache_service_proto_goTypes  = []any{
		(*ListCachedContentsRequest)(nil),  // 0: qclaogui.generativelanguage.v1beta1.ListCachedContentsRequest
		(*ListCachedContentsResponse)(nil), // 1: qclaogui.generativelanguage.v1beta1.ListCachedContentsResponse
		(*CreateCachedContentRequest)(nil), // 2: qclaogui.generativelanguage.v1beta1.CreateCachedContentRequest
		(*GetCachedContentRequest)(nil),    // 3: qclaogui.generativelanguage.v1beta1.GetCachedContentRequest
		(*UpdateCachedContentRequest)(nil), // 4: qclaogui.generativelanguage.v1beta1.UpdateCachedContentRequest
		(*DeleteCachedContentRequest)(nil), // 5: qclaogui.generativelanguage.v1beta1.DeleteCachedContentRequest
		(*CachedContent)(nil),              // 6: qclaogui.generativelanguage.v1beta1.CachedContent
		(*fieldmaskpb.FieldMask)(nil),      // 7: google.protobuf.FieldMask
		(*emptypb.Empty)(nil),              // 8: google.protobuf.Empty
	}
)

var file_qclaogui_generativelanguage_v1beta1_cache_service_proto_depIdxs = []int32{
	6, // 0: qclaogui.generativelanguage.v1beta1.ListCachedContentsResponse.cached_contents:type_name -> qclaogui.generativelanguage.v1beta1.CachedContent
	6, // 1: qclaogui.generativelanguage.v1beta1.CreateCachedContentRequest.cached_content:type_name -> qclaogui.generativelanguage.v1beta1.CachedContent
	6, // 2: qclaogui.generativelanguage.v1beta1.UpdateCachedContentRequest.cached_content:type_name -> qclaogui.generativelanguage.v1beta1.CachedContent
	7, // 3: qclaogui.generativelanguage.v1beta1.UpdateCachedContentRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 4: qclaogui.generativelanguage.v1beta1.CacheService.ListCachedContents:input_type -> qclaogui.generativelanguage.v1beta1.ListCachedContentsRequest
	2, // 5: qclaogui.generativelanguage.v1beta1.CacheService.CreateCachedContent:input_type -> qclaogui.generativelanguage.v1beta1.CreateCachedContentRequest
	3, // 6: qclaogui.generativelanguage.v1beta1.CacheService.GetCachedContent:input_type -> qclaogui.generativelanguage.v1beta1.GetCachedContentRequest
	4, // 7: qclaogui.generativelanguage.v1beta1.CacheService.UpdateCachedContent:input_type -> qclaogui.generativelanguage.v1beta1.UpdateCachedContentRequest
	5, // 8: qclaogui.generativelanguage.v1beta1.CacheService.DeleteCachedContent:input_type -> qclaogui.generativelanguage.v1beta1.DeleteCachedContentRequest
	1, // 9: qclaogui.generativelanguage.v1beta1.CacheService.ListCachedContents:output_type -> qclaogui.generativelanguage.v1beta1.ListCachedContentsResponse
	6, // 10: qclaogui.generativelanguage.v1beta1.CacheService.CreateCachedContent:output_type -> qclaogui.generativelanguage.v1beta1.CachedContent
	6, // 11: qclaogui.generativelanguage.v1beta1.CacheService.GetCachedContent:output_type -> qclaogui.generativelanguage.v1beta1.CachedContent
	6, // 12: qclaogui.generativelanguage.v1beta1.CacheService.UpdateCachedContent:output_type -> qclaogui.generativelanguage.v1beta1.CachedContent
	8, // 13: qclaogui.generativelanguage.v1beta1.CacheService.DeleteCachedContent:output_type -> google.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_qclaogui_generativelanguage_v1beta1_cache_service_proto_init() }
func file_qclaogui_generativelanguage_v1beta1_cache_service_proto_init() {
	if File_qclaogui_generativelanguage_v1beta1_cache_service_proto != nil {
		return
	}
	file_qclaogui_generativelanguage_v1beta1_cached_content_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDesc), len(file_qclaogui_generativelanguage_v1beta1_cache_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qclaogui_generativelanguage_v1beta1_cache_service_proto_goTypes,
		DependencyIndexes: file_qclaogui_generativelanguage_v1beta1_cache_service_proto_depIdxs,
		MessageInfos:      file_qclaogui_generativelanguage_v1beta1_cache_service_proto_msgTypes,
	}.Build()
	File_qclaogui_generativelanguage_v1beta1_cache_service_proto = out.File
	file_qclaogui_generativelanguage_v1beta1_cache_service_proto_goTypes = nil
	file_qclaogui_generativelanguage_v1beta1_cache_service_proto_depIdxs = nil
}
