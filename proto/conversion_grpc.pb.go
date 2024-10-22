// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: conversion.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ConversionService_Convert_FullMethodName = "/conversion.ConversionService/Convert"
)

// ConversionServiceClient is the client API for ConversionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConversionServiceClient interface {
	Convert(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error)
}

type conversionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversionServiceClient(cc grpc.ClientConnInterface) ConversionServiceClient {
	return &conversionServiceClient{cc}
}

func (c *conversionServiceClient) Convert(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConvertResponse)
	err := c.cc.Invoke(ctx, ConversionService_Convert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionServiceServer is the server API for ConversionService service.
// All implementations must embed UnimplementedConversionServiceServer
// for forward compatibility.
type ConversionServiceServer interface {
	Convert(context.Context, *ConvertRequest) (*ConvertResponse, error)
	mustEmbedUnimplementedConversionServiceServer()
}

// UnimplementedConversionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConversionServiceServer struct{}

func (UnimplementedConversionServiceServer) Convert(context.Context, *ConvertRequest) (*ConvertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Convert not implemented")
}
func (UnimplementedConversionServiceServer) mustEmbedUnimplementedConversionServiceServer() {}
func (UnimplementedConversionServiceServer) testEmbeddedByValue()                           {}

// UnsafeConversionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversionServiceServer will
// result in compilation errors.
type UnsafeConversionServiceServer interface {
	mustEmbedUnimplementedConversionServiceServer()
}

func RegisterConversionServiceServer(s grpc.ServiceRegistrar, srv ConversionServiceServer) {
	// If the following call pancis, it indicates UnimplementedConversionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ConversionService_ServiceDesc, srv)
}

func _ConversionService_Convert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionServiceServer).Convert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConversionService_Convert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionServiceServer).Convert(ctx, req.(*ConvertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConversionService_ServiceDesc is the grpc.ServiceDesc for ConversionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConversionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "conversion.ConversionService",
	HandlerType: (*ConversionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Convert",
			Handler:    _ConversionService_Convert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conversion.proto",
}
