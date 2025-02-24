// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: ydb_coordination_v1.proto

package Ydb_Coordination_V1

import (
	context "context"
	Ydb_Coordination "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Coordination"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CoordinationService_Session_FullMethodName      = "/Ydb.Coordination.V1.CoordinationService/Session"
	CoordinationService_CreateNode_FullMethodName   = "/Ydb.Coordination.V1.CoordinationService/CreateNode"
	CoordinationService_AlterNode_FullMethodName    = "/Ydb.Coordination.V1.CoordinationService/AlterNode"
	CoordinationService_DropNode_FullMethodName     = "/Ydb.Coordination.V1.CoordinationService/DropNode"
	CoordinationService_DescribeNode_FullMethodName = "/Ydb.Coordination.V1.CoordinationService/DescribeNode"
)

// CoordinationServiceClient is the client API for CoordinationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoordinationServiceClient interface {
	// *
	// Bidirectional stream used to establish a session with a coordination node
	//
	// Relevant APIs for managing semaphores, distributed locking, creating or
	// restoring a previously established session are described using nested
	// messages in SessionRequest and SessionResponse. Session is established
	// with a specific coordination node (previously created using CreateNode
	// below) and semaphores are local to that coordination node.
	Session(ctx context.Context, opts ...grpc.CallOption) (CoordinationService_SessionClient, error)
	// Creates a new coordination node
	CreateNode(ctx context.Context, in *Ydb_Coordination.CreateNodeRequest, opts ...grpc.CallOption) (*Ydb_Coordination.CreateNodeResponse, error)
	// Modifies settings of a coordination node
	AlterNode(ctx context.Context, in *Ydb_Coordination.AlterNodeRequest, opts ...grpc.CallOption) (*Ydb_Coordination.AlterNodeResponse, error)
	// Drops a coordination node
	DropNode(ctx context.Context, in *Ydb_Coordination.DropNodeRequest, opts ...grpc.CallOption) (*Ydb_Coordination.DropNodeResponse, error)
	// Describes a coordination node
	DescribeNode(ctx context.Context, in *Ydb_Coordination.DescribeNodeRequest, opts ...grpc.CallOption) (*Ydb_Coordination.DescribeNodeResponse, error)
}

type coordinationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoordinationServiceClient(cc grpc.ClientConnInterface) CoordinationServiceClient {
	return &coordinationServiceClient{cc}
}

func (c *coordinationServiceClient) Session(ctx context.Context, opts ...grpc.CallOption) (CoordinationService_SessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &CoordinationService_ServiceDesc.Streams[0], CoordinationService_Session_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &coordinationServiceSessionClient{stream}
	return x, nil
}

type CoordinationService_SessionClient interface {
	Send(*Ydb_Coordination.SessionRequest) error
	Recv() (*Ydb_Coordination.SessionResponse, error)
	grpc.ClientStream
}

type coordinationServiceSessionClient struct {
	grpc.ClientStream
}

func (x *coordinationServiceSessionClient) Send(m *Ydb_Coordination.SessionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *coordinationServiceSessionClient) Recv() (*Ydb_Coordination.SessionResponse, error) {
	m := new(Ydb_Coordination.SessionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coordinationServiceClient) CreateNode(ctx context.Context, in *Ydb_Coordination.CreateNodeRequest, opts ...grpc.CallOption) (*Ydb_Coordination.CreateNodeResponse, error) {
	out := new(Ydb_Coordination.CreateNodeResponse)
	err := c.cc.Invoke(ctx, CoordinationService_CreateNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinationServiceClient) AlterNode(ctx context.Context, in *Ydb_Coordination.AlterNodeRequest, opts ...grpc.CallOption) (*Ydb_Coordination.AlterNodeResponse, error) {
	out := new(Ydb_Coordination.AlterNodeResponse)
	err := c.cc.Invoke(ctx, CoordinationService_AlterNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinationServiceClient) DropNode(ctx context.Context, in *Ydb_Coordination.DropNodeRequest, opts ...grpc.CallOption) (*Ydb_Coordination.DropNodeResponse, error) {
	out := new(Ydb_Coordination.DropNodeResponse)
	err := c.cc.Invoke(ctx, CoordinationService_DropNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinationServiceClient) DescribeNode(ctx context.Context, in *Ydb_Coordination.DescribeNodeRequest, opts ...grpc.CallOption) (*Ydb_Coordination.DescribeNodeResponse, error) {
	out := new(Ydb_Coordination.DescribeNodeResponse)
	err := c.cc.Invoke(ctx, CoordinationService_DescribeNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinationServiceServer is the server API for CoordinationService service.
// All implementations must embed UnimplementedCoordinationServiceServer
// for forward compatibility
type CoordinationServiceServer interface {
	// *
	// Bidirectional stream used to establish a session with a coordination node
	//
	// Relevant APIs for managing semaphores, distributed locking, creating or
	// restoring a previously established session are described using nested
	// messages in SessionRequest and SessionResponse. Session is established
	// with a specific coordination node (previously created using CreateNode
	// below) and semaphores are local to that coordination node.
	Session(CoordinationService_SessionServer) error
	// Creates a new coordination node
	CreateNode(context.Context, *Ydb_Coordination.CreateNodeRequest) (*Ydb_Coordination.CreateNodeResponse, error)
	// Modifies settings of a coordination node
	AlterNode(context.Context, *Ydb_Coordination.AlterNodeRequest) (*Ydb_Coordination.AlterNodeResponse, error)
	// Drops a coordination node
	DropNode(context.Context, *Ydb_Coordination.DropNodeRequest) (*Ydb_Coordination.DropNodeResponse, error)
	// Describes a coordination node
	DescribeNode(context.Context, *Ydb_Coordination.DescribeNodeRequest) (*Ydb_Coordination.DescribeNodeResponse, error)
	mustEmbedUnimplementedCoordinationServiceServer()
}

// UnimplementedCoordinationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCoordinationServiceServer struct {
}

func (UnimplementedCoordinationServiceServer) Session(CoordinationService_SessionServer) error {
	return status.Errorf(codes.Unimplemented, "method Session not implemented")
}
func (UnimplementedCoordinationServiceServer) CreateNode(context.Context, *Ydb_Coordination.CreateNodeRequest) (*Ydb_Coordination.CreateNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNode not implemented")
}
func (UnimplementedCoordinationServiceServer) AlterNode(context.Context, *Ydb_Coordination.AlterNodeRequest) (*Ydb_Coordination.AlterNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterNode not implemented")
}
func (UnimplementedCoordinationServiceServer) DropNode(context.Context, *Ydb_Coordination.DropNodeRequest) (*Ydb_Coordination.DropNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropNode not implemented")
}
func (UnimplementedCoordinationServiceServer) DescribeNode(context.Context, *Ydb_Coordination.DescribeNodeRequest) (*Ydb_Coordination.DescribeNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeNode not implemented")
}
func (UnimplementedCoordinationServiceServer) mustEmbedUnimplementedCoordinationServiceServer() {}

// UnsafeCoordinationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoordinationServiceServer will
// result in compilation errors.
type UnsafeCoordinationServiceServer interface {
	mustEmbedUnimplementedCoordinationServiceServer()
}

func RegisterCoordinationServiceServer(s grpc.ServiceRegistrar, srv CoordinationServiceServer) {
	s.RegisterService(&CoordinationService_ServiceDesc, srv)
}

func _CoordinationService_Session_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CoordinationServiceServer).Session(&coordinationServiceSessionServer{stream})
}

type CoordinationService_SessionServer interface {
	Send(*Ydb_Coordination.SessionResponse) error
	Recv() (*Ydb_Coordination.SessionRequest, error)
	grpc.ServerStream
}

type coordinationServiceSessionServer struct {
	grpc.ServerStream
}

func (x *coordinationServiceSessionServer) Send(m *Ydb_Coordination.SessionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *coordinationServiceSessionServer) Recv() (*Ydb_Coordination.SessionRequest, error) {
	m := new(Ydb_Coordination.SessionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CoordinationService_CreateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Coordination.CreateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinationServiceServer).CreateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoordinationService_CreateNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinationServiceServer).CreateNode(ctx, req.(*Ydb_Coordination.CreateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoordinationService_AlterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Coordination.AlterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinationServiceServer).AlterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoordinationService_AlterNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinationServiceServer).AlterNode(ctx, req.(*Ydb_Coordination.AlterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoordinationService_DropNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Coordination.DropNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinationServiceServer).DropNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoordinationService_DropNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinationServiceServer).DropNode(ctx, req.(*Ydb_Coordination.DropNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoordinationService_DescribeNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Coordination.DescribeNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinationServiceServer).DescribeNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoordinationService_DescribeNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinationServiceServer).DescribeNode(ctx, req.(*Ydb_Coordination.DescribeNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoordinationService_ServiceDesc is the grpc.ServiceDesc for CoordinationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoordinationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.Coordination.V1.CoordinationService",
	HandlerType: (*CoordinationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNode",
			Handler:    _CoordinationService_CreateNode_Handler,
		},
		{
			MethodName: "AlterNode",
			Handler:    _CoordinationService_AlterNode_Handler,
		},
		{
			MethodName: "DropNode",
			Handler:    _CoordinationService_DropNode_Handler,
		},
		{
			MethodName: "DescribeNode",
			Handler:    _CoordinationService_DescribeNode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Session",
			Handler:       _CoordinationService_Session_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ydb_coordination_v1.proto",
}
