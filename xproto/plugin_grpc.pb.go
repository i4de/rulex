// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package xproto

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

// XPluginClient is the client API for XPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XPluginClient interface {
	//
	Init(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	//
	Start(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	//
	Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	//
	PluginMetaInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*XPluginMetaInfo, error)
}

type xPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewXPluginClient(cc grpc.ClientConnInterface) XPluginClient {
	return &xPluginClient{cc}
}

func (c *xPluginClient) Init(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/xplugin.XPlugin/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xPluginClient) Start(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/xplugin.XPlugin/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xPluginClient) Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/xplugin.XPlugin/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xPluginClient) PluginMetaInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*XPluginMetaInfo, error) {
	out := new(XPluginMetaInfo)
	err := c.cc.Invoke(ctx, "/xplugin.XPlugin/PluginMetaInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XPluginServer is the server API for XPlugin service.
// All implementations must embed UnimplementedXPluginServer
// for forward compatibility
type XPluginServer interface {
	//
	Init(context.Context, *Request) (*Response, error)
	//
	Start(context.Context, *Request) (*Response, error)
	//
	Stop(context.Context, *Request) (*Response, error)
	//
	PluginMetaInfo(context.Context, *Request) (*XPluginMetaInfo, error)
	mustEmbedUnimplementedXPluginServer()
}

// UnimplementedXPluginServer must be embedded to have forward compatible implementations.
type UnimplementedXPluginServer struct {
}

func (UnimplementedXPluginServer) Init(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedXPluginServer) Start(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedXPluginServer) Stop(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedXPluginServer) PluginMetaInfo(context.Context, *Request) (*XPluginMetaInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginMetaInfo not implemented")
}
func (UnimplementedXPluginServer) mustEmbedUnimplementedXPluginServer() {}

// UnsafeXPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XPluginServer will
// result in compilation errors.
type UnsafeXPluginServer interface {
	mustEmbedUnimplementedXPluginServer()
}

func RegisterXPluginServer(s grpc.ServiceRegistrar, srv XPluginServer) {
	s.RegisterService(&XPlugin_ServiceDesc, srv)
}

func _XPlugin_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XPluginServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xplugin.XPlugin/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XPluginServer).Init(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _XPlugin_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XPluginServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xplugin.XPlugin/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XPluginServer).Start(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _XPlugin_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XPluginServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xplugin.XPlugin/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XPluginServer).Stop(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _XPlugin_PluginMetaInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XPluginServer).PluginMetaInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xplugin.XPlugin/PluginMetaInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XPluginServer).PluginMetaInfo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// XPlugin_ServiceDesc is the grpc.ServiceDesc for XPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var XPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xplugin.XPlugin",
	HandlerType: (*XPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _XPlugin_Init_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _XPlugin_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _XPlugin_Stop_Handler,
		},
		{
			MethodName: "PluginMetaInfo",
			Handler:    _XPlugin_PluginMetaInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin.proto",
}