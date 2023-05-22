// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: arithmetic/arithmetic.proto

package arithmetic

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
	Arithmetic_Add_FullMethodName      = "/arithmetic.Arithmetic/Add"
	Arithmetic_Divide_FullMethodName   = "/arithmetic.Arithmetic/Divide"
	Arithmetic_Multiply_FullMethodName = "/arithmetic.Arithmetic/Multiply"
	Arithmetic_Subtract_FullMethodName = "/arithmetic.Arithmetic/Subtract"
)

// ArithmeticClient is the client API for Arithmetic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArithmeticClient interface {
	Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Divide(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Subtract(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type arithmeticClient struct {
	cc grpc.ClientConnInterface
}

func NewArithmeticClient(cc grpc.ClientConnInterface) ArithmeticClient {
	return &arithmeticClient{cc}
}

func (c *arithmeticClient) Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Arithmetic_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticClient) Divide(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Arithmetic_Divide_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticClient) Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Arithmetic_Multiply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticClient) Subtract(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Arithmetic_Subtract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithmeticServer is the server API for Arithmetic service.
// All implementations must embed UnimplementedArithmeticServer
// for forward compatibility
type ArithmeticServer interface {
	Add(context.Context, *Request) (*Response, error)
	Divide(context.Context, *Request) (*Response, error)
	Multiply(context.Context, *Request) (*Response, error)
	Subtract(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedArithmeticServer()
}

// UnimplementedArithmeticServer must be embedded to have forward compatible implementations.
type UnimplementedArithmeticServer struct {
}

func (UnimplementedArithmeticServer) Add(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedArithmeticServer) Divide(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Divide not implemented")
}
func (UnimplementedArithmeticServer) Multiply(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedArithmeticServer) Subtract(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (UnimplementedArithmeticServer) mustEmbedUnimplementedArithmeticServer() {}

// UnsafeArithmeticServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArithmeticServer will
// result in compilation errors.
type UnsafeArithmeticServer interface {
	mustEmbedUnimplementedArithmeticServer()
}

func RegisterArithmeticServer(s grpc.ServiceRegistrar, srv ArithmeticServer) {
	s.RegisterService(&Arithmetic_ServiceDesc, srv)
}

func _Arithmetic_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Arithmetic_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServer).Add(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arithmetic_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Arithmetic_Divide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServer).Divide(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arithmetic_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Arithmetic_Multiply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServer).Multiply(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arithmetic_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Arithmetic_Subtract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServer).Subtract(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Arithmetic_ServiceDesc is the grpc.ServiceDesc for Arithmetic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Arithmetic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "arithmetic.Arithmetic",
	HandlerType: (*ArithmeticServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Arithmetic_Add_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _Arithmetic_Divide_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _Arithmetic_Multiply_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _Arithmetic_Subtract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arithmetic/arithmetic.proto",
}
