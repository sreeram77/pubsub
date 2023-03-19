// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: subscriber/subscriber.proto

package subscriber

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

// SubscriberClient is the client API for Subscriber service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriberClient interface {
	Subscribe(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Ack, error)
}

type subscriberClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriberClient(cc grpc.ClientConnInterface) SubscriberClient {
	return &subscriberClient{cc}
}

func (c *subscriberClient) Subscribe(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/Subscriber/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriberServer is the server API for Subscriber service.
// All implementations must embed UnimplementedSubscriberServer
// for forward compatibility
type SubscriberServer interface {
	Subscribe(context.Context, *Topic) (*Ack, error)
	mustEmbedUnimplementedSubscriberServer()
}

// UnimplementedSubscriberServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriberServer struct {
}

func (UnimplementedSubscriberServer) Subscribe(context.Context, *Topic) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSubscriberServer) mustEmbedUnimplementedSubscriberServer() {}

// UnsafeSubscriberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriberServer will
// result in compilation errors.
type UnsafeSubscriberServer interface {
	mustEmbedUnimplementedSubscriberServer()
}

func RegisterSubscriberServer(s grpc.ServiceRegistrar, srv SubscriberServer) {
	s.RegisterService(&Subscriber_ServiceDesc, srv)
}

func _Subscriber_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriberServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Subscriber/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriberServer).Subscribe(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

// Subscriber_ServiceDesc is the grpc.ServiceDesc for Subscriber service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Subscriber_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Subscriber",
	HandlerType: (*SubscriberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _Subscriber_Subscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscriber/subscriber.proto",
}
