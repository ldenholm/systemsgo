// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// AdderServiceClient is the client API for AdderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdderServiceClient interface {
	Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type adderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdderServiceClient(cc grpc.ClientConnInterface) AdderServiceClient {
	return &adderServiceClient{cc}
}

func (c *adderServiceClient) Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.AdderService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adderServiceClient) Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.AdderService/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdderServiceServer is the server API for AdderService service.
// All implementations must embed UnimplementedAdderServiceServer
// for forward compatibility
type AdderServiceServer interface {
	Add(context.Context, *Request) (*Response, error)
	Multiply(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedAdderServiceServer()
}

// UnimplementedAdderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdderServiceServer struct {
}

func (UnimplementedAdderServiceServer) Add(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedAdderServiceServer) Multiply(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedAdderServiceServer) mustEmbedUnimplementedAdderServiceServer() {}

// UnsafeAdderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdderServiceServer will
// result in compilation errors.
type UnsafeAdderServiceServer interface {
	mustEmbedUnimplementedAdderServiceServer()
}

func RegisterAdderServiceServer(s grpc.ServiceRegistrar, srv AdderServiceServer) {
	s.RegisterService(&AdderService_ServiceDesc, srv)
}

func _AdderService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AdderService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServiceServer).Add(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdderService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AdderService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServiceServer).Multiply(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// AdderService_ServiceDesc is the grpc.ServiceDesc for AdderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AdderService",
	HandlerType: (*AdderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AdderService_Add_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _AdderService_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/service.proto",
}
