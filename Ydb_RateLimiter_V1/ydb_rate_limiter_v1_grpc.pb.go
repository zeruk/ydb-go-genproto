// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Ydb_RateLimiter_V1

import (
	context "context"
	Ydb_RateLimiter "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_RateLimiter"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RateLimiterServiceClient is the client API for RateLimiterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RateLimiterServiceClient interface {
	// Create a new resource in existing coordination node.
	CreateResource(ctx context.Context, in *Ydb_RateLimiter.CreateResourceRequest, opts ...grpc.CallOption) (*Ydb_RateLimiter.CreateResourceResponse, error)
	// Update a resource in coordination node.
	AlterResource(ctx context.Context, in *Ydb_RateLimiter.AlterResourceRequest, opts ...grpc.CallOption) (*Ydb_RateLimiter.AlterResourceResponse, error)
	// Delete a resource from coordination node.
	DropResource(ctx context.Context, in *Ydb_RateLimiter.DropResourceRequest, opts ...grpc.CallOption) (*Ydb_RateLimiter.DropResourceResponse, error)
	// List resources in given coordination node.
	ListResources(ctx context.Context, in *Ydb_RateLimiter.ListResourcesRequest, opts ...grpc.CallOption) (*Ydb_RateLimiter.ListResourcesResponse, error)
	// Describe properties of resource in coordination node.
	DescribeResource(ctx context.Context, in *Ydb_RateLimiter.DescribeResourceRequest, opts ...grpc.CallOption) (*Ydb_RateLimiter.DescribeResourceResponse, error)
	// Take units for usage of a resource in coordination node.
	AcquireResource(ctx context.Context, in *Ydb_RateLimiter.AcquireResourceRequest, opts ...grpc.CallOption) (*Ydb_RateLimiter.AcquireResourceResponse, error)
}

type rateLimiterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRateLimiterServiceClient(cc grpc.ClientConnInterface) RateLimiterServiceClient {
	return &rateLimiterServiceClient{cc}
}

func (c *rateLimiterServiceClient) CreateResource(ctx context.Context, in *Ydb_RateLimiter.CreateResourceRequest, opts ...grpc.CallOption) (*Ydb_RateLimiter.CreateResourceResponse, error) {
	out := new(Ydb_RateLimiter.CreateResourceResponse)
	err := c.cc.Invoke(ctx, "/Ydb.RateLimiter.V1.RateLimiterService/CreateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateLimiterServiceClient) AlterResource(ctx context.Context, in *Ydb_RateLimiter.AlterResourceRequest, opts ...grpc.CallOption) (*Ydb_RateLimiter.AlterResourceResponse, error) {
	out := new(Ydb_RateLimiter.AlterResourceResponse)
	err := c.cc.Invoke(ctx, "/Ydb.RateLimiter.V1.RateLimiterService/AlterResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateLimiterServiceClient) DropResource(ctx context.Context, in *Ydb_RateLimiter.DropResourceRequest, opts ...grpc.CallOption) (*Ydb_RateLimiter.DropResourceResponse, error) {
	out := new(Ydb_RateLimiter.DropResourceResponse)
	err := c.cc.Invoke(ctx, "/Ydb.RateLimiter.V1.RateLimiterService/DropResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateLimiterServiceClient) ListResources(ctx context.Context, in *Ydb_RateLimiter.ListResourcesRequest, opts ...grpc.CallOption) (*Ydb_RateLimiter.ListResourcesResponse, error) {
	out := new(Ydb_RateLimiter.ListResourcesResponse)
	err := c.cc.Invoke(ctx, "/Ydb.RateLimiter.V1.RateLimiterService/ListResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateLimiterServiceClient) DescribeResource(ctx context.Context, in *Ydb_RateLimiter.DescribeResourceRequest, opts ...grpc.CallOption) (*Ydb_RateLimiter.DescribeResourceResponse, error) {
	out := new(Ydb_RateLimiter.DescribeResourceResponse)
	err := c.cc.Invoke(ctx, "/Ydb.RateLimiter.V1.RateLimiterService/DescribeResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateLimiterServiceClient) AcquireResource(ctx context.Context, in *Ydb_RateLimiter.AcquireResourceRequest, opts ...grpc.CallOption) (*Ydb_RateLimiter.AcquireResourceResponse, error) {
	out := new(Ydb_RateLimiter.AcquireResourceResponse)
	err := c.cc.Invoke(ctx, "/Ydb.RateLimiter.V1.RateLimiterService/AcquireResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimiterServiceServer is the server API for RateLimiterService service.
// All implementations must embed UnimplementedRateLimiterServiceServer
// for forward compatibility
type RateLimiterServiceServer interface {
	// Create a new resource in existing coordination node.
	CreateResource(context.Context, *Ydb_RateLimiter.CreateResourceRequest) (*Ydb_RateLimiter.CreateResourceResponse, error)
	// Update a resource in coordination node.
	AlterResource(context.Context, *Ydb_RateLimiter.AlterResourceRequest) (*Ydb_RateLimiter.AlterResourceResponse, error)
	// Delete a resource from coordination node.
	DropResource(context.Context, *Ydb_RateLimiter.DropResourceRequest) (*Ydb_RateLimiter.DropResourceResponse, error)
	// List resources in given coordination node.
	ListResources(context.Context, *Ydb_RateLimiter.ListResourcesRequest) (*Ydb_RateLimiter.ListResourcesResponse, error)
	// Describe properties of resource in coordination node.
	DescribeResource(context.Context, *Ydb_RateLimiter.DescribeResourceRequest) (*Ydb_RateLimiter.DescribeResourceResponse, error)
	// Take units for usage of a resource in coordination node.
	AcquireResource(context.Context, *Ydb_RateLimiter.AcquireResourceRequest) (*Ydb_RateLimiter.AcquireResourceResponse, error)
	mustEmbedUnimplementedRateLimiterServiceServer()
}

// UnimplementedRateLimiterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRateLimiterServiceServer struct {
}

func (UnimplementedRateLimiterServiceServer) CreateResource(context.Context, *Ydb_RateLimiter.CreateResourceRequest) (*Ydb_RateLimiter.CreateResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}
func (UnimplementedRateLimiterServiceServer) AlterResource(context.Context, *Ydb_RateLimiter.AlterResourceRequest) (*Ydb_RateLimiter.AlterResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterResource not implemented")
}
func (UnimplementedRateLimiterServiceServer) DropResource(context.Context, *Ydb_RateLimiter.DropResourceRequest) (*Ydb_RateLimiter.DropResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropResource not implemented")
}
func (UnimplementedRateLimiterServiceServer) ListResources(context.Context, *Ydb_RateLimiter.ListResourcesRequest) (*Ydb_RateLimiter.ListResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResources not implemented")
}
func (UnimplementedRateLimiterServiceServer) DescribeResource(context.Context, *Ydb_RateLimiter.DescribeResourceRequest) (*Ydb_RateLimiter.DescribeResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeResource not implemented")
}
func (UnimplementedRateLimiterServiceServer) AcquireResource(context.Context, *Ydb_RateLimiter.AcquireResourceRequest) (*Ydb_RateLimiter.AcquireResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcquireResource not implemented")
}
func (UnimplementedRateLimiterServiceServer) mustEmbedUnimplementedRateLimiterServiceServer() {}

// UnsafeRateLimiterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RateLimiterServiceServer will
// result in compilation errors.
type UnsafeRateLimiterServiceServer interface {
	mustEmbedUnimplementedRateLimiterServiceServer()
}

func RegisterRateLimiterServiceServer(s grpc.ServiceRegistrar, srv RateLimiterServiceServer) {
	s.RegisterService(&RateLimiterService_ServiceDesc, srv)
}

func _RateLimiterService_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_RateLimiter.CreateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimiterServiceServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.RateLimiter.V1.RateLimiterService/CreateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimiterServiceServer).CreateResource(ctx, req.(*Ydb_RateLimiter.CreateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateLimiterService_AlterResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_RateLimiter.AlterResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimiterServiceServer).AlterResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.RateLimiter.V1.RateLimiterService/AlterResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimiterServiceServer).AlterResource(ctx, req.(*Ydb_RateLimiter.AlterResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateLimiterService_DropResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_RateLimiter.DropResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimiterServiceServer).DropResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.RateLimiter.V1.RateLimiterService/DropResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimiterServiceServer).DropResource(ctx, req.(*Ydb_RateLimiter.DropResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateLimiterService_ListResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_RateLimiter.ListResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimiterServiceServer).ListResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.RateLimiter.V1.RateLimiterService/ListResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimiterServiceServer).ListResources(ctx, req.(*Ydb_RateLimiter.ListResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateLimiterService_DescribeResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_RateLimiter.DescribeResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimiterServiceServer).DescribeResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.RateLimiter.V1.RateLimiterService/DescribeResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimiterServiceServer).DescribeResource(ctx, req.(*Ydb_RateLimiter.DescribeResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateLimiterService_AcquireResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_RateLimiter.AcquireResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimiterServiceServer).AcquireResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.RateLimiter.V1.RateLimiterService/AcquireResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimiterServiceServer).AcquireResource(ctx, req.(*Ydb_RateLimiter.AcquireResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RateLimiterService_ServiceDesc is the grpc.ServiceDesc for RateLimiterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RateLimiterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.RateLimiter.V1.RateLimiterService",
	HandlerType: (*RateLimiterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResource",
			Handler:    _RateLimiterService_CreateResource_Handler,
		},
		{
			MethodName: "AlterResource",
			Handler:    _RateLimiterService_AlterResource_Handler,
		},
		{
			MethodName: "DropResource",
			Handler:    _RateLimiterService_DropResource_Handler,
		},
		{
			MethodName: "ListResources",
			Handler:    _RateLimiterService_ListResources_Handler,
		},
		{
			MethodName: "DescribeResource",
			Handler:    _RateLimiterService_DescribeResource_Handler,
		},
		{
			MethodName: "AcquireResource",
			Handler:    _RateLimiterService_AcquireResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ydb_rate_limiter_v1.proto",
}
