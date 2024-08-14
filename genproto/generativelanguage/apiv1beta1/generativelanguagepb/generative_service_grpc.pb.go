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
// source: qclaogui/generativelanguage/v1beta1/generative_service.proto

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
	GenerativeService_GenerateContent_FullMethodName       = "/qclaogui.generativelanguage.v1beta1.GenerativeService/GenerateContent"
	GenerativeService_GenerateAnswer_FullMethodName        = "/qclaogui.generativelanguage.v1beta1.GenerativeService/GenerateAnswer"
	GenerativeService_StreamGenerateContent_FullMethodName = "/qclaogui.generativelanguage.v1beta1.GenerativeService/StreamGenerateContent"
	GenerativeService_EmbedContent_FullMethodName          = "/qclaogui.generativelanguage.v1beta1.GenerativeService/EmbedContent"
	GenerativeService_BatchEmbedContents_FullMethodName    = "/qclaogui.generativelanguage.v1beta1.GenerativeService/BatchEmbedContents"
	GenerativeService_CountTokens_FullMethodName           = "/qclaogui.generativelanguage.v1beta1.GenerativeService/CountTokens"
)

// GenerativeServiceClient is the client API for GenerativeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// API for using Large Models that generate multimodal content and have
// additional capabilities beyond text generation.
type GenerativeServiceClient interface {
	// Generates a model response given an input `GenerateContentRequest`.
	// Refer to the [text generation
	// guide](https://ai.google.dev/gemini-api/docs/text-generation) for detailed
	// usage information. Input capabilities differ between models, including
	// tuned models. Refer to the [model
	// guide](https://ai.google.dev/gemini-api/docs/models/gemini) and [tuning
	// guide](https://ai.google.dev/gemini-api/docs/model-tuning) for details.
	GenerateContent(ctx context.Context, in *GenerateContentRequest, opts ...grpc.CallOption) (*GenerateContentResponse, error)
	// Generates a grounded answer from the model given an input
	// `GenerateAnswerRequest`.
	GenerateAnswer(ctx context.Context, in *GenerateAnswerRequest, opts ...grpc.CallOption) (*GenerateAnswerResponse, error)
	// Generates a [streamed
	// response](https://ai.google.dev/gemini-api/docs/text-generation?lang=python#generate-a-text-stream)
	// from the model given an input `GenerateContentRequest`.
	StreamGenerateContent(ctx context.Context, in *GenerateContentRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GenerateContentResponse], error)
	// Generates a text embedding vector from the input `Content` using the
	// specified [Gemini Embedding
	// model](https://ai.google.dev/gemini-api/docs/models/gemini#text-embedding).
	EmbedContent(ctx context.Context, in *EmbedContentRequest, opts ...grpc.CallOption) (*EmbedContentResponse, error)
	// Generates multiple embedding vectors from the input `Content` which
	// consists of a batch of strings represented as `EmbedContentRequest`
	// objects.
	BatchEmbedContents(ctx context.Context, in *BatchEmbedContentsRequest, opts ...grpc.CallOption) (*BatchEmbedContentsResponse, error)
	// Runs a model's tokenizer on input `Content` and returns the token count.
	// Refer to the [tokens guide](https://ai.google.dev/gemini-api/docs/tokens)
	// to learn more about tokens.
	CountTokens(ctx context.Context, in *CountTokensRequest, opts ...grpc.CallOption) (*CountTokensResponse, error)
}

type generativeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGenerativeServiceClient(cc grpc.ClientConnInterface) GenerativeServiceClient {
	return &generativeServiceClient{cc}
}

