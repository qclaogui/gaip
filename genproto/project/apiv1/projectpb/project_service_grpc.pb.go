// (-- api-linter: core::0191::java-package=disabled
// (-- api-linter: core::0191::java-multiple-files=disabled
// (-- api-linter: core::0191::java-outer-classname=disabled
//     aip.dev/not-precedent: We need to do this because reasons. --)

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: qclaogui/project/v1/project_service.proto

package projectpb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProjectService_CreateProject_FullMethodName = "/qclaogui.project.v1.ProjectService/CreateProject"
	ProjectService_GetProject_FullMethodName    = "/qclaogui.project.v1.ProjectService/GetProject"
	ProjectService_ListProjects_FullMethodName  = "/qclaogui.project.v1.ProjectService/ListProjects"
	ProjectService_DeleteProject_FullMethodName = "/qclaogui.project.v1.ProjectService/DeleteProject"
)

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	// Creates a new project.
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*Project, error)
	// Gets the specified project.
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error)
	// Lists projects.
	ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
	// Deletes the specified project.
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, ProjectService_CreateProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, ProjectService_GetProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, ProjectService_ListProjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProjectService_DeleteProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations should embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	// Creates a new project.
	CreateProject(context.Context, *CreateProjectRequest) (*Project, error)
	// Gets the specified project.
	GetProject(context.Context, *GetProjectRequest) (*Project, error)
	// Lists projects.
	ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error)
	// Deletes the specified project.
	DeleteProject(context.Context, *DeleteProjectRequest) (*emptypb.Empty, error)
}

// UnimplementedProjectServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) CreateProject(context.Context, *CreateProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectServiceServer) GetProject(context.Context, *GetProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedProjectServiceServer) ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
func (UnimplementedProjectServiceServer) DeleteProject(context.Context, *DeleteProjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_ListProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).ListProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_ListProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).ListProjects(ctx, req.(*ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_DeleteProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qclaogui.project.v1.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _ProjectService_CreateProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _ProjectService_GetProject_Handler,
		},
		{
			MethodName: "ListProjects",
			Handler:    _ProjectService_ListProjects_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _ProjectService_DeleteProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qclaogui/project/v1/project_service.proto",
}
