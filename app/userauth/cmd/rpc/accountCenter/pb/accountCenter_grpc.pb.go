// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: accountCenter.proto

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

const (
	AccountCenterClient_Register_FullMethodName     = "/accountCenter.accountCenterClient/Register"
	AccountCenterClient_Login_FullMethodName        = "/accountCenter.accountCenterClient/Login"
	AccountCenterClient_SetPassword_FullMethodName  = "/accountCenter.accountCenterClient/SetPassword"
	AccountCenterClient_CcnuLogin_FullMethodName    = "/accountCenter.accountCenterClient/ccnuLogin"
	AccountCenterClient_SetStudentID_FullMethodName = "/accountCenter.accountCenterClient/SetStudentID"
	AccountCenterClient_SetEmail_FullMethodName     = "/accountCenter.accountCenterClient/SetEmail"
)

// AccountCenterClientClient is the client API for AccountCenterClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountCenterClientClient interface {
	Register(ctx context.Context, in *RegisterDataReq, opts ...grpc.CallOption) (*RegisterDataResp, error)
	Login(ctx context.Context, in *LoginVerifyReq, opts ...grpc.CallOption) (*LoginVerifyResp, error)
	SetPassword(ctx context.Context, in *SetPasswordReq, opts ...grpc.CallOption) (*SetPasswordResp, error)
	CcnuLogin(ctx context.Context, in *CcnuLoginReq, opts ...grpc.CallOption) (*CcnuLoginResp, error)
	SetStudentID(ctx context.Context, in *SetStudentIDReq, opts ...grpc.CallOption) (*SetStudentIDResp, error)
	SetEmail(ctx context.Context, in *SetEmailReq, opts ...grpc.CallOption) (*SetEmailResp, error)
}

type accountCenterClientClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountCenterClientClient(cc grpc.ClientConnInterface) AccountCenterClientClient {
	return &accountCenterClientClient{cc}
}

func (c *accountCenterClientClient) Register(ctx context.Context, in *RegisterDataReq, opts ...grpc.CallOption) (*RegisterDataResp, error) {
	out := new(RegisterDataResp)
	err := c.cc.Invoke(ctx, AccountCenterClient_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClientClient) Login(ctx context.Context, in *LoginVerifyReq, opts ...grpc.CallOption) (*LoginVerifyResp, error) {
	out := new(LoginVerifyResp)
	err := c.cc.Invoke(ctx, AccountCenterClient_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClientClient) SetPassword(ctx context.Context, in *SetPasswordReq, opts ...grpc.CallOption) (*SetPasswordResp, error) {
	out := new(SetPasswordResp)
	err := c.cc.Invoke(ctx, AccountCenterClient_SetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClientClient) CcnuLogin(ctx context.Context, in *CcnuLoginReq, opts ...grpc.CallOption) (*CcnuLoginResp, error) {
	out := new(CcnuLoginResp)
	err := c.cc.Invoke(ctx, AccountCenterClient_CcnuLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClientClient) SetStudentID(ctx context.Context, in *SetStudentIDReq, opts ...grpc.CallOption) (*SetStudentIDResp, error) {
	out := new(SetStudentIDResp)
	err := c.cc.Invoke(ctx, AccountCenterClient_SetStudentID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClientClient) SetEmail(ctx context.Context, in *SetEmailReq, opts ...grpc.CallOption) (*SetEmailResp, error) {
	out := new(SetEmailResp)
	err := c.cc.Invoke(ctx, AccountCenterClient_SetEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountCenterClientServer is the server API for AccountCenterClient service.
// All implementations must embed UnimplementedAccountCenterClientServer
// for forward compatibility
type AccountCenterClientServer interface {
	Register(context.Context, *RegisterDataReq) (*RegisterDataResp, error)
	Login(context.Context, *LoginVerifyReq) (*LoginVerifyResp, error)
	SetPassword(context.Context, *SetPasswordReq) (*SetPasswordResp, error)
	CcnuLogin(context.Context, *CcnuLoginReq) (*CcnuLoginResp, error)
	SetStudentID(context.Context, *SetStudentIDReq) (*SetStudentIDResp, error)
	SetEmail(context.Context, *SetEmailReq) (*SetEmailResp, error)
	mustEmbedUnimplementedAccountCenterClientServer()
}

// UnimplementedAccountCenterClientServer must be embedded to have forward compatible implementations.
type UnimplementedAccountCenterClientServer struct {
}

func (UnimplementedAccountCenterClientServer) Register(context.Context, *RegisterDataReq) (*RegisterDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAccountCenterClientServer) Login(context.Context, *LoginVerifyReq) (*LoginVerifyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAccountCenterClientServer) SetPassword(context.Context, *SetPasswordReq) (*SetPasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPassword not implemented")
}
func (UnimplementedAccountCenterClientServer) CcnuLogin(context.Context, *CcnuLoginReq) (*CcnuLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CcnuLogin not implemented")
}
func (UnimplementedAccountCenterClientServer) SetStudentID(context.Context, *SetStudentIDReq) (*SetStudentIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStudentID not implemented")
}
func (UnimplementedAccountCenterClientServer) SetEmail(context.Context, *SetEmailReq) (*SetEmailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEmail not implemented")
}
func (UnimplementedAccountCenterClientServer) mustEmbedUnimplementedAccountCenterClientServer() {}

// UnsafeAccountCenterClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountCenterClientServer will
// result in compilation errors.
type UnsafeAccountCenterClientServer interface {
	mustEmbedUnimplementedAccountCenterClientServer()
}

func RegisterAccountCenterClientServer(s grpc.ServiceRegistrar, srv AccountCenterClientServer) {
	s.RegisterService(&AccountCenterClient_ServiceDesc, srv)
}

func _AccountCenterClient_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterClientServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountCenterClient_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterClientServer).Register(ctx, req.(*RegisterDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenterClient_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginVerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterClientServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountCenterClient_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterClientServer).Login(ctx, req.(*LoginVerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenterClient_SetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterClientServer).SetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountCenterClient_SetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterClientServer).SetPassword(ctx, req.(*SetPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenterClient_CcnuLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CcnuLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterClientServer).CcnuLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountCenterClient_CcnuLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterClientServer).CcnuLogin(ctx, req.(*CcnuLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenterClient_SetStudentID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStudentIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterClientServer).SetStudentID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountCenterClient_SetStudentID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterClientServer).SetStudentID(ctx, req.(*SetStudentIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenterClient_SetEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterClientServer).SetEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountCenterClient_SetEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterClientServer).SetEmail(ctx, req.(*SetEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountCenterClient_ServiceDesc is the grpc.ServiceDesc for AccountCenterClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountCenterClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accountCenter.accountCenterClient",
	HandlerType: (*AccountCenterClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AccountCenterClient_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AccountCenterClient_Login_Handler,
		},
		{
			MethodName: "SetPassword",
			Handler:    _AccountCenterClient_SetPassword_Handler,
		},
		{
			MethodName: "ccnuLogin",
			Handler:    _AccountCenterClient_CcnuLogin_Handler,
		},
		{
			MethodName: "SetStudentID",
			Handler:    _AccountCenterClient_SetStudentID_Handler,
		},
		{
			MethodName: "SetEmail",
			Handler:    _AccountCenterClient_SetEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accountCenter.proto",
}
