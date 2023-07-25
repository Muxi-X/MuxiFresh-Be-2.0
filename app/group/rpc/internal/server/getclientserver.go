// Code generated by goctl. DO NOT EDIT.
// Source: intro.proto

package server

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/rpc/internal/logic"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/rpc/internal/svc"
	pb2 "MuxiFresh2.0/MuXiFresh-Be-2.0/app/group/rpc/pb"
	"context"
)

type GetClientServer struct {
	svcCtx *svc.ServiceContext
	pb2.UnimplementedGetClientServer
}

func NewGetClientServer(svcCtx *svc.ServiceContext) *GetClientServer {
	return &GetClientServer{
		svcCtx: svcCtx,
	}
}

func (s *GetClientServer) GetGroupIntro(ctx context.Context, in *pb2.GetReq) (*pb2.GetResp, error) {
	l := logic.NewGetGroupIntroLogic(ctx, s.svcCtx)
	return l.GetGroupIntro(in)
}
