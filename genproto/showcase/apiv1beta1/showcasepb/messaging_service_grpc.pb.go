// (-- api-linter: core::0191::java-package=disabled
// (-- api-linter: core::0191::java-multiple-files=disabled
// (-- api-linter: core::0191::java-outer-classname=disabled
//     aip.dev/not-precedent: We need to do this because reasons. --)

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: qclaogui/showcase/v1beta1/messaging_service.proto

package showcasepb

import (
	context "context"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MessagingService_CreateRoom_FullMethodName   = "/qclaogui.showcase.v1beta1.MessagingService/CreateRoom"
	MessagingService_GetRoom_FullMethodName      = "/qclaogui.showcase.v1beta1.MessagingService/GetRoom"
	MessagingService_UpdateRoom_FullMethodName   = "/qclaogui.showcase.v1beta1.MessagingService/UpdateRoom"
	MessagingService_DeleteRoom_FullMethodName   = "/qclaogui.showcase.v1beta1.MessagingService/DeleteRoom"
	MessagingService_ListRooms_FullMethodName    = "/qclaogui.showcase.v1beta1.MessagingService/ListRooms"
	MessagingService_CreateBlurb_FullMethodName  = "/qclaogui.showcase.v1beta1.MessagingService/CreateBlurb"
	MessagingService_GetBlurb_FullMethodName     = "/qclaogui.showcase.v1beta1.MessagingService/GetBlurb"
	MessagingService_UpdateBlurb_FullMethodName  = "/qclaogui.showcase.v1beta1.MessagingService/UpdateBlurb"
	MessagingService_DeleteBlurb_FullMethodName  = "/qclaogui.showcase.v1beta1.MessagingService/DeleteBlurb"
	MessagingService_ListBlurbs_FullMethodName   = "/qclaogui.showcase.v1beta1.MessagingService/ListBlurbs"
	MessagingService_SearchBlurbs_FullMethodName = "/qclaogui.showcase.v1beta1.MessagingService/SearchBlurbs"
	MessagingService_StreamBlurbs_FullMethodName = "/qclaogui.showcase.v1beta1.MessagingService/StreamBlurbs"
	MessagingService_SendBlurbs_FullMethodName   = "/qclaogui.showcase.v1beta1.MessagingService/SendBlurbs"
	MessagingService_Connect_FullMethodName      = "/qclaogui.showcase.v1beta1.MessagingService/Connect"
)

// MessagingServiceClient is the client API for MessagingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A simple messaging service that implements chat rooms and profile posts.
//
// This messaging service showcases the features that API clients
// generated by gapic-generators implement.
type MessagingServiceClient interface {
	// Creates a room.
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*Room, error)
	// Retrieves the Room with the given resource name.
	GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*Room, error)
	// Updates a room.
	UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*Room, error)
	// Deletes a room and all of its blurbs.
	DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Lists all chat rooms.
	ListRooms(ctx context.Context, in *ListRoomsRequest, opts ...grpc.CallOption) (*ListRoomsResponse, error)
	// Creates a blurb. If the parent is a room, the blurb is understood to be a
	// message in that room. If the parent is a profile, the blurb is understood
	// to be a post on the profile.
	CreateBlurb(ctx context.Context, in *CreateBlurbRequest, opts ...grpc.CallOption) (*Blurb, error)
	// Retrieves the Blurb with the given resource name.
	GetBlurb(ctx context.Context, in *GetBlurbRequest, opts ...grpc.CallOption) (*Blurb, error)
	// Updates a blurb.
	UpdateBlurb(ctx context.Context, in *UpdateBlurbRequest, opts ...grpc.CallOption) (*Blurb, error)
	// Deletes a blurb.
	DeleteBlurb(ctx context.Context, in *DeleteBlurbRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Lists blurbs for a specific chat room or user profile depending on the
	// parent resource name.
	ListBlurbs(ctx context.Context, in *ListBlurbsRequest, opts ...grpc.CallOption) (*ListBlurbsResponse, error)
	// This method searches through all blurbs across all rooms and profiles
	// for blurbs containing to words found in the query. Only posts that
	// contain an exact match of a queried word will be returned.
	SearchBlurbs(ctx context.Context, in *SearchBlurbsRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// This returns a stream that emits the blurbs that are created for a
	// particular chat room or user profile.
	StreamBlurbs(ctx context.Context, in *StreamBlurbsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamBlurbsResponse], error)
	// This is a stream to create multiple blurbs. If an invalid blurb is
	// requested to be created, the stream will close with an error.
	SendBlurbs(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CreateBlurbRequest, SendBlurbsResponse], error)
	// This method starts a bidirectional stream that receives all blurbs that
	// are being created after the stream has started and sends requests to create
	// blurbs. If an invalid blurb is requested to be created, the stream will
	// close with an error.
	Connect(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ConnectRequest, StreamBlurbsResponse], error)
}

type messagingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingServiceClient(cc grpc.ClientConnInterface) MessagingServiceClient {
	return &messagingServiceClient{cc}
}

func (c *messagingServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*Room, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Room)
	err := c.cc.Invoke(ctx, MessagingService_CreateRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*Room, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Room)
	err := c.cc.Invoke(ctx, MessagingService_GetRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*Room, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Room)
	err := c.cc.Invoke(ctx, MessagingService_UpdateRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MessagingService_DeleteRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) ListRooms(ctx context.Context, in *ListRoomsRequest, opts ...grpc.CallOption) (*ListRoomsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRoomsResponse)
	err := c.cc.Invoke(ctx, MessagingService_ListRooms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) CreateBlurb(ctx context.Context, in *CreateBlurbRequest, opts ...grpc.CallOption) (*Blurb, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Blurb)
	err := c.cc.Invoke(ctx, MessagingService_CreateBlurb_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) GetBlurb(ctx context.Context, in *GetBlurbRequest, opts ...grpc.CallOption) (*Blurb, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Blurb)
	err := c.cc.Invoke(ctx, MessagingService_GetBlurb_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) UpdateBlurb(ctx context.Context, in *UpdateBlurbRequest, opts ...grpc.CallOption) (*Blurb, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Blurb)
	err := c.cc.Invoke(ctx, MessagingService_UpdateBlurb_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) DeleteBlurb(ctx context.Context, in *DeleteBlurbRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MessagingService_DeleteBlurb_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) ListBlurbs(ctx context.Context, in *ListBlurbsRequest, opts ...grpc.CallOption) (*ListBlurbsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBlurbsResponse)
	err := c.cc.Invoke(ctx, MessagingService_ListBlurbs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) SearchBlurbs(ctx context.Context, in *SearchBlurbsRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, MessagingService_SearchBlurbs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) StreamBlurbs(ctx context.Context, in *StreamBlurbsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamBlurbsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MessagingService_ServiceDesc.Streams[0], MessagingService_StreamBlurbs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamBlurbsRequest, StreamBlurbsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MessagingService_StreamBlurbsClient = grpc.ServerStreamingClient[StreamBlurbsResponse]

func (c *messagingServiceClient) SendBlurbs(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CreateBlurbRequest, SendBlurbsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MessagingService_ServiceDesc.Streams[1], MessagingService_SendBlurbs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CreateBlurbRequest, SendBlurbsResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MessagingService_SendBlurbsClient = grpc.ClientStreamingClient[CreateBlurbRequest, SendBlurbsResponse]

func (c *messagingServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ConnectRequest, StreamBlurbsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MessagingService_ServiceDesc.Streams[2], MessagingService_Connect_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ConnectRequest, StreamBlurbsResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MessagingService_ConnectClient = grpc.BidiStreamingClient[ConnectRequest, StreamBlurbsResponse]

// MessagingServiceServer is the server API for MessagingService service.
// All implementations should embed UnimplementedMessagingServiceServer
// for forward compatibility.
//
// A simple messaging service that implements chat rooms and profile posts.
//
// This messaging service showcases the features that API clients
// generated by gapic-generators implement.
type MessagingServiceServer interface {
	// Creates a room.
	CreateRoom(context.Context, *CreateRoomRequest) (*Room, error)
	// Retrieves the Room with the given resource name.
	GetRoom(context.Context, *GetRoomRequest) (*Room, error)
	// Updates a room.
	UpdateRoom(context.Context, *UpdateRoomRequest) (*Room, error)
	// Deletes a room and all of its blurbs.
	DeleteRoom(context.Context, *DeleteRoomRequest) (*emptypb.Empty, error)
	// Lists all chat rooms.
	ListRooms(context.Context, *ListRoomsRequest) (*ListRoomsResponse, error)
	// Creates a blurb. If the parent is a room, the blurb is understood to be a
	// message in that room. If the parent is a profile, the blurb is understood
	// to be a post on the profile.
	CreateBlurb(context.Context, *CreateBlurbRequest) (*Blurb, error)
	// Retrieves the Blurb with the given resource name.
	GetBlurb(context.Context, *GetBlurbRequest) (*Blurb, error)
	// Updates a blurb.
	UpdateBlurb(context.Context, *UpdateBlurbRequest) (*Blurb, error)
	// Deletes a blurb.
	DeleteBlurb(context.Context, *DeleteBlurbRequest) (*emptypb.Empty, error)
	// Lists blurbs for a specific chat room or user profile depending on the
	// parent resource name.
	ListBlurbs(context.Context, *ListBlurbsRequest) (*ListBlurbsResponse, error)
	// This method searches through all blurbs across all rooms and profiles
	// for blurbs containing to words found in the query. Only posts that
	// contain an exact match of a queried word will be returned.
	SearchBlurbs(context.Context, *SearchBlurbsRequest) (*longrunningpb.Operation, error)
	// This returns a stream that emits the blurbs that are created for a
	// particular chat room or user profile.
	StreamBlurbs(*StreamBlurbsRequest, grpc.ServerStreamingServer[StreamBlurbsResponse]) error
	// This is a stream to create multiple blurbs. If an invalid blurb is
	// requested to be created, the stream will close with an error.
	SendBlurbs(grpc.ClientStreamingServer[CreateBlurbRequest, SendBlurbsResponse]) error
	// This method starts a bidirectional stream that receives all blurbs that
	// are being created after the stream has started and sends requests to create
	// blurbs. If an invalid blurb is requested to be created, the stream will
	// close with an error.
	Connect(grpc.BidiStreamingServer[ConnectRequest, StreamBlurbsResponse]) error
}

// UnimplementedMessagingServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessagingServiceServer struct{}

func (UnimplementedMessagingServiceServer) CreateRoom(context.Context, *CreateRoomRequest) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}

func (UnimplementedMessagingServiceServer) GetRoom(context.Context, *GetRoomRequest) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}

func (UnimplementedMessagingServiceServer) UpdateRoom(context.Context, *UpdateRoomRequest) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoom not implemented")
}

