// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.19.1
// source: grpc_script.proto

package grpc_script

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

// ScriptServiceClient is the client API for ScriptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScriptServiceClient interface {
	OnScriptCall(ctx context.Context, in *ScriptRequest, opts ...grpc.CallOption) (*ScriptResponse, error)
	OnImportSet(ctx context.Context, in *ImportSetRequest, opts ...grpc.CallOption) (*ImportSetResponse, error)
}

type scriptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScriptServiceClient(cc grpc.ClientConnInterface) ScriptServiceClient {
	return &scriptServiceClient{cc}
}

func (c *scriptServiceClient) OnScriptCall(ctx context.Context, in *ScriptRequest, opts ...grpc.CallOption) (*ScriptResponse, error) {
	out := new(ScriptResponse)
	err := c.cc.Invoke(ctx, "/grpc_script.ScriptService/OnScriptCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scriptServiceClient) OnImportSet(ctx context.Context, in *ImportSetRequest, opts ...grpc.CallOption) (*ImportSetResponse, error) {
	out := new(ImportSetResponse)
	err := c.cc.Invoke(ctx, "/grpc_script.ScriptService/OnImportSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScriptServiceServer is the server API for ScriptService service.
// All implementations must embed UnimplementedScriptServiceServer
// for forward compatibility
type ScriptServiceServer interface {
	OnScriptCall(context.Context, *ScriptRequest) (*ScriptResponse, error)
	OnImportSet(context.Context, *ImportSetRequest) (*ImportSetResponse, error)
	mustEmbedUnimplementedScriptServiceServer()
}

// UnimplementedScriptServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScriptServiceServer struct {
}

func (UnimplementedScriptServiceServer) OnScriptCall(context.Context, *ScriptRequest) (*ScriptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnScriptCall not implemented")
}
func (UnimplementedScriptServiceServer) OnImportSet(context.Context, *ImportSetRequest) (*ImportSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnImportSet not implemented")
}
func (UnimplementedScriptServiceServer) mustEmbedUnimplementedScriptServiceServer() {}

// UnsafeScriptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScriptServiceServer will
// result in compilation errors.
type UnsafeScriptServiceServer interface {
	mustEmbedUnimplementedScriptServiceServer()
}

func RegisterScriptServiceServer(s grpc.ServiceRegistrar, srv ScriptServiceServer) {
	s.RegisterService(&ScriptService_ServiceDesc, srv)
}

func _ScriptService_OnScriptCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScriptServiceServer).OnScriptCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_script.ScriptService/OnScriptCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScriptServiceServer).OnScriptCall(ctx, req.(*ScriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScriptService_OnImportSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScriptServiceServer).OnImportSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_script.ScriptService/OnImportSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScriptServiceServer).OnImportSet(ctx, req.(*ImportSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScriptService_ServiceDesc is the grpc.ServiceDesc for ScriptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScriptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_script.ScriptService",
	HandlerType: (*ScriptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnScriptCall",
			Handler:    _ScriptService_OnScriptCall_Handler,
		},
		{
			MethodName: "OnImportSet",
			Handler:    _ScriptService_OnImportSet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_script.proto",
}
