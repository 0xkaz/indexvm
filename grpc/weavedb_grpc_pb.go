package grpc

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

// DBClient is the client API for DB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DBClient interface {
	Query(ctx context.Context, in *WeaveDBRequest, opts ...grpc.CallOption) (*WeaveDBReply, error)
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
}

type dBClient struct {
	cc grpc.ClientConnInterface
}

func NewDBClient(cc grpc.ClientConnInterface) DBClient {
	return &dBClient{cc}
}

func (c *dBClient) Query(ctx context.Context, in *WeaveDBRequest, opts ...grpc.CallOption) (*WeaveDBReply, error) {
	out := new(WeaveDBReply)
	err := c.cc.Invoke(ctx, "/weavedb.DB/query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/weavedb.DB/sayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/weavedb.DB/ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBServer is the server API for DB service.
// All implementations must embed UnimplementedDBServer
// for forward compatibility
type DBServer interface {
	Query(context.Context, *WeaveDBRequest) (*WeaveDBReply, error)
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
	mustEmbedUnimplementedDBServer()
}

// UnimplementedDBServer must be embedded to have forward compatible implementations.
type UnimplementedDBServer struct {
}

func (UnimplementedDBServer) Query(context.Context, *WeaveDBRequest) (*WeaveDBReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedDBServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedDBServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDBServer) mustEmbedUnimplementedDBServer() {}

// UnsafeDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DBServer will
// result in compilation errors.
type UnsafeDBServer interface {
	mustEmbedUnimplementedDBServer()
}

func RegisterDBServer(s grpc.ServiceRegistrar, srv DBServer) {
	s.RegisterService(&DB_ServiceDesc, srv)
}

func _DB_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeaveDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weavedb.DB/query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Query(ctx, req.(*WeaveDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weavedb.DB/sayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weavedb.DB/ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DB_ServiceDesc is the grpc.ServiceDesc for DB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "weavedb.DB",
	HandlerType: (*DBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "query",
			Handler:    _DB_Query_Handler,
		},
		{
			MethodName: "sayHello",
			Handler:    _DB_SayHello_Handler,
		},
		{
			MethodName: "ping",
			Handler:    _DB_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weavedb.proto",
}
