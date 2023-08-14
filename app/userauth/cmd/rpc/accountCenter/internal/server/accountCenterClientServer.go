// Code generated by goctl. DO NOT EDIT.
// Source: accountCenter.proto

package server

import (
	"context"

	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/accountCenter/internal/logic"
	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/accountCenter/internal/svc"
	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/accountCenter/pb"
)

type AccountCenterClientServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedAccountCenterClientServer
}

func NewAccountCenterClientServer(svcCtx *svc.ServiceContext) *AccountCenterClientServer {
	return &AccountCenterClientServer{
		svcCtx: svcCtx,
	}
}

func (s *AccountCenterClientServer) Register(ctx context.Context, in *pb.RegisterDataReq) (*pb.RegisterDataResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *AccountCenterClientServer) Login(ctx context.Context, in *pb.LoginVerifyReq) (*pb.LoginVerifyResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *AccountCenterClientServer) SetPassword(ctx context.Context, in *pb.SetPasswordReq) (*pb.SetPasswordResp, error) {
	l := logic.NewSetPasswordLogic(ctx, s.svcCtx)
	return l.SetPassword(in)
}

func (s *AccountCenterClientServer) CcnuLogin(ctx context.Context, in *pb.CcnuLoginReq) (*pb.CcnuLoginResp, error) {
	l := logic.NewCcnuLoginLogic(ctx, s.svcCtx)
	return l.CcnuLogin(in)
}

func (s *AccountCenterClientServer) SetStudentID(ctx context.Context, in *pb.SetStudentIDReq) (*pb.SetStudentIDResp, error) {
	l := logic.NewSetStudentIDLogic(ctx, s.svcCtx)
	return l.SetStudentID(in)
}

func (s *AccountCenterClientServer) SetEmail(ctx context.Context, in *pb.SetEmailReq) (*pb.SetEmailResp, error) {
	l := logic.NewSetEmailLogic(ctx, s.svcCtx)
	return l.SetEmail(in)
}
