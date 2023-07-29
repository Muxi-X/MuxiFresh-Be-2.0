// Code generated by goctl. DO NOT EDIT.
// Source: submission.proto

package submissionclient

import (
	"context"

	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Completion                 = pb.Completion
	GetAllSubmissionStatusReq  = pb.GetAllSubmissionStatusReq
	GetAllSubmissionStatusResp = pb.GetAllSubmissionStatusResp
	GetMySubmissionStatusReq   = pb.GetMySubmissionStatusReq
	GetMySubmissionStatusResp  = pb.GetMySubmissionStatusResp
	GetSubmissionInfoReq       = pb.GetSubmissionInfoReq
	GetSubmissionInfoResp      = pb.GetSubmissionInfoResp
	SetSubmissionReq           = pb.SetSubmissionReq
	SetSubmissionResp          = pb.SetSubmissionResp

	SubmissionClient interface {
		GetMySubmissionStatus(ctx context.Context, in *GetMySubmissionStatusReq, opts ...grpc.CallOption) (*GetMySubmissionStatusResp, error)
		SetSubmission(ctx context.Context, in *SetSubmissionReq, opts ...grpc.CallOption) (*SetSubmissionResp, error)
		GetSubmissionInfo(ctx context.Context, in *GetSubmissionInfoReq, opts ...grpc.CallOption) (*GetSubmissionInfoResp, error)
		GetAllSubmissionStatus(ctx context.Context, in *GetAllSubmissionStatusReq, opts ...grpc.CallOption) (*GetAllSubmissionStatusResp, error)
	}

	defaultSubmissionClient struct {
		cli zrpc.Client
	}
)

func NewSubmissionClient(cli zrpc.Client) SubmissionClient {
	return &defaultSubmissionClient{
		cli: cli,
	}
}

func (m *defaultSubmissionClient) GetMySubmissionStatus(ctx context.Context, in *GetMySubmissionStatusReq, opts ...grpc.CallOption) (*GetMySubmissionStatusResp, error) {
	client := pb.NewSubmissionClientClient(m.cli.Conn())
	return client.GetMySubmissionStatus(ctx, in, opts...)
}

func (m *defaultSubmissionClient) SetSubmission(ctx context.Context, in *SetSubmissionReq, opts ...grpc.CallOption) (*SetSubmissionResp, error) {
	client := pb.NewSubmissionClientClient(m.cli.Conn())
	return client.SetSubmission(ctx, in, opts...)
}

func (m *defaultSubmissionClient) GetSubmissionInfo(ctx context.Context, in *GetSubmissionInfoReq, opts ...grpc.CallOption) (*GetSubmissionInfoResp, error) {
	client := pb.NewSubmissionClientClient(m.cli.Conn())
	return client.GetSubmissionInfo(ctx, in, opts...)
}

func (m *defaultSubmissionClient) GetAllSubmissionStatus(ctx context.Context, in *GetAllSubmissionStatusReq, opts ...grpc.CallOption) (*GetAllSubmissionStatusResp, error) {
	client := pb.NewSubmissionClientClient(m.cli.Conn())
	return client.GetAllSubmissionStatus(ctx, in, opts...)
}