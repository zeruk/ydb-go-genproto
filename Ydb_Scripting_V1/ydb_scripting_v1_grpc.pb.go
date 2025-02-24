// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: ydb_scripting_v1.proto

package Ydb_Scripting_V1

import (
	context "context"
	Ydb_Scripting "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Scripting"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ScriptingService_ExecuteYql_FullMethodName       = "/Ydb.Scripting.V1.ScriptingService/ExecuteYql"
	ScriptingService_StreamExecuteYql_FullMethodName = "/Ydb.Scripting.V1.ScriptingService/StreamExecuteYql"
	ScriptingService_ExplainYql_FullMethodName       = "/Ydb.Scripting.V1.ScriptingService/ExplainYql"
)

// ScriptingServiceClient is the client API for ScriptingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScriptingServiceClient interface {
	ExecuteYql(ctx context.Context, in *Ydb_Scripting.ExecuteYqlRequest, opts ...grpc.CallOption) (*Ydb_Scripting.ExecuteYqlResponse, error)
	// Executes yql request with streaming result.
	StreamExecuteYql(ctx context.Context, in *Ydb_Scripting.ExecuteYqlRequest, opts ...grpc.CallOption) (ScriptingService_StreamExecuteYqlClient, error)
	ExplainYql(ctx context.Context, in *Ydb_Scripting.ExplainYqlRequest, opts ...grpc.CallOption) (*Ydb_Scripting.ExplainYqlResponse, error)
}

type scriptingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScriptingServiceClient(cc grpc.ClientConnInterface) ScriptingServiceClient {
	return &scriptingServiceClient{cc}
}

func (c *scriptingServiceClient) ExecuteYql(ctx context.Context, in *Ydb_Scripting.ExecuteYqlRequest, opts ...grpc.CallOption) (*Ydb_Scripting.ExecuteYqlResponse, error) {
	out := new(Ydb_Scripting.ExecuteYqlResponse)
	err := c.cc.Invoke(ctx, ScriptingService_ExecuteYql_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scriptingServiceClient) StreamExecuteYql(ctx context.Context, in *Ydb_Scripting.ExecuteYqlRequest, opts ...grpc.CallOption) (ScriptingService_StreamExecuteYqlClient, error) {
	stream, err := c.cc.NewStream(ctx, &ScriptingService_ServiceDesc.Streams[0], ScriptingService_StreamExecuteYql_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &scriptingServiceStreamExecuteYqlClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ScriptingService_StreamExecuteYqlClient interface {
	Recv() (*Ydb_Scripting.ExecuteYqlPartialResponse, error)
	grpc.ClientStream
}

type scriptingServiceStreamExecuteYqlClient struct {
	grpc.ClientStream
}

func (x *scriptingServiceStreamExecuteYqlClient) Recv() (*Ydb_Scripting.ExecuteYqlPartialResponse, error) {
	m := new(Ydb_Scripting.ExecuteYqlPartialResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scriptingServiceClient) ExplainYql(ctx context.Context, in *Ydb_Scripting.ExplainYqlRequest, opts ...grpc.CallOption) (*Ydb_Scripting.ExplainYqlResponse, error) {
	out := new(Ydb_Scripting.ExplainYqlResponse)
	err := c.cc.Invoke(ctx, ScriptingService_ExplainYql_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScriptingServiceServer is the server API for ScriptingService service.
// All implementations must embed UnimplementedScriptingServiceServer
// for forward compatibility
type ScriptingServiceServer interface {
	ExecuteYql(context.Context, *Ydb_Scripting.ExecuteYqlRequest) (*Ydb_Scripting.ExecuteYqlResponse, error)
	// Executes yql request with streaming result.
	StreamExecuteYql(*Ydb_Scripting.ExecuteYqlRequest, ScriptingService_StreamExecuteYqlServer) error
	ExplainYql(context.Context, *Ydb_Scripting.ExplainYqlRequest) (*Ydb_Scripting.ExplainYqlResponse, error)
	mustEmbedUnimplementedScriptingServiceServer()
}

// UnimplementedScriptingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScriptingServiceServer struct {
}

func (UnimplementedScriptingServiceServer) ExecuteYql(context.Context, *Ydb_Scripting.ExecuteYqlRequest) (*Ydb_Scripting.ExecuteYqlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteYql not implemented")
}
func (UnimplementedScriptingServiceServer) StreamExecuteYql(*Ydb_Scripting.ExecuteYqlRequest, ScriptingService_StreamExecuteYqlServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamExecuteYql not implemented")
}
func (UnimplementedScriptingServiceServer) ExplainYql(context.Context, *Ydb_Scripting.ExplainYqlRequest) (*Ydb_Scripting.ExplainYqlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExplainYql not implemented")
}
func (UnimplementedScriptingServiceServer) mustEmbedUnimplementedScriptingServiceServer() {}

// UnsafeScriptingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScriptingServiceServer will
// result in compilation errors.
type UnsafeScriptingServiceServer interface {
	mustEmbedUnimplementedScriptingServiceServer()
}

func RegisterScriptingServiceServer(s grpc.ServiceRegistrar, srv ScriptingServiceServer) {
	s.RegisterService(&ScriptingService_ServiceDesc, srv)
}

func _ScriptingService_ExecuteYql_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Scripting.ExecuteYqlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScriptingServiceServer).ExecuteYql(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScriptingService_ExecuteYql_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScriptingServiceServer).ExecuteYql(ctx, req.(*Ydb_Scripting.ExecuteYqlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScriptingService_StreamExecuteYql_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Ydb_Scripting.ExecuteYqlRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ScriptingServiceServer).StreamExecuteYql(m, &scriptingServiceStreamExecuteYqlServer{stream})
}

type ScriptingService_StreamExecuteYqlServer interface {
	Send(*Ydb_Scripting.ExecuteYqlPartialResponse) error
	grpc.ServerStream
}

type scriptingServiceStreamExecuteYqlServer struct {
	grpc.ServerStream
}

func (x *scriptingServiceStreamExecuteYqlServer) Send(m *Ydb_Scripting.ExecuteYqlPartialResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ScriptingService_ExplainYql_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Scripting.ExplainYqlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScriptingServiceServer).ExplainYql(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScriptingService_ExplainYql_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScriptingServiceServer).ExplainYql(ctx, req.(*Ydb_Scripting.ExplainYqlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScriptingService_ServiceDesc is the grpc.ServiceDesc for ScriptingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScriptingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.Scripting.V1.ScriptingService",
	HandlerType: (*ScriptingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteYql",
			Handler:    _ScriptingService_ExecuteYql_Handler,
		},
		{
			MethodName: "ExplainYql",
			Handler:    _ScriptingService_ExplainYql_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamExecuteYql",
			Handler:       _ScriptingService_StreamExecuteYql_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ydb_scripting_v1.proto",
}
