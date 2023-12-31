// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: filemeta.proto

package filemeta

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
	Filemeta_GetFileMeta_FullMethodName        = "/upload.filemeta/GetFileMeta"
	Filemeta_GetUserFileMeta_FullMethodName    = "/upload.filemeta/GetUserFileMeta"
	Filemeta_UpdataUserFileMeta_FullMethodName = "/upload.filemeta/UpdataUserFileMeta"
)

// FilemetaClient is the client API for Filemeta service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilemetaClient interface {
	GetFileMeta(ctx context.Context, in *GetFileMetaReq, opts ...grpc.CallOption) (*FileMeta, error)
	GetUserFileMeta(ctx context.Context, in *GetUserFileMetaReq, opts ...grpc.CallOption) (*GetUserFileMetaResp, error)
	UpdataUserFileMeta(ctx context.Context, in *UpdataUserFileMetaReq, opts ...grpc.CallOption) (*CommonResp, error)
}

type filemetaClient struct {
	cc grpc.ClientConnInterface
}

func NewFilemetaClient(cc grpc.ClientConnInterface) FilemetaClient {
	return &filemetaClient{cc}
}

func (c *filemetaClient) GetFileMeta(ctx context.Context, in *GetFileMetaReq, opts ...grpc.CallOption) (*FileMeta, error) {
	out := new(FileMeta)
	err := c.cc.Invoke(ctx, Filemeta_GetFileMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filemetaClient) GetUserFileMeta(ctx context.Context, in *GetUserFileMetaReq, opts ...grpc.CallOption) (*GetUserFileMetaResp, error) {
	out := new(GetUserFileMetaResp)
	err := c.cc.Invoke(ctx, Filemeta_GetUserFileMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filemetaClient) UpdataUserFileMeta(ctx context.Context, in *UpdataUserFileMetaReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Filemeta_UpdataUserFileMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilemetaServer is the server API for Filemeta service.
// All implementations must embed UnimplementedFilemetaServer
// for forward compatibility
type FilemetaServer interface {
	GetFileMeta(context.Context, *GetFileMetaReq) (*FileMeta, error)
	GetUserFileMeta(context.Context, *GetUserFileMetaReq) (*GetUserFileMetaResp, error)
	UpdataUserFileMeta(context.Context, *UpdataUserFileMetaReq) (*CommonResp, error)
	mustEmbedUnimplementedFilemetaServer()
}

// UnimplementedFilemetaServer must be embedded to have forward compatible implementations.
type UnimplementedFilemetaServer struct {
}

func (UnimplementedFilemetaServer) GetFileMeta(context.Context, *GetFileMetaReq) (*FileMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileMeta not implemented")
}
func (UnimplementedFilemetaServer) GetUserFileMeta(context.Context, *GetUserFileMetaReq) (*GetUserFileMetaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFileMeta not implemented")
}
func (UnimplementedFilemetaServer) UpdataUserFileMeta(context.Context, *UpdataUserFileMetaReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdataUserFileMeta not implemented")
}
func (UnimplementedFilemetaServer) mustEmbedUnimplementedFilemetaServer() {}

// UnsafeFilemetaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilemetaServer will
// result in compilation errors.
type UnsafeFilemetaServer interface {
	mustEmbedUnimplementedFilemetaServer()
}

func RegisterFilemetaServer(s grpc.ServiceRegistrar, srv FilemetaServer) {
	s.RegisterService(&Filemeta_ServiceDesc, srv)
}

func _Filemeta_GetFileMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileMetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilemetaServer).GetFileMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Filemeta_GetFileMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilemetaServer).GetFileMeta(ctx, req.(*GetFileMetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Filemeta_GetUserFileMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFileMetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilemetaServer).GetUserFileMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Filemeta_GetUserFileMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilemetaServer).GetUserFileMeta(ctx, req.(*GetUserFileMetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Filemeta_UpdataUserFileMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdataUserFileMetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilemetaServer).UpdataUserFileMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Filemeta_UpdataUserFileMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilemetaServer).UpdataUserFileMeta(ctx, req.(*UpdataUserFileMetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Filemeta_ServiceDesc is the grpc.ServiceDesc for Filemeta service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Filemeta_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "upload.filemeta",
	HandlerType: (*FilemetaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFileMeta",
			Handler:    _Filemeta_GetFileMeta_Handler,
		},
		{
			MethodName: "GetUserFileMeta",
			Handler:    _Filemeta_GetUserFileMeta_Handler,
		},
		{
			MethodName: "UpdataUserFileMeta",
			Handler:    _Filemeta_UpdataUserFileMeta_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "filemeta.proto",
}
