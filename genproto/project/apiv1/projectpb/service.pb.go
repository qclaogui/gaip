// (-- api-linter: core::0191::java-package=disabled
// (-- api-linter: core::0191::java-multiple-files=disabled
// (-- api-linter: core::0191::java-outer-classname=disabled
//     aip.dev/not-precedent: We need to do this because reasons. --)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: qclaogui/project/v1/service.proto

package projectpb

import (
	reflect "reflect"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
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

var File_qclaogui_project_v1_service_proto protoreflect.FileDescriptor

var file_qclaogui_project_v1_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x96,
	0x04, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x81, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x29, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x27, 0xda, 0x41,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x78, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x26, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x71, 0x63,
	0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x24, 0xda, 0x41, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x12,
	0x79, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12,
	0x28, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x78, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x29, 0x2e, 0x71, 0x63,
	0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x24,
	0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x7d, 0x1a, 0x11, 0xca, 0x41, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f,
	0x73, 0x74, 0x3a, 0x39, 0x30, 0x39, 0x35, 0x32, 0xda, 0x04, 0x0a, 0x0f, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x1d, 0xda, 0x41,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x08, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x6c, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x71, 0x63,
	0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x21, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x6d, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x25, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x76, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1f, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d,
	0x12, 0x6f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26,
	0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x21,
	0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a,
	0x7d, 0x1a, 0x11, 0xca, 0x41, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a,
	0x39, 0x30, 0x39, 0x35, 0x32, 0xac, 0x04, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x03, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x20, 0x2e,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xb8, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c,
	0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x65, 0x63, 0x68, 0x6f, 0x8a, 0xd3, 0xe4, 0x93,
	0x02, 0x9a, 0x02, 0x12, 0x08, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x7b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x2a, 0x7d, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x21, 0x7b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f,
	0x2a, 0x2f, 0x2a, 0x2a, 0x7d, 0x12, 0x22, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x18, 0x7b, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x2a, 0x2a, 0x12, 0x30, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x26, 0x7b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x2a, 0x2a, 0x7d, 0x12, 0x31, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x27, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x2a, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x3d, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x2a, 0x2a, 0x12, 0x18,
	0x0a, 0x0c, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x08,
	0x7b, 0x62, 0x61, 0x7a, 0x3d, 0x2a, 0x2a, 0x7d, 0x12, 0x23, 0x0a, 0x0c, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x13, 0x7b, 0x71, 0x75, 0x78, 0x3d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x2a, 0x2a, 0x12, 0x80, 0x01,
	0x0a, 0x04, 0x57, 0x61, 0x69, 0x74, 0x12, 0x20, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0xca, 0x41, 0x1c, 0x0a, 0x0c, 0x57, 0x61,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x57, 0x61, 0x69, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01,
	0x2a, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x77, 0x61, 0x69, 0x74,
	0x1a, 0x11, 0xca, 0x41, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x39,
	0x30, 0x39, 0x35, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_qclaogui_project_v1_service_proto_goTypes = []interface{}{
	(*CreateProjectRequest)(nil),    // 0: qclaogui.project.v1.CreateProjectRequest
	(*GetProjectRequest)(nil),       // 1: qclaogui.project.v1.GetProjectRequest
	(*ListProjectsRequest)(nil),     // 2: qclaogui.project.v1.ListProjectsRequest
	(*DeleteProjectRequest)(nil),    // 3: qclaogui.project.v1.DeleteProjectRequest
	(*CreateUserRequest)(nil),       // 4: qclaogui.project.v1.CreateUserRequest
	(*GetUserRequest)(nil),          // 5: qclaogui.project.v1.GetUserRequest
	(*ListUsersRequest)(nil),        // 6: qclaogui.project.v1.ListUsersRequest
	(*UpdateUserRequest)(nil),       // 7: qclaogui.project.v1.UpdateUserRequest
	(*DeleteUserRequest)(nil),       // 8: qclaogui.project.v1.DeleteUserRequest
	(*EchoRequest)(nil),             // 9: qclaogui.project.v1.EchoRequest
	(*WaitRequest)(nil),             // 10: qclaogui.project.v1.WaitRequest
	(*Project)(nil),                 // 11: qclaogui.project.v1.Project
	(*ListProjectsResponse)(nil),    // 12: qclaogui.project.v1.ListProjectsResponse
	(*emptypb.Empty)(nil),           // 13: google.protobuf.Empty
	(*User)(nil),                    // 14: qclaogui.project.v1.User
	(*ListUsersResponse)(nil),       // 15: qclaogui.project.v1.ListUsersResponse
	(*EchoResponse)(nil),            // 16: qclaogui.project.v1.EchoResponse
	(*longrunningpb.Operation)(nil), // 17: google.longrunning.Operation
}
var file_qclaogui_project_v1_service_proto_depIdxs = []int32{
	0,  // 0: qclaogui.project.v1.ProjectService.CreateProject:input_type -> qclaogui.project.v1.CreateProjectRequest
	1,  // 1: qclaogui.project.v1.ProjectService.GetProject:input_type -> qclaogui.project.v1.GetProjectRequest
	2,  // 2: qclaogui.project.v1.ProjectService.ListProjects:input_type -> qclaogui.project.v1.ListProjectsRequest
	3,  // 3: qclaogui.project.v1.ProjectService.DeleteProject:input_type -> qclaogui.project.v1.DeleteProjectRequest
	4,  // 4: qclaogui.project.v1.IdentityService.CreateUser:input_type -> qclaogui.project.v1.CreateUserRequest
	5,  // 5: qclaogui.project.v1.IdentityService.GetUser:input_type -> qclaogui.project.v1.GetUserRequest
	6,  // 6: qclaogui.project.v1.IdentityService.ListUsers:input_type -> qclaogui.project.v1.ListUsersRequest
	7,  // 7: qclaogui.project.v1.IdentityService.UpdateUser:input_type -> qclaogui.project.v1.UpdateUserRequest
	8,  // 8: qclaogui.project.v1.IdentityService.DeleteUser:input_type -> qclaogui.project.v1.DeleteUserRequest
	9,  // 9: qclaogui.project.v1.EchoService.Echo:input_type -> qclaogui.project.v1.EchoRequest
	10, // 10: qclaogui.project.v1.EchoService.Wait:input_type -> qclaogui.project.v1.WaitRequest
	11, // 11: qclaogui.project.v1.ProjectService.CreateProject:output_type -> qclaogui.project.v1.Project
	11, // 12: qclaogui.project.v1.ProjectService.GetProject:output_type -> qclaogui.project.v1.Project
	12, // 13: qclaogui.project.v1.ProjectService.ListProjects:output_type -> qclaogui.project.v1.ListProjectsResponse
	13, // 14: qclaogui.project.v1.ProjectService.DeleteProject:output_type -> google.protobuf.Empty
	14, // 15: qclaogui.project.v1.IdentityService.CreateUser:output_type -> qclaogui.project.v1.User
	14, // 16: qclaogui.project.v1.IdentityService.GetUser:output_type -> qclaogui.project.v1.User
	15, // 17: qclaogui.project.v1.IdentityService.ListUsers:output_type -> qclaogui.project.v1.ListUsersResponse
	14, // 18: qclaogui.project.v1.IdentityService.UpdateUser:output_type -> qclaogui.project.v1.User
	13, // 19: qclaogui.project.v1.IdentityService.DeleteUser:output_type -> google.protobuf.Empty
	16, // 20: qclaogui.project.v1.EchoService.Echo:output_type -> qclaogui.project.v1.EchoResponse
	17, // 21: qclaogui.project.v1.EchoService.Wait:output_type -> google.longrunning.Operation
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_qclaogui_project_v1_service_proto_init() }
func file_qclaogui_project_v1_service_proto_init() {
	if File_qclaogui_project_v1_service_proto != nil {
		return
	}
	file_qclaogui_project_v1_echo_proto_init()
	file_qclaogui_project_v1_project_proto_init()
	file_qclaogui_project_v1_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_qclaogui_project_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_qclaogui_project_v1_service_proto_goTypes,
		DependencyIndexes: file_qclaogui_project_v1_service_proto_depIdxs,
	}.Build()
	File_qclaogui_project_v1_service_proto = out.File
	file_qclaogui_project_v1_service_proto_rawDesc = nil
	file_qclaogui_project_v1_service_proto_goTypes = nil
	file_qclaogui_project_v1_service_proto_depIdxs = nil
}
