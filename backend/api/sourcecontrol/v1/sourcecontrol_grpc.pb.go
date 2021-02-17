// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sourcecontrolv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SourceControlAPIClient is the client API for SourceControlAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SourceControlAPIClient interface {
	GetRepositoryOptions(ctx context.Context, in *GetRepositoryOptionsRequest, opts ...grpc.CallOption) (*GetRepositoryOptionsResponse, error)
	CreateRepository(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*CreateRepositoryResponse, error)
}

type sourceControlAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSourceControlAPIClient(cc grpc.ClientConnInterface) SourceControlAPIClient {
	return &sourceControlAPIClient{cc}
}

func (c *sourceControlAPIClient) GetRepositoryOptions(ctx context.Context, in *GetRepositoryOptionsRequest, opts ...grpc.CallOption) (*GetRepositoryOptionsResponse, error) {
	out := new(GetRepositoryOptionsResponse)
	err := c.cc.Invoke(ctx, "/clutch.sourcecontrol.v1.SourceControlAPI/GetRepositoryOptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceControlAPIClient) CreateRepository(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*CreateRepositoryResponse, error) {
	out := new(CreateRepositoryResponse)
	err := c.cc.Invoke(ctx, "/clutch.sourcecontrol.v1.SourceControlAPI/CreateRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SourceControlAPIServer is the server API for SourceControlAPI service.
// All implementations should embed UnimplementedSourceControlAPIServer
// for forward compatibility
type SourceControlAPIServer interface {
	GetRepositoryOptions(context.Context, *GetRepositoryOptionsRequest) (*GetRepositoryOptionsResponse, error)
	CreateRepository(context.Context, *CreateRepositoryRequest) (*CreateRepositoryResponse, error)
}

// UnimplementedSourceControlAPIServer should be embedded to have forward compatible implementations.
type UnimplementedSourceControlAPIServer struct {
}

func (UnimplementedSourceControlAPIServer) GetRepositoryOptions(context.Context, *GetRepositoryOptionsRequest) (*GetRepositoryOptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepositoryOptions not implemented")
}
func (UnimplementedSourceControlAPIServer) CreateRepository(context.Context, *CreateRepositoryRequest) (*CreateRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepository not implemented")
}

// UnsafeSourceControlAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SourceControlAPIServer will
// result in compilation errors.
type UnsafeSourceControlAPIServer interface {
	mustEmbedUnimplementedSourceControlAPIServer()
}

func RegisterSourceControlAPIServer(s grpc.ServiceRegistrar, srv SourceControlAPIServer) {
	s.RegisterService(&SourceControlAPI_ServiceDesc, srv)
}

func _SourceControlAPI_GetRepositoryOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepositoryOptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceControlAPIServer).GetRepositoryOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.sourcecontrol.v1.SourceControlAPI/GetRepositoryOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceControlAPIServer).GetRepositoryOptions(ctx, req.(*GetRepositoryOptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceControlAPI_CreateRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceControlAPIServer).CreateRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.sourcecontrol.v1.SourceControlAPI/CreateRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceControlAPIServer).CreateRepository(ctx, req.(*CreateRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SourceControlAPI_ServiceDesc is the grpc.ServiceDesc for SourceControlAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SourceControlAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.sourcecontrol.v1.SourceControlAPI",
	HandlerType: (*SourceControlAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRepositoryOptions",
			Handler:    _SourceControlAPI_GetRepositoryOptions_Handler,
		},
		{
			MethodName: "CreateRepository",
			Handler:    _SourceControlAPI_CreateRepository_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sourcecontrol/v1/sourcecontrol.proto",
}