func (UnimplementedMessagingServiceServer) DeleteRoom(context.Context, *DeleteRoomRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}

func (UnimplementedMessagingServiceServer) ListRooms(context.Context, *ListRoomsRequest) (*ListRoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRooms not implemented")
}

func (UnimplementedMessagingServiceServer) CreateBlurb(context.Context, *CreateBlurbRequest) (*Blurb, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlurb not implemented")
}

func (UnimplementedMessagingServiceServer) GetBlurb(context.Context, *GetBlurbRequest) (*Blurb, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlurb not implemented")
}

func (UnimplementedMessagingServiceServer) UpdateBlurb(context.Context, *UpdateBlurbRequest) (*Blurb, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlurb not implemented")
}

func (UnimplementedMessagingServiceServer) DeleteBlurb(context.Context, *DeleteBlurbRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlurb not implemented")
}

func (UnimplementedMessagingServiceServer) ListBlurbs(context.Context, *ListBlurbsRequest) (*ListBlurbsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBlurbs not implemented")
}

func (UnimplementedMessagingServiceServer) SearchBlurbs(context.Context, *SearchBlurbsRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBlurbs not implemented")
}

func (UnimplementedMessagingServiceServer) StreamBlurbs(*StreamBlurbsRequest, grpc.ServerStreamingServer[StreamBlurbsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamBlurbs not implemented")
}

func (UnimplementedMessagingServiceServer) SendBlurbs(grpc.ClientStreamingServer[CreateBlurbRequest, SendBlurbsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SendBlurbs not implemented")
}

func (UnimplementedMessagingServiceServer) Connect(grpc.BidiStreamingServer[ConnectRequest, StreamBlurbsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedMessagingServiceServer) testEmbeddedByValue() {}

// UnsafeMessagingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServiceServer will
// result in compilation errors.
type UnsafeMessagingServiceServer interface {
	mustEmbedUnimplementedMessagingServiceServer()
}

func RegisterMessagingServiceServer(s grpc.ServiceRegistrar, srv MessagingServiceServer) {
	// If the following call pancis, it indicates UnimplementedMessagingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MessagingService_ServiceDesc, srv)
}

func _MessagingService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_CreateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_GetRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).GetRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_GetRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).GetRoom(ctx, req.(*GetRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_UpdateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).UpdateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_UpdateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).UpdateRoom(ctx, req.(*UpdateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_DeleteRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).DeleteRoom(ctx, req.(*DeleteRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_ListRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).ListRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_ListRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).ListRooms(ctx, req.(*ListRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_CreateBlurb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlurbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).CreateBlurb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_CreateBlurb_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).CreateBlurb(ctx, req.(*CreateBlurbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_GetBlurb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlurbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).GetBlurb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_GetBlurb_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).GetBlurb(ctx, req.(*GetBlurbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_UpdateBlurb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlurbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).UpdateBlurb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_UpdateBlurb_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).UpdateBlurb(ctx, req.(*UpdateBlurbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_DeleteBlurb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlurbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).DeleteBlurb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_DeleteBlurb_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).DeleteBlurb(ctx, req.(*DeleteBlurbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_ListBlurbs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBlurbsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).ListBlurbs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_ListBlurbs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).ListBlurbs(ctx, req.(*ListBlurbsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_SearchBlurbs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBlurbsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).SearchBlurbs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_SearchBlurbs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).SearchBlurbs(ctx, req.(*SearchBlurbsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_StreamBlurbs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamBlurbsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagingServiceServer).StreamBlurbs(m, &grpc.GenericServerStream[StreamBlurbsRequest, StreamBlurbsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MessagingService_StreamBlurbsServer = grpc.ServerStreamingServer[StreamBlurbsResponse]

func _MessagingService_SendBlurbs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessagingServiceServer).SendBlurbs(&grpc.GenericServerStream[CreateBlurbRequest, SendBlurbsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MessagingService_SendBlurbsServer = grpc.ClientStreamingServer[CreateBlurbRequest, SendBlurbsResponse]

func _MessagingService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessagingServiceServer).Connect(&grpc.GenericServerStream[ConnectRequest, StreamBlurbsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MessagingService_ConnectServer = grpc.BidiStreamingServer[ConnectRequest, StreamBlurbsResponse]

// MessagingService_ServiceDesc is the grpc.ServiceDesc for MessagingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessagingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qclaogui.showcase.v1beta1.MessagingService",
	HandlerType: (*MessagingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _MessagingService_CreateRoom_Handler,
		},
		{
			MethodName: "GetRoom",
			Handler:    _MessagingService_GetRoom_Handler,
		},
		{
			MethodName: "UpdateRoom",
			Handler:    _MessagingService_UpdateRoom_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _MessagingService_DeleteRoom_Handler,
		},
		{
			MethodName: "ListRooms",
			Handler:    _MessagingService_ListRooms_Handler,
		},
		{
			MethodName: "CreateBlurb",
			Handler:    _MessagingService_CreateBlurb_Handler,
		},
		{
			MethodName: "GetBlurb",
			Handler:    _MessagingService_GetBlurb_Handler,
		},
		{
			MethodName: "UpdateBlurb",
			Handler:    _MessagingService_UpdateBlurb_Handler,
		},
		{
			MethodName: "DeleteBlurb",
			Handler:    _MessagingService_DeleteBlurb_Handler,
		},
		{
			MethodName: "ListBlurbs",
			Handler:    _MessagingService_ListBlurbs_Handler,
		},
		{
			MethodName: "SearchBlurbs",
			Handler:    _MessagingService_SearchBlurbs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamBlurbs",
			Handler:       _MessagingService_StreamBlurbs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendBlurbs",
			Handler:       _MessagingService_SendBlurbs_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Connect",
			Handler:       _MessagingService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "qclaogui/showcase/v1beta1/messaging_service.proto",
}
