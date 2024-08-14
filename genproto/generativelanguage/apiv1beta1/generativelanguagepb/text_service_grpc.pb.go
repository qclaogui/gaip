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
// source: qclaogui/generativelanguage/v1beta1/text_service.proto

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
	TextService_GenerateText_FullMethodName    = "/qclaogui.generativelanguage.v1beta1.TextService/GenerateText"
	TextService_EmbedText_FullMethodName       = "/qclaogui.generativelanguage.v1beta1.TextService/EmbedText"
	TextService_BatchEmbedText_FullMethodName  = "/qclaogui.generativelanguage.v1beta1.TextService/BatchEmbedText"
	TextService_CountTextTokens_FullMethodName = "/qclaogui.generativelanguage.v1beta1.TextService/CountTextTokens"
)

// TextServiceClient is the client API for TextService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// API for using Generative Language Models (GLMs) trained to generate text.
//
// Also known as Large Language Models (LLM)s, these generate text given an
// input prompt from the user.
type TextServiceClient interface {
	// Generates a response from the model given an input message.
	GenerateText(ctx context.Context, in *GenerateTextRequest, opts ...grpc.CallOption) (*GenerateTextResponse, error)
	// Generates an embedding from the model given an input message.
	EmbedText(ctx context.Context, in *EmbedTextRequest, opts ...grpc.CallOption) (*EmbedTextResponse, error)
	// Generates multiple embeddings from the model given input text in a
	// synchronous call.
	BatchEmbedText(ctx context.Context, in *BatchEmbedTextRequest, opts ...grpc.CallOption) (*BatchEmbedTextResponse, error)
	// Runs a model's tokenizer on a text and returns the token count.
	CountTextTokens(ctx context.Context, in *CountTextTokensRequest, opts ...grpc.CallOption) (*CountTextTokensResponse, error)
}

type textServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTextServiceClient(cc grpc.ClientConnInterface) TextServiceClient {
	return &textServiceClient{cc}
}

func (c *textServiceClient) GenerateText(ctx context.Context, in *GenerateTextRequest, opts ...grpc.CallOption) (*GenerateTextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateTextResponse)
	err := c.cc.Invoke(ctx, TextService_GenerateText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textServiceClient) EmbedText(ctx context.Context, in *EmbedTextRequest, opts ...grpc.CallOption) (*EmbedTextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmbedTextResponse)
	err := c.cc.Invoke(ctx, TextService_EmbedText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textServiceClient) BatchEmbedText(ctx context.Context, in *BatchEmbedTextRequest, opts ...grpc.CallOption) (*BatchEmbedTextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchEmbedTextResponse)
	err := c.cc.Invoke(ctx, TextService_BatchEmbedText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textServiceClient) CountTextTokens(ctx context.Context, in *CountTextTokensRequest, opts ...grpc.CallOption) (*CountTextTokensResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountTextTokensResponse)
	err := c.cc.Invoke(ctx, TextService_CountTextTokens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextServiceServer is the server API for TextService service.
// All implementations should embed UnimplementedTextServiceServer
// for forward compatibility.
//
// API for using Generative Language Models (GLMs) trained to generate text.
//
// Also known as Large Language Models (LLM)s, these generate text given an
// input prompt from the user.
type TextServiceServer interface {
	// Generates a response from the model given an input message.
	GenerateText(context.Context, *GenerateTextRequest) (*GenerateTextResponse, error)
	// Generates an embedding from the model given an input message.
	EmbedText(context.Context, *EmbedTextRequest) (*EmbedTextResponse, error)
	// Generates multiple embeddings from the model given input text in a
	// synchronous call.
	BatchEmbedText(context.Context, *BatchEmbedTextRequest) (*BatchEmbedTextResponse, error)
	// Runs a model's tokenizer on a text and returns the token count.
	CountTextTokens(context.Context, *CountTextTokensRequest) (*CountTextTokensResponse, error)
}

// UnimplementedTextServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTextServiceServer struct{}

func (UnimplementedTextServiceServer) GenerateText(context.Context, *GenerateTextRequest) (*GenerateTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateText not implemented")
}
func (UnimplementedTextServiceServer) EmbedText(context.Context, *EmbedTextRequest) (*EmbedTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmbedText not implemented")
}
func (UnimplementedTextServiceServer) BatchEmbedText(context.Context, *BatchEmbedTextRequest) (*BatchEmbedTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchEmbedText not implemented")
}
func (UnimplementedTextServiceServer) CountTextTokens(context.Context, *CountTextTokensRequest) (*CountTextTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountTextTokens not implemented")
}
func (UnimplementedTextServiceServer) testEmbeddedByValue() {}

// UnsafeTextServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextServiceServer will
// result in compilation errors.
type UnsafeTextServiceServer interface {
	mustEmbedUnimplementedTextServiceServer()
}

func RegisterTextServiceServer(s grpc.ServiceRegistrar, srv TextServiceServer) {
	// If the following call pancis, it indicates UnimplementedTextServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TextService_ServiceDesc, srv)
}

func _TextService_GenerateText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextServiceServer).GenerateText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TextService_GenerateText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextServiceServer).GenerateText(ctx, req.(*GenerateTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextService_EmbedText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmbedTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextServiceServer).EmbedText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TextService_EmbedText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextServiceServer).EmbedText(ctx, req.(*EmbedTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextService_BatchEmbedText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchEmbedTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextServiceServer).BatchEmbedText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TextService_BatchEmbedText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextServiceServer).BatchEmbedText(ctx, req.(*BatchEmbedTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextService_CountTextTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountTextTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextServiceServer).CountTextTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TextService_CountTextTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextServiceServer).CountTextTokens(ctx, req.(*CountTextTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TextService_ServiceDesc is the grpc.ServiceDesc for TextService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qclaogui.generativelanguage.v1beta1.TextService",
	HandlerType: (*TextServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateText",
			Handler:    _TextService_GenerateText_Handler,
		},
		{
			MethodName: "EmbedText",
			Handler:    _TextService_EmbedText_Handler,
		},
		{
			MethodName: "BatchEmbedText",
			Handler:    _TextService_BatchEmbedText_Handler,
		},
		{
			MethodName: "CountTextTokens",
			Handler:    _TextService_CountTextTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qclaogui/generativelanguage/v1beta1/text_service.proto",
}
