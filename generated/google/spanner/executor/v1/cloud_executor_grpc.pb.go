// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: google/spanner/executor/v1/cloud_executor.proto

package executorpb

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

// SpannerExecutorProxyClient is the client API for SpannerExecutorProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpannerExecutorProxyClient interface {
	// ExecuteActionAsync is a streaming call that starts executing a new Spanner
	// action.
	//
	// For each request, the server will reply with one or more responses, but
	// only the last response will contain status in the outcome.
	//
	// Responses can be matched to requests by action_id. It is allowed to have
	// multiple actions in flight--in that case, actions are be executed in
	// parallel.
	ExecuteActionAsync(ctx context.Context, opts ...grpc.CallOption) (SpannerExecutorProxy_ExecuteActionAsyncClient, error)
}

type spannerExecutorProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewSpannerExecutorProxyClient(cc grpc.ClientConnInterface) SpannerExecutorProxyClient {
	return &spannerExecutorProxyClient{cc}
}

func (c *spannerExecutorProxyClient) ExecuteActionAsync(ctx context.Context, opts ...grpc.CallOption) (SpannerExecutorProxy_ExecuteActionAsyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &SpannerExecutorProxy_ServiceDesc.Streams[0], "/google.spanner.executor.v1.SpannerExecutorProxy/ExecuteActionAsync", opts...)
	if err != nil {
		return nil, err
	}
	x := &spannerExecutorProxyExecuteActionAsyncClient{stream}
	return x, nil
}

type SpannerExecutorProxy_ExecuteActionAsyncClient interface {
	Send(*SpannerAsyncActionRequest) error
	Recv() (*SpannerAsyncActionResponse, error)
	grpc.ClientStream
}

type spannerExecutorProxyExecuteActionAsyncClient struct {
	grpc.ClientStream
}

func (x *spannerExecutorProxyExecuteActionAsyncClient) Send(m *SpannerAsyncActionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *spannerExecutorProxyExecuteActionAsyncClient) Recv() (*SpannerAsyncActionResponse, error) {
	m := new(SpannerAsyncActionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpannerExecutorProxyServer is the server API for SpannerExecutorProxy service.
// All implementations must embed UnimplementedSpannerExecutorProxyServer
// for forward compatibility
type SpannerExecutorProxyServer interface {
	// ExecuteActionAsync is a streaming call that starts executing a new Spanner
	// action.
	//
	// For each request, the server will reply with one or more responses, but
	// only the last response will contain status in the outcome.
	//
	// Responses can be matched to requests by action_id. It is allowed to have
	// multiple actions in flight--in that case, actions are be executed in
	// parallel.
	ExecuteActionAsync(SpannerExecutorProxy_ExecuteActionAsyncServer) error
	mustEmbedUnimplementedSpannerExecutorProxyServer()
}

// UnimplementedSpannerExecutorProxyServer must be embedded to have forward compatible implementations.
type UnimplementedSpannerExecutorProxyServer struct {
}

func (UnimplementedSpannerExecutorProxyServer) ExecuteActionAsync(SpannerExecutorProxy_ExecuteActionAsyncServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecuteActionAsync not implemented")
}
func (UnimplementedSpannerExecutorProxyServer) mustEmbedUnimplementedSpannerExecutorProxyServer() {}

// UnsafeSpannerExecutorProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpannerExecutorProxyServer will
// result in compilation errors.
type UnsafeSpannerExecutorProxyServer interface {
	mustEmbedUnimplementedSpannerExecutorProxyServer()
}

func RegisterSpannerExecutorProxyServer(s grpc.ServiceRegistrar, srv SpannerExecutorProxyServer) {
	s.RegisterService(&SpannerExecutorProxy_ServiceDesc, srv)
}

func _SpannerExecutorProxy_ExecuteActionAsync_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SpannerExecutorProxyServer).ExecuteActionAsync(&spannerExecutorProxyExecuteActionAsyncServer{stream})
}

type SpannerExecutorProxy_ExecuteActionAsyncServer interface {
	Send(*SpannerAsyncActionResponse) error
	Recv() (*SpannerAsyncActionRequest, error)
	grpc.ServerStream
}

type spannerExecutorProxyExecuteActionAsyncServer struct {
	grpc.ServerStream
}

func (x *spannerExecutorProxyExecuteActionAsyncServer) Send(m *SpannerAsyncActionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *spannerExecutorProxyExecuteActionAsyncServer) Recv() (*SpannerAsyncActionRequest, error) {
	m := new(SpannerAsyncActionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpannerExecutorProxy_ServiceDesc is the grpc.ServiceDesc for SpannerExecutorProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpannerExecutorProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.spanner.executor.v1.SpannerExecutorProxy",
	HandlerType: (*SpannerExecutorProxyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteActionAsync",
			Handler:       _SpannerExecutorProxy_ExecuteActionAsync_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "google/spanner/executor/v1/cloud_executor.proto",
}
