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
// - protoc             v5.29.1
// source: qclaogui/generativelanguage/v1beta1/permission_service.proto

package generativelanguagepb

import (
	context "context"

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
	PermissionService_CreatePermission_FullMethodName  = "/qclaogui.generativelanguage.v1beta1.PermissionService/CreatePermission"
	PermissionService_GetPermission_FullMethodName     = "/qclaogui.generativelanguage.v1beta1.PermissionService/GetPermission"
	PermissionService_ListPermissions_FullMethodName   = "/qclaogui.generativelanguage.v1beta1.PermissionService/ListPermissions"
	PermissionService_UpdatePermission_FullMethodName  = "/qclaogui.generativelanguage.v1beta1.PermissionService/UpdatePermission"
	PermissionService_DeletePermission_FullMethodName  = "/qclaogui.generativelanguage.v1beta1.PermissionService/DeletePermission"
	PermissionService_TransferOwnership_FullMethodName = "/qclaogui.generativelanguage.v1beta1.PermissionService/TransferOwnership"
)

// PermissionServiceClient is the client API for PermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Provides methods for managing permissions to PaLM API resources.
type PermissionServiceClient interface {
	// Create a permission to a specific resource.
	CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*Permission, error)
	// Gets information about a specific Permission.
	GetPermission(ctx context.Context, in *GetPermissionRequest, opts ...grpc.CallOption) (*Permission, error)
	// Lists permissions for the specific resource.
	ListPermissions(ctx context.Context, in *ListPermissionsRequest, opts ...grpc.CallOption) (*ListPermissionsResponse, error)
	// Updates the permission.
	UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*Permission, error)
	// Deletes the permission.
	DeletePermission(ctx context.Context, in *DeletePermissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Transfers ownership of the tuned model.
	// This is the only way to change ownership of the tuned model.
	// The current owner will be downgraded to writer role.
	TransferOwnership(ctx context.Context, in *TransferOwnershipRequest, opts ...grpc.CallOption) (*TransferOwnershipResponse, error)
}

type permissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionServiceClient(cc grpc.ClientConnInterface) PermissionServiceClient {
	return &permissionServiceClient{cc}
}

func (c *permissionServiceClient) CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*Permission, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Permission)
	err := c.cc.Invoke(ctx, PermissionService_CreatePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetPermission(ctx context.Context, in *GetPermissionRequest, opts ...grpc.CallOption) (*Permission, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Permission)
	err := c.cc.Invoke(ctx, PermissionService_GetPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) ListPermissions(ctx context.Context, in *ListPermissionsRequest, opts ...grpc.CallOption) (*ListPermissionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPermissionsResponse)
	err := c.cc.Invoke(ctx, PermissionService_ListPermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*Permission, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Permission)
	err := c.cc.Invoke(ctx, PermissionService_UpdatePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) DeletePermission(ctx context.Context, in *DeletePermissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PermissionService_DeletePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) TransferOwnership(ctx context.Context, in *TransferOwnershipRequest, opts ...grpc.CallOption) (*TransferOwnershipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransferOwnershipResponse)
	err := c.cc.Invoke(ctx, PermissionService_TransferOwnership_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServiceServer is the server API for PermissionService service.
// All implementations should embed UnimplementedPermissionServiceServer
// for forward compatibility.
//
// Provides methods for managing permissions to PaLM API resources.
type PermissionServiceServer interface {
	// Create a permission to a specific resource.
	CreatePermission(context.Context, *CreatePermissionRequest) (*Permission, error)
	// Gets information about a specific Permission.
	GetPermission(context.Context, *GetPermissionRequest) (*Permission, error)
	// Lists permissions for the specific resource.
	ListPermissions(context.Context, *ListPermissionsRequest) (*ListPermissionsResponse, error)
	// Updates the permission.
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*Permission, error)
	// Deletes the permission.
	DeletePermission(context.Context, *DeletePermissionRequest) (*emptypb.Empty, error)
	// Transfers ownership of the tuned model.
	// This is the only way to change ownership of the tuned model.
	// The current owner will be downgraded to writer role.
	TransferOwnership(context.Context, *TransferOwnershipRequest) (*TransferOwnershipResponse, error)
}

// UnimplementedPermissionServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPermissionServiceServer struct{}

func (UnimplementedPermissionServiceServer) CreatePermission(context.Context, *CreatePermissionRequest) (*Permission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermission not implemented")
}

func (UnimplementedPermissionServiceServer) GetPermission(context.Context, *GetPermissionRequest) (*Permission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermission not implemented")
}

func (UnimplementedPermissionServiceServer) ListPermissions(context.Context, *ListPermissionsRequest) (*ListPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPermissions not implemented")
}

func (UnimplementedPermissionServiceServer) UpdatePermission(context.Context, *UpdatePermissionRequest) (*Permission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePermission not implemented")
}

func (UnimplementedPermissionServiceServer) DeletePermission(context.Context, *DeletePermissionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePermission not implemented")
}

func (UnimplementedPermissionServiceServer) TransferOwnership(context.Context, *TransferOwnershipRequest) (*TransferOwnershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferOwnership not implemented")
}
func (UnimplementedPermissionServiceServer) testEmbeddedByValue() {}

// UnsafePermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionServiceServer will
// result in compilation errors.
type UnsafePermissionServiceServer interface {
	mustEmbedUnimplementedPermissionServiceServer()
}

func RegisterPermissionServiceServer(s grpc.ServiceRegistrar, srv PermissionServiceServer) {
	// If the following call pancis, it indicates UnimplementedPermissionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PermissionService_ServiceDesc, srv)
}

func _PermissionService_CreatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).CreatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_CreatePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).CreatePermission(ctx, req.(*CreatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_GetPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetPermission(ctx, req.(*GetPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_ListPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).ListPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_ListPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).ListPermissions(ctx, req.(*ListPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_UpdatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).UpdatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_UpdatePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).UpdatePermission(ctx, req.(*UpdatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_DeletePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).DeletePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_DeletePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).DeletePermission(ctx, req.(*DeletePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_TransferOwnership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferOwnershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).TransferOwnership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_TransferOwnership_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).TransferOwnership(ctx, req.(*TransferOwnershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionService_ServiceDesc is the grpc.ServiceDesc for PermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qclaogui.generativelanguage.v1beta1.PermissionService",
	HandlerType: (*PermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePermission",
			Handler:    _PermissionService_CreatePermission_Handler,
		},
		{
			MethodName: "GetPermission",
			Handler:    _PermissionService_GetPermission_Handler,
		},
		{
			MethodName: "ListPermissions",
			Handler:    _PermissionService_ListPermissions_Handler,
		},
		{
			MethodName: "UpdatePermission",
			Handler:    _PermissionService_UpdatePermission_Handler,
		},
		{
			MethodName: "DeletePermission",
			Handler:    _PermissionService_DeletePermission_Handler,
		},
		{
			MethodName: "TransferOwnership",
			Handler:    _PermissionService_TransferOwnership_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qclaogui/generativelanguage/v1beta1/permission_service.proto",
}