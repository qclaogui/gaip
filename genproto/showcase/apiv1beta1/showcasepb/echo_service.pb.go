// (-- api-linter: core::0191::java-package=disabled
// (-- api-linter: core::0191::java-multiple-files=disabled
// (-- api-linter: core::0191::java-outer-classname=disabled
//     aip.dev/not-precedent: We need to do this because reasons. --)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.1
// source: qclaogui/showcase/v1beta1/echo_service.proto

package showcasepb

import (
	reflect "reflect"
	unsafe "unsafe"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
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

var File_qclaogui_showcase_v1beta1_echo_service_proto protoreflect.FileDescriptor

var file_qclaogui_showcase_v1beta1_echo_service_proto_rawDesc = string([]byte{
	0x0a, 0x2c, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61,
	0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xff, 0x0a, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x98, 0x03, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x26,
	0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xbe, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x65, 0x63, 0x68, 0x6f, 0x8a,
	0xd3, 0xe4, 0x93, 0x02, 0x9a, 0x02, 0x12, 0x08, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x19, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x7b, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x2a, 0x7d, 0x12, 0x2b, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x7b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f, 0x6e,
	0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x2a, 0x2a, 0x7d, 0x12, 0x22, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x18, 0x7b, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x2a, 0x2a, 0x12, 0x30, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x26, 0x7b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x2a, 0x2a, 0x7d, 0x12, 0x31,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x27, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x3d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x2a,
	0x2a, 0x12, 0x18, 0x0a, 0x0c, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x08, 0x7b, 0x62, 0x61, 0x7a, 0x3d, 0x2a, 0x2a, 0x7d, 0x12, 0x23, 0x0a, 0x0c, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x13, 0x7b, 0x71, 0x75,
	0x78, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x2a, 0x2a,
	0x12, 0xa3, 0x01, 0x0a, 0x10, 0x45, 0x63, 0x68, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x32, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69,
	0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x8e, 0x01, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x61, 0x6e,
	0x64, 0x12, 0x28, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73, 0x68, 0x6f,
	0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x78,
	0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x71, 0x63,
	0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0xda, 0x41, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22,
	0x14, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x65,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x30, 0x01, 0x12, 0x7e, 0x0a, 0x07, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x12, 0x26, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73, 0x68,
	0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45,
	0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x28, 0x01, 0x12, 0x5b, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x26, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67,
	0x75, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x92, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x64, 0x45, 0x78,
	0x70, 0x61, 0x6e, 0x64, 0x12, 0x2d, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e,
	0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x64, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73,
	0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x64, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x12, 0x8b, 0x01, 0x0a, 0x04, 0x57, 0x61,
	0x69, 0x74, 0x12, 0x26, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73, 0x68,
	0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x57,
	0x61, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3c, 0xca, 0x41, 0x1c, 0x0a, 0x0c,
	0x57, 0x61, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x57, 0x61,
	0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x63,
	0x68, 0x6f, 0x3a, 0x77, 0x61, 0x69, 0x74, 0x12, 0x7a, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x27, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x1a, 0x22, 0xca, 0x41, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73,
	0x74, 0x3a, 0x39, 0x30, 0x39, 0x35, 0x8a, 0xd4, 0xdb, 0xd2, 0x0f, 0x0b, 0x76, 0x31, 0x5f, 0x32,
	0x30, 0x32, 0x34, 0x30, 0x35, 0x30, 0x36, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67,
	0x61, 0x69, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f,
	0x77, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var file_qclaogui_showcase_v1beta1_echo_service_proto_goTypes = []any{
	(*EchoRequest)(nil),              // 0: qclaogui.showcase.v1beta1.EchoRequest
	(*EchoErrorDetailsRequest)(nil),  // 1: qclaogui.showcase.v1beta1.EchoErrorDetailsRequest
	(*ExpandRequest)(nil),            // 2: qclaogui.showcase.v1beta1.ExpandRequest
	(*PagedExpandRequest)(nil),       // 3: qclaogui.showcase.v1beta1.PagedExpandRequest
	(*WaitRequest)(nil),              // 4: qclaogui.showcase.v1beta1.WaitRequest
	(*BlockRequest)(nil),             // 5: qclaogui.showcase.v1beta1.BlockRequest
	(*EchoResponse)(nil),             // 6: qclaogui.showcase.v1beta1.EchoResponse
	(*EchoErrorDetailsResponse)(nil), // 7: qclaogui.showcase.v1beta1.EchoErrorDetailsResponse
	(*PagedExpandResponse)(nil),      // 8: qclaogui.showcase.v1beta1.PagedExpandResponse
	(*longrunningpb.Operation)(nil),  // 9: google.longrunning.Operation
	(*BlockResponse)(nil),            // 10: qclaogui.showcase.v1beta1.BlockResponse
}

var file_qclaogui_showcase_v1beta1_echo_service_proto_depIdxs = []int32{
	0,  // 0: qclaogui.showcase.v1beta1.EchoService.Echo:input_type -> qclaogui.showcase.v1beta1.EchoRequest
	1,  // 1: qclaogui.showcase.v1beta1.EchoService.EchoErrorDetails:input_type -> qclaogui.showcase.v1beta1.EchoErrorDetailsRequest
	2,  // 2: qclaogui.showcase.v1beta1.EchoService.Expand:input_type -> qclaogui.showcase.v1beta1.ExpandRequest
	0,  // 3: qclaogui.showcase.v1beta1.EchoService.Collect:input_type -> qclaogui.showcase.v1beta1.EchoRequest
	0,  // 4: qclaogui.showcase.v1beta1.EchoService.Chat:input_type -> qclaogui.showcase.v1beta1.EchoRequest
	3,  // 5: qclaogui.showcase.v1beta1.EchoService.PagedExpand:input_type -> qclaogui.showcase.v1beta1.PagedExpandRequest
	4,  // 6: qclaogui.showcase.v1beta1.EchoService.Wait:input_type -> qclaogui.showcase.v1beta1.WaitRequest
	5,  // 7: qclaogui.showcase.v1beta1.EchoService.Block:input_type -> qclaogui.showcase.v1beta1.BlockRequest
	6,  // 8: qclaogui.showcase.v1beta1.EchoService.Echo:output_type -> qclaogui.showcase.v1beta1.EchoResponse
	7,  // 9: qclaogui.showcase.v1beta1.EchoService.EchoErrorDetails:output_type -> qclaogui.showcase.v1beta1.EchoErrorDetailsResponse
	6,  // 10: qclaogui.showcase.v1beta1.EchoService.Expand:output_type -> qclaogui.showcase.v1beta1.EchoResponse
	6,  // 11: qclaogui.showcase.v1beta1.EchoService.Collect:output_type -> qclaogui.showcase.v1beta1.EchoResponse
	6,  // 12: qclaogui.showcase.v1beta1.EchoService.Chat:output_type -> qclaogui.showcase.v1beta1.EchoResponse
	8,  // 13: qclaogui.showcase.v1beta1.EchoService.PagedExpand:output_type -> qclaogui.showcase.v1beta1.PagedExpandResponse
	9,  // 14: qclaogui.showcase.v1beta1.EchoService.Wait:output_type -> google.longrunning.Operation
	10, // 15: qclaogui.showcase.v1beta1.EchoService.Block:output_type -> qclaogui.showcase.v1beta1.BlockResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_qclaogui_showcase_v1beta1_echo_service_proto_init() }
func file_qclaogui_showcase_v1beta1_echo_service_proto_init() {
	if File_qclaogui_showcase_v1beta1_echo_service_proto != nil {
		return
	}
	file_qclaogui_showcase_v1beta1_echo_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qclaogui_showcase_v1beta1_echo_service_proto_rawDesc), len(file_qclaogui_showcase_v1beta1_echo_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qclaogui_showcase_v1beta1_echo_service_proto_goTypes,
		DependencyIndexes: file_qclaogui_showcase_v1beta1_echo_service_proto_depIdxs,
	}.Build()
	File_qclaogui_showcase_v1beta1_echo_service_proto = out.File
	file_qclaogui_showcase_v1beta1_echo_service_proto_goTypes = nil
	file_qclaogui_showcase_v1beta1_echo_service_proto_depIdxs = nil
}
