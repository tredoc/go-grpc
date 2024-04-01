// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: src/ping.proto

package pb

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

// ResponserClient is the client API for Responser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResponserClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error)
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (Responser_GetListClient, error)
	SendList(ctx context.Context, opts ...grpc.CallOption) (Responser_SendListClient, error)
	HandleJob(ctx context.Context, opts ...grpc.CallOption) (Responser_HandleJobClient, error)
}

type responserClient struct {
	cc grpc.ClientConnInterface
}

func NewResponserClient(cc grpc.ClientConnInterface) ResponserClient {
	return &responserClient{cc}
}

func (c *responserClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/pb.Responser/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *responserClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (Responser_GetListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Responser_ServiceDesc.Streams[0], "/pb.Responser/GetList", opts...)
	if err != nil {
		return nil, err
	}
	x := &responserGetListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Responser_GetListClient interface {
	Recv() (*GetListResponse, error)
	grpc.ClientStream
}

type responserGetListClient struct {
	grpc.ClientStream
}

func (x *responserGetListClient) Recv() (*GetListResponse, error) {
	m := new(GetListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *responserClient) SendList(ctx context.Context, opts ...grpc.CallOption) (Responser_SendListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Responser_ServiceDesc.Streams[1], "/pb.Responser/SendList", opts...)
	if err != nil {
		return nil, err
	}
	x := &responserSendListClient{stream}
	return x, nil
}

type Responser_SendListClient interface {
	Send(*SendListRequest) error
	CloseAndRecv() (*SendListResponse, error)
	grpc.ClientStream
}

type responserSendListClient struct {
	grpc.ClientStream
}

func (x *responserSendListClient) Send(m *SendListRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *responserSendListClient) CloseAndRecv() (*SendListResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *responserClient) HandleJob(ctx context.Context, opts ...grpc.CallOption) (Responser_HandleJobClient, error) {
	stream, err := c.cc.NewStream(ctx, &Responser_ServiceDesc.Streams[2], "/pb.Responser/HandleJob", opts...)
	if err != nil {
		return nil, err
	}
	x := &responserHandleJobClient{stream}
	return x, nil
}

type Responser_HandleJobClient interface {
	Send(*HandleJobRequest) error
	Recv() (*HandleJobResponse, error)
	grpc.ClientStream
}

type responserHandleJobClient struct {
	grpc.ClientStream
}

func (x *responserHandleJobClient) Send(m *HandleJobRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *responserHandleJobClient) Recv() (*HandleJobResponse, error) {
	m := new(HandleJobResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ResponserServer is the server API for Responser service.
// All implementations must embed UnimplementedResponserServer
// for forward compatibility
type ResponserServer interface {
	Ping(context.Context, *PingRequest) (*PongResponse, error)
	GetList(*GetListRequest, Responser_GetListServer) error
	SendList(Responser_SendListServer) error
	HandleJob(Responser_HandleJobServer) error
	mustEmbedUnimplementedResponserServer()
}

// UnimplementedResponserServer must be embedded to have forward compatible implementations.
type UnimplementedResponserServer struct {
}

func (UnimplementedResponserServer) Ping(context.Context, *PingRequest) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedResponserServer) GetList(*GetListRequest, Responser_GetListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedResponserServer) SendList(Responser_SendListServer) error {
	return status.Errorf(codes.Unimplemented, "method SendList not implemented")
}
func (UnimplementedResponserServer) HandleJob(Responser_HandleJobServer) error {
	return status.Errorf(codes.Unimplemented, "method HandleJob not implemented")
}
func (UnimplementedResponserServer) mustEmbedUnimplementedResponserServer() {}

// UnsafeResponserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResponserServer will
// result in compilation errors.
type UnsafeResponserServer interface {
	mustEmbedUnimplementedResponserServer()
}

func RegisterResponserServer(s grpc.ServiceRegistrar, srv ResponserServer) {
	s.RegisterService(&Responser_ServiceDesc, srv)
}

func _Responser_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResponserServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Responser/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResponserServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Responser_GetList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResponserServer).GetList(m, &responserGetListServer{stream})
}

type Responser_GetListServer interface {
	Send(*GetListResponse) error
	grpc.ServerStream
}

type responserGetListServer struct {
	grpc.ServerStream
}

func (x *responserGetListServer) Send(m *GetListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Responser_SendList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ResponserServer).SendList(&responserSendListServer{stream})
}

type Responser_SendListServer interface {
	SendAndClose(*SendListResponse) error
	Recv() (*SendListRequest, error)
	grpc.ServerStream
}

type responserSendListServer struct {
	grpc.ServerStream
}

func (x *responserSendListServer) SendAndClose(m *SendListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *responserSendListServer) Recv() (*SendListRequest, error) {
	m := new(SendListRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Responser_HandleJob_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ResponserServer).HandleJob(&responserHandleJobServer{stream})
}

type Responser_HandleJobServer interface {
	Send(*HandleJobResponse) error
	Recv() (*HandleJobRequest, error)
	grpc.ServerStream
}

type responserHandleJobServer struct {
	grpc.ServerStream
}

func (x *responserHandleJobServer) Send(m *HandleJobResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *responserHandleJobServer) Recv() (*HandleJobRequest, error) {
	m := new(HandleJobRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Responser_ServiceDesc is the grpc.ServiceDesc for Responser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Responser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Responser",
	HandlerType: (*ResponserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Responser_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetList",
			Handler:       _Responser_GetList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendList",
			Handler:       _Responser_SendList_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "HandleJob",
			Handler:       _Responser_HandleJob_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "src/ping.proto",
}
