// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: qclaogui/library/v1/service.proto

package librarypb

import (
	reflect "reflect"

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

var File_qclaogui_library_v1_service_proto protoreflect.FileDescriptor

var file_qclaogui_library_v1_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x71, 0x63,
	0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x71, 0x63,
	0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc3, 0x0b,
	0x0a, 0x0e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x76, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12,
	0x27, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f,
	0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x68, 0x65, 0x6c, 0x66, 0x22, 0x22, 0xda, 0x41, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x3a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x0b, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x12, 0x71, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x65, 0x6c, 0x66, 0x12, 0x24, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x23, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x75, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x12, 0x27, 0x2e, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68,
	0x65, 0x6c, 0x76, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x76,
	0x65, 0x73, 0x12, 0x73, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c,
	0x66, 0x12, 0x27, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x23, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x2a, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x73, 0x68, 0x65,
	0x6c, 0x76, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x8e, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x72, 0x67,
	0x65, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f,
	0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x72, 0x67, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x38,
	0xda, 0x41, 0x10, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x68,
	0x65, 0x6c, 0x66, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x76,
	0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x2f,
	0x2a, 0x7d, 0x3a, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x26, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67,
	0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x38, 0xda, 0x41, 0x0b, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x62, 0x6f, 0x6f, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24,
	0x3a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x3d, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x76, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x23, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22,
	0x2b, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c,
	0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65,
	0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x89, 0x01, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x25, 0x2e, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0xda, 0x41, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x2f,
	0x2a, 0x7d, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x79, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x26, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2b, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x2f, 0x2a, 0x7d, 0x12, 0x90, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x26, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x3f, 0xda, 0x41, 0x10, 0x62, 0x6f, 0x6f, 0x6b, 0x2c, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26,
	0x3a, 0x01, 0x2a, 0x32, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x3d, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x91, 0x01, 0x0a, 0x08, 0x4d, 0x6f, 0x76, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x24, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x44, 0xda, 0x41, 0x15, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x6d, 0x6f, 0x76, 0x65, 0x1a, 0x17, 0xca, 0x41, 0x14, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67, 0x61, 0x69, 0x70, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_qclaogui_library_v1_service_proto_goTypes = []any{
	(*CreateShelfRequest)(nil),  // 0: qclaogui.library.v1.CreateShelfRequest
	(*GetShelfRequest)(nil),     // 1: qclaogui.library.v1.GetShelfRequest
	(*ListShelvesRequest)(nil),  // 2: qclaogui.library.v1.ListShelvesRequest
	(*DeleteShelfRequest)(nil),  // 3: qclaogui.library.v1.DeleteShelfRequest
	(*MergeShelvesRequest)(nil), // 4: qclaogui.library.v1.MergeShelvesRequest
	(*CreateBookRequest)(nil),   // 5: qclaogui.library.v1.CreateBookRequest
	(*GetBookRequest)(nil),      // 6: qclaogui.library.v1.GetBookRequest
	(*ListBooksRequest)(nil),    // 7: qclaogui.library.v1.ListBooksRequest
	(*DeleteBookRequest)(nil),   // 8: qclaogui.library.v1.DeleteBookRequest
	(*UpdateBookRequest)(nil),   // 9: qclaogui.library.v1.UpdateBookRequest
	(*MoveBookRequest)(nil),     // 10: qclaogui.library.v1.MoveBookRequest
	(*Shelf)(nil),               // 11: qclaogui.library.v1.Shelf
	(*ListShelvesResponse)(nil), // 12: qclaogui.library.v1.ListShelvesResponse
	(*emptypb.Empty)(nil),       // 13: google.protobuf.Empty
	(*Book)(nil),                // 14: qclaogui.library.v1.Book
	(*ListBooksResponse)(nil),   // 15: qclaogui.library.v1.ListBooksResponse
}
var file_qclaogui_library_v1_service_proto_depIdxs = []int32{
	0,  // 0: qclaogui.library.v1.LibraryService.CreateShelf:input_type -> qclaogui.library.v1.CreateShelfRequest
	1,  // 1: qclaogui.library.v1.LibraryService.GetShelf:input_type -> qclaogui.library.v1.GetShelfRequest
	2,  // 2: qclaogui.library.v1.LibraryService.ListShelves:input_type -> qclaogui.library.v1.ListShelvesRequest
	3,  // 3: qclaogui.library.v1.LibraryService.DeleteShelf:input_type -> qclaogui.library.v1.DeleteShelfRequest
	4,  // 4: qclaogui.library.v1.LibraryService.MergeShelves:input_type -> qclaogui.library.v1.MergeShelvesRequest
	5,  // 5: qclaogui.library.v1.LibraryService.CreateBook:input_type -> qclaogui.library.v1.CreateBookRequest
	6,  // 6: qclaogui.library.v1.LibraryService.GetBook:input_type -> qclaogui.library.v1.GetBookRequest
	7,  // 7: qclaogui.library.v1.LibraryService.ListBooks:input_type -> qclaogui.library.v1.ListBooksRequest
	8,  // 8: qclaogui.library.v1.LibraryService.DeleteBook:input_type -> qclaogui.library.v1.DeleteBookRequest
	9,  // 9: qclaogui.library.v1.LibraryService.UpdateBook:input_type -> qclaogui.library.v1.UpdateBookRequest
	10, // 10: qclaogui.library.v1.LibraryService.MoveBook:input_type -> qclaogui.library.v1.MoveBookRequest
	11, // 11: qclaogui.library.v1.LibraryService.CreateShelf:output_type -> qclaogui.library.v1.Shelf
	11, // 12: qclaogui.library.v1.LibraryService.GetShelf:output_type -> qclaogui.library.v1.Shelf
	12, // 13: qclaogui.library.v1.LibraryService.ListShelves:output_type -> qclaogui.library.v1.ListShelvesResponse
	13, // 14: qclaogui.library.v1.LibraryService.DeleteShelf:output_type -> google.protobuf.Empty
	11, // 15: qclaogui.library.v1.LibraryService.MergeShelves:output_type -> qclaogui.library.v1.Shelf
	14, // 16: qclaogui.library.v1.LibraryService.CreateBook:output_type -> qclaogui.library.v1.Book
	14, // 17: qclaogui.library.v1.LibraryService.GetBook:output_type -> qclaogui.library.v1.Book
	15, // 18: qclaogui.library.v1.LibraryService.ListBooks:output_type -> qclaogui.library.v1.ListBooksResponse
	13, // 19: qclaogui.library.v1.LibraryService.DeleteBook:output_type -> google.protobuf.Empty
	14, // 20: qclaogui.library.v1.LibraryService.UpdateBook:output_type -> qclaogui.library.v1.Book
	14, // 21: qclaogui.library.v1.LibraryService.MoveBook:output_type -> qclaogui.library.v1.Book
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_qclaogui_library_v1_service_proto_init() }
func file_qclaogui_library_v1_service_proto_init() {
	if File_qclaogui_library_v1_service_proto != nil {
		return
	}
	file_qclaogui_library_v1_book_proto_init()
	file_qclaogui_library_v1_shelf_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_qclaogui_library_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qclaogui_library_v1_service_proto_goTypes,
		DependencyIndexes: file_qclaogui_library_v1_service_proto_depIdxs,
	}.Build()
	File_qclaogui_library_v1_service_proto = out.File
	file_qclaogui_library_v1_service_proto_rawDesc = nil
	file_qclaogui_library_v1_service_proto_goTypes = nil
	file_qclaogui_library_v1_service_proto_depIdxs = nil
}
