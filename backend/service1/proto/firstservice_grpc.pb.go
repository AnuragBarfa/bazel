// Run from cm "protoc --proto_path=proto proto/*.proto --go-grpc_out=." to generate go client manually.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: firstservice.proto

package __

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

const (
	FirstService_SayHello_FullMethodName      = "/firstservice.FirstService/SayHello"
	FirstService_SayHelloAgain_FullMethodName = "/firstservice.FirstService/SayHelloAgain"
)

// FirstServiceClient is the client API for FirstService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FirstServiceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type firstServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFirstServiceClient(cc grpc.ClientConnInterface) FirstServiceClient {
	return &firstServiceClient{cc}
}

func (c *firstServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, FirstService_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firstServiceClient) SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, FirstService_SayHelloAgain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FirstServiceServer is the server API for FirstService service.
// All implementations must embed UnimplementedFirstServiceServer
// for forward compatibility
type FirstServiceServer interface {
	SayHello(context.Context, *Message) (*Message, error)
	SayHelloAgain(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedFirstServiceServer()
}

// UnimplementedFirstServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFirstServiceServer struct {
}

func (UnimplementedFirstServiceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedFirstServiceServer) SayHelloAgain(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloAgain not implemented")
}
func (UnimplementedFirstServiceServer) mustEmbedUnimplementedFirstServiceServer() {}

// UnsafeFirstServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FirstServiceServer will
// result in compilation errors.
type UnsafeFirstServiceServer interface {
	mustEmbedUnimplementedFirstServiceServer()
}

func RegisterFirstServiceServer(s grpc.ServiceRegistrar, srv FirstServiceServer) {
	s.RegisterService(&FirstService_ServiceDesc, srv)
}

func _FirstService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirstServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FirstService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirstServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirstService_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirstServiceServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FirstService_SayHelloAgain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirstServiceServer).SayHelloAgain(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FirstService_ServiceDesc is the grpc.ServiceDesc for FirstService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FirstService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "firstservice.FirstService",
	HandlerType: (*FirstServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _FirstService_SayHello_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _FirstService_SayHelloAgain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "firstservice.proto",
}