func (c *generativeServiceClient) GenerateContent(ctx context.Context, in *GenerateContentRequest, opts ...grpc.CallOption) (*GenerateContentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateContentResponse)
	err := c.cc.Invoke(ctx, GenerativeService_GenerateContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generativeServiceClient) GenerateAnswer(ctx context.Context, in *GenerateAnswerRequest, opts ...grpc.CallOption) (*GenerateAnswerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateAnswerResponse)
	err := c.cc.Invoke(ctx, GenerativeService_GenerateAnswer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generativeServiceClient) StreamGenerateContent(ctx context.Context, in *GenerateContentRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GenerateContentResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GenerativeService_ServiceDesc.Streams[0], GenerativeService_StreamGenerateContent_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GenerateContentRequest, GenerateContentResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GenerativeService_StreamGenerateContentClient = grpc.ServerStreamingClient[GenerateContentResponse]

func (c *generativeServiceClient) EmbedContent(ctx context.Context, in *EmbedContentRequest, opts ...grpc.CallOption) (*EmbedContentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmbedContentResponse)
	err := c.cc.Invoke(ctx, GenerativeService_EmbedContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generativeServiceClient) BatchEmbedContents(ctx context.Context, in *BatchEmbedContentsRequest, opts ...grpc.CallOption) (*BatchEmbedContentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchEmbedContentsResponse)
	err := c.cc.Invoke(ctx, GenerativeService_BatchEmbedContents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generativeServiceClient) CountTokens(ctx context.Context, in *CountTokensRequest, opts ...grpc.CallOption) (*CountTokensResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountTokensResponse)
	err := c.cc.Invoke(ctx, GenerativeService_CountTokens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenerativeServiceServer is the server API for GenerativeService service.
// All implementations should embed UnimplementedGenerativeServiceServer
// for forward compatibility.
//
// API for using Large Models that generate multimodal content and have
// additional capabilities beyond text generation.
type GenerativeServiceServer interface {
	// Generates a model response given an input `GenerateContentRequest`.
	// Refer to the [text generation
	// guide](https://ai.google.dev/gemini-api/docs/text-generation) for detailed
	// usage information. Input capabilities differ between models, including
	// tuned models. Refer to the [model
	// guide](https://ai.google.dev/gemini-api/docs/models/gemini) and [tuning
	// guide](https://ai.google.dev/gemini-api/docs/model-tuning) for details.
	GenerateContent(context.Context, *GenerateContentRequest) (*GenerateContentResponse, error)
	// Generates a grounded answer from the model given an input
	// `GenerateAnswerRequest`.
	GenerateAnswer(context.Context, *GenerateAnswerRequest) (*GenerateAnswerResponse, error)
	// Generates a [streamed
	// response](https://ai.google.dev/gemini-api/docs/text-generation?lang=python#generate-a-text-stream)
	// from the model given an input `GenerateContentRequest`.
	StreamGenerateContent(*GenerateContentRequest, grpc.ServerStreamingServer[GenerateContentResponse]) error
	// Generates a text embedding vector from the input `Content` using the
	// specified [Gemini Embedding
	// model](https://ai.google.dev/gemini-api/docs/models/gemini#text-embedding).
	EmbedContent(context.Context, *EmbedContentRequest) (*EmbedContentResponse, error)
	// Generates multiple embedding vectors from the input `Content` which
	// consists of a batch of strings represented as `EmbedContentRequest`
	// objects.
	BatchEmbedContents(context.Context, *BatchEmbedContentsRequest) (*BatchEmbedContentsResponse, error)
	// Runs a model's tokenizer on input `Content` and returns the token count.
	// Refer to the [tokens guide](https://ai.google.dev/gemini-api/docs/tokens)
	// to learn more about tokens.
	CountTokens(context.Context, *CountTokensRequest) (*CountTokensResponse, error)
}

// UnimplementedGenerativeServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGenerativeServiceServer struct{}

func (UnimplementedGenerativeServiceServer) GenerateContent(context.Context, *GenerateContentRequest) (*GenerateContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateContent not implemented")
}
func (UnimplementedGenerativeServiceServer) GenerateAnswer(context.Context, *GenerateAnswerRequest) (*GenerateAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAnswer not implemented")
}
func (UnimplementedGenerativeServiceServer) StreamGenerateContent(*GenerateContentRequest, grpc.ServerStreamingServer[GenerateContentResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamGenerateContent not implemented")
}
func (UnimplementedGenerativeServiceServer) EmbedContent(context.Context, *EmbedContentRequest) (*EmbedContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmbedContent not implemented")
}
func (UnimplementedGenerativeServiceServer) BatchEmbedContents(context.Context, *BatchEmbedContentsRequest) (*BatchEmbedContentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchEmbedContents not implemented")
}
func (UnimplementedGenerativeServiceServer) CountTokens(context.Context, *CountTokensRequest) (*CountTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountTokens not implemented")
}
func (UnimplementedGenerativeServiceServer) testEmbeddedByValue() {}

// UnsafeGenerativeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenerativeServiceServer will
// result in compilation errors.
type UnsafeGenerativeServiceServer interface {
	mustEmbedUnimplementedGenerativeServiceServer()
}

func RegisterGenerativeServiceServer(s grpc.ServiceRegistrar, srv GenerativeServiceServer) {
	// If the following call pancis, it indicates UnimplementedGenerativeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GenerativeService_ServiceDesc, srv)
}

func _GenerativeService_GenerateContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerativeServiceServer).GenerateContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenerativeService_GenerateContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerativeServiceServer).GenerateContent(ctx, req.(*GenerateContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenerativeService_GenerateAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerativeServiceServer).GenerateAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenerativeService_GenerateAnswer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerativeServiceServer).GenerateAnswer(ctx, req.(*GenerateAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenerativeService_StreamGenerateContent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GenerateContentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GenerativeServiceServer).StreamGenerateContent(m, &grpc.GenericServerStream[GenerateContentRequest, GenerateContentResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GenerativeService_StreamGenerateContentServer = grpc.ServerStreamingServer[GenerateContentResponse]

func _GenerativeService_EmbedContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmbedContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerativeServiceServer).EmbedContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenerativeService_EmbedContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerativeServiceServer).EmbedContent(ctx, req.(*EmbedContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenerativeService_BatchEmbedContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchEmbedContentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerativeServiceServer).BatchEmbedContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenerativeService_BatchEmbedContents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerativeServiceServer).BatchEmbedContents(ctx, req.(*BatchEmbedContentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenerativeService_CountTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerativeServiceServer).CountTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenerativeService_CountTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerativeServiceServer).CountTokens(ctx, req.(*CountTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GenerativeService_ServiceDesc is the grpc.ServiceDesc for GenerativeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenerativeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qclaogui.generativelanguage.v1beta1.GenerativeService",
	HandlerType: (*GenerativeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateContent",
			Handler:    _GenerativeService_GenerateContent_Handler,
		},
		{
			MethodName: "GenerateAnswer",
			Handler:    _GenerativeService_GenerateAnswer_Handler,
		},
		{
			MethodName: "EmbedContent",
			Handler:    _GenerativeService_EmbedContent_Handler,
		},
		{
			MethodName: "BatchEmbedContents",
			Handler:    _GenerativeService_BatchEmbedContents_Handler,
		},
		{
			MethodName: "CountTokens",
			Handler:    _GenerativeService_CountTokens_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamGenerateContent",
			Handler:       _GenerativeService_StreamGenerateContent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "qclaogui/generativelanguage/v1beta1/generative_service.proto",
}
