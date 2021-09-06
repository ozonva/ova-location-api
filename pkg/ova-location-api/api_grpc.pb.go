// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ova_location_api

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

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	CreateLocationV1(ctx context.Context, in *CreateLocationV1Request, opts ...grpc.CallOption) (*LocationV1Response, error)
	GetLocationV1(ctx context.Context, in *GetLocationV1Request, opts ...grpc.CallOption) (*LocationV1Response, error)
	ListLocationsV1(ctx context.Context, in *ListLocationV1Request, opts ...grpc.CallOption) (*ListLocationsV1Response, error)
	RemoveLocationV1(ctx context.Context, in *RemoveLocationV1Request, opts ...grpc.CallOption) (*RemoveV1Response, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) CreateLocationV1(ctx context.Context, in *CreateLocationV1Request, opts ...grpc.CallOption) (*LocationV1Response, error) {
	out := new(LocationV1Response)
	err := c.cc.Invoke(ctx, "/api.api/CreateLocationV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetLocationV1(ctx context.Context, in *GetLocationV1Request, opts ...grpc.CallOption) (*LocationV1Response, error) {
	out := new(LocationV1Response)
	err := c.cc.Invoke(ctx, "/api.api/GetLocationV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListLocationsV1(ctx context.Context, in *ListLocationV1Request, opts ...grpc.CallOption) (*ListLocationsV1Response, error) {
	out := new(ListLocationsV1Response)
	err := c.cc.Invoke(ctx, "/api.api/ListLocationsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) RemoveLocationV1(ctx context.Context, in *RemoveLocationV1Request, opts ...grpc.CallOption) (*RemoveV1Response, error) {
	out := new(RemoveV1Response)
	err := c.cc.Invoke(ctx, "/api.api/RemoveLocationV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	CreateLocationV1(context.Context, *CreateLocationV1Request) (*LocationV1Response, error)
	GetLocationV1(context.Context, *GetLocationV1Request) (*LocationV1Response, error)
	ListLocationsV1(context.Context, *ListLocationV1Request) (*ListLocationsV1Response, error)
	RemoveLocationV1(context.Context, *RemoveLocationV1Request) (*RemoveV1Response, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) CreateLocationV1(context.Context, *CreateLocationV1Request) (*LocationV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLocationV1 not implemented")
}
func (UnimplementedApiServer) GetLocationV1(context.Context, *GetLocationV1Request) (*LocationV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocationV1 not implemented")
}
func (UnimplementedApiServer) ListLocationsV1(context.Context, *ListLocationV1Request) (*ListLocationsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLocationsV1 not implemented")
}
func (UnimplementedApiServer) RemoveLocationV1(context.Context, *RemoveLocationV1Request) (*RemoveV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLocationV1 not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_CreateLocationV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLocationV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).CreateLocationV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.api/CreateLocationV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).CreateLocationV1(ctx, req.(*CreateLocationV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetLocationV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetLocationV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.api/GetLocationV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetLocationV1(ctx, req.(*GetLocationV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ListLocationsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLocationV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ListLocationsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.api/ListLocationsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ListLocationsV1(ctx, req.(*ListLocationV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_RemoveLocationV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLocationV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RemoveLocationV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.api/RemoveLocationV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RemoveLocationV1(ctx, req.(*RemoveLocationV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLocationV1",
			Handler:    _Api_CreateLocationV1_Handler,
		},
		{
			MethodName: "GetLocationV1",
			Handler:    _Api_GetLocationV1_Handler,
		},
		{
			MethodName: "ListLocationsV1",
			Handler:    _Api_ListLocationsV1_Handler,
		},
		{
			MethodName: "RemoveLocationV1",
			Handler:    _Api_RemoveLocationV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
