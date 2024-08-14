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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: qclaogui/generativelanguage/v1beta1/discuss_service.proto

package generativelanguagepb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DiscussService_GenerateMessage_FullMethodName    = "/qclaogui.generativelanguage.v1beta1.DiscussService/GenerateMessage"
	DiscussService_CountMessageTokens_FullMethodName = "/qclaogui.generativelanguage.v1beta1.DiscussService/CountMessageTokens"
)

// DiscussServiceClient is the client API for DiscussService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// An API for using Generative Language Models (GLMs) in dialog applications.
//
// Also known as large language models (LLMs), this API provides models that
// are trained for multi-turn dialog.
type DiscussServiceClient interface {
	// Generates a response from the model given an input `MessagePrompt`.
	GenerateMessage(ctx context.Context, in *GenerateMessageRequest, opts ...grpc.CallOption) (*GenerateMessageResponse, error)
	// Runs a model's tokenizer on a string and returns the token count.
	CountMessageTokens(ctx context.Context, in *CountMessageTokensRequest, opts ...grpc.CallOption) (*CountMessageTokensResponse, error)
}

type discussServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscussServiceClient(cc grpc.ClientConnInterface) DiscussServiceClient {
	return &discussServiceClient{cc}
}

func (c *discussServiceClient) GenerateMessage(ctx context.Context, in *GenerateMessageRequest, opts ...grpc.CallOption) (*GenerateMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateMessageResponse)
	err := c.cc.Invoke(ctx, DiscussService_GenerateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discussServiceClient) CountMessageTokens(ctx context.Context, in *CountMessageTokensRequest, opts ...grpc.CallOption) (*CountMessageTokensResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountMessageTokensResponse)
	err := c.cc.Invoke(ctx, DiscussService_CountMessageTokens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscussServiceServer is the server API for DiscussService service.
// All implementations should embed UnimplementedDiscussServiceServer
// for forward compatibility.
//
// An API for using Generative Language Models (GLMs) in dialog applications.
//
// Also known as large language models (LLMs), this API provides models that
// are trained for multi-turn dialog.
type DiscussServiceServer interface {
	// Generates a response from the model given an input `MessagePrompt`.
	GenerateMessage(context.Context, *GenerateMessageRequest) (*GenerateMessageResponse, error)
	// Runs a model's tokenizer on a string and returns the token count.
	CountMessageTokens(context.Context, *CountMessageTokensRequest) (*CountMessageTokensResponse, error)
}

// UnimplementedDiscussServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDiscussServiceServer struct{}

func (UnimplementedDiscussServiceServer) GenerateMessage(context.Context, *GenerateMessageRequest) (*GenerateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateMessage not implemented")
}
func (UnimplementedDiscussServiceServer) CountMessageTokens(context.Context, *CountMessageTokensRequest) (*CountMessageTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountMessageTokens not implemented")
}
func (UnimplementedDiscussServiceServer) testEmbeddedByValue() {}

// UnsafeDiscussServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscussServiceServer will
// result in compilation errors.
type UnsafeDiscussServiceServer interface {
	mustEmbedUnimplementedDiscussServiceServer()
}

func RegisterDiscussServiceServer(s grpc.ServiceRegistrar, srv DiscussServiceServer) {
	// If the following call pancis, it indicates UnimplementedDiscussServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DiscussService_ServiceDesc, srv)
}

func _DiscussService_GenerateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscussServiceServer).GenerateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscussService_GenerateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscussServiceServer).GenerateMessage(ctx, req.(*GenerateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscussService_CountMessageTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountMessageTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscussServiceServer).CountMessageTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscussService_CountMessageTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscussServiceServer).CountMessageTokens(ctx, req.(*CountMessageTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscussService_ServiceDesc is the grpc.ServiceDesc for DiscussService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscussService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qclaogui.generativelanguage.v1beta1.DiscussService",
	HandlerType: (*DiscussServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateMessage",
			Handler:    _DiscussService_GenerateMessage_Handler,
		},
		{
			MethodName: "CountMessageTokens",
			Handler:    _DiscussService_CountMessageTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qclaogui/generativelanguage/v1beta1/discuss_service.proto",
}
