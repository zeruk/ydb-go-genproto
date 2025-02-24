// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: ydb_export_v1.proto

package Ydb_Export_V1

import (
	context "context"
	Ydb_Export "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Export"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ExportService_ExportToYt_FullMethodName = "/Ydb.Export.V1.ExportService/ExportToYt"
	ExportService_ExportToS3_FullMethodName = "/Ydb.Export.V1.ExportService/ExportToS3"
)

// ExportServiceClient is the client API for ExportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExportServiceClient interface {
	// Exports data to YT.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	ExportToYt(ctx context.Context, in *Ydb_Export.ExportToYtRequest, opts ...grpc.CallOption) (*Ydb_Export.ExportToYtResponse, error)
	// Exports data to S3.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	ExportToS3(ctx context.Context, in *Ydb_Export.ExportToS3Request, opts ...grpc.CallOption) (*Ydb_Export.ExportToS3Response, error)
}

type exportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExportServiceClient(cc grpc.ClientConnInterface) ExportServiceClient {
	return &exportServiceClient{cc}
}

func (c *exportServiceClient) ExportToYt(ctx context.Context, in *Ydb_Export.ExportToYtRequest, opts ...grpc.CallOption) (*Ydb_Export.ExportToYtResponse, error) {
	out := new(Ydb_Export.ExportToYtResponse)
	err := c.cc.Invoke(ctx, ExportService_ExportToYt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exportServiceClient) ExportToS3(ctx context.Context, in *Ydb_Export.ExportToS3Request, opts ...grpc.CallOption) (*Ydb_Export.ExportToS3Response, error) {
	out := new(Ydb_Export.ExportToS3Response)
	err := c.cc.Invoke(ctx, ExportService_ExportToS3_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExportServiceServer is the server API for ExportService service.
// All implementations must embed UnimplementedExportServiceServer
// for forward compatibility
type ExportServiceServer interface {
	// Exports data to YT.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	ExportToYt(context.Context, *Ydb_Export.ExportToYtRequest) (*Ydb_Export.ExportToYtResponse, error)
	// Exports data to S3.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	ExportToS3(context.Context, *Ydb_Export.ExportToS3Request) (*Ydb_Export.ExportToS3Response, error)
	mustEmbedUnimplementedExportServiceServer()
}

// UnimplementedExportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExportServiceServer struct {
}

func (UnimplementedExportServiceServer) ExportToYt(context.Context, *Ydb_Export.ExportToYtRequest) (*Ydb_Export.ExportToYtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportToYt not implemented")
}
func (UnimplementedExportServiceServer) ExportToS3(context.Context, *Ydb_Export.ExportToS3Request) (*Ydb_Export.ExportToS3Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportToS3 not implemented")
}
func (UnimplementedExportServiceServer) mustEmbedUnimplementedExportServiceServer() {}

// UnsafeExportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExportServiceServer will
// result in compilation errors.
type UnsafeExportServiceServer interface {
	mustEmbedUnimplementedExportServiceServer()
}

func RegisterExportServiceServer(s grpc.ServiceRegistrar, srv ExportServiceServer) {
	s.RegisterService(&ExportService_ServiceDesc, srv)
}

func _ExportService_ExportToYt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Export.ExportToYtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExportServiceServer).ExportToYt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExportService_ExportToYt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExportServiceServer).ExportToYt(ctx, req.(*Ydb_Export.ExportToYtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExportService_ExportToS3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Export.ExportToS3Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExportServiceServer).ExportToS3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExportService_ExportToS3_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExportServiceServer).ExportToS3(ctx, req.(*Ydb_Export.ExportToS3Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ExportService_ServiceDesc is the grpc.ServiceDesc for ExportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.Export.V1.ExportService",
	HandlerType: (*ExportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExportToYt",
			Handler:    _ExportService_ExportToYt_Handler,
		},
		{
			MethodName: "ExportToS3",
			Handler:    _ExportService_ExportToS3_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ydb_export_v1.proto",
}
