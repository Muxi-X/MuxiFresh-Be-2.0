// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: intro.proto

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

// GetClientClient is the client API for GetClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetClientClient interface {
	GetGroupIntro(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
}

type getClientClient struct {
	cc grpc.ClientConnInterface
}

func NewGetClientClient(cc grpc.ClientConnInterface) GetClientClient {
	return &getClientClient{cc}
}

func (c *getClientClient) GetGroupIntro(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, "/intro.GetClient/GetGroupIntro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetClientServer is the server API for GetClient service.
// All implementations must embed UnimplementedGetClientServer
// for forward compatibility
type GetClientServer interface {
	GetGroupIntro(context.Context, *GetReq) (*GetResp, error)
	mustEmbedUnimplementedGetClientServer()
}

// UnimplementedGetClientServer must be embedded to have forward compatible implementations.
type UnimplementedGetClientServer struct {
}

func (UnimplementedGetClientServer) GetGroupIntro(context.Context, *GetReq) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupIntro not implemented")
}
func (UnimplementedGetClientServer) mustEmbedUnimplementedGetClientServer() {}

// UnsafeGetClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetClientServer will
// result in compilation errors.
type UnsafeGetClientServer interface {
	mustEmbedUnimplementedGetClientServer()
}

func RegisterGetClientServer(s grpc.ServiceRegistrar, srv GetClientServer) {
	s.RegisterService(&GetClient_ServiceDesc, srv)
}

func _GetClient_GetGroupIntro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetClientServer).GetGroupIntro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/intro.GetClient/GetGroupIntro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetClientServer).GetGroupIntro(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GetClient_ServiceDesc is the grpc.ServiceDesc for GetClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "intro.GetClient",
	HandlerType: (*GetClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroupIntro",
			Handler:    _GetClient_GetGroupIntro_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "intro.proto",
}
