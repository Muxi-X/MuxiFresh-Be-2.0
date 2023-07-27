// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package commentclient

import (
	"context"

	"MuXiFresh-Be-2.0/app/task/cmd/rpc/comment/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Comment                  = pb.Comment
	DelSubmissionCommentReq  = pb.DelSubmissionCommentReq
	DelSubmissionCommentResp = pb.DelSubmissionCommentResp
	GetSubmissionCommentReq  = pb.GetSubmissionCommentReq
	GetSubmissionCommentResp = pb.GetSubmissionCommentResp
	IsMyCommentReq           = pb.IsMyCommentReq
	IsMyCommentResp          = pb.IsMyCommentResp
	SetSubmissionCommentReq  = pb.SetSubmissionCommentReq
	SetSubmissionCommentResp = pb.SetSubmissionCommentResp

	CommentClient interface {
		GetSubmissionComment(ctx context.Context, in *GetSubmissionCommentReq, opts ...grpc.CallOption) (*GetSubmissionCommentResp, error)
		SetSubmissionComment(ctx context.Context, in *SetSubmissionCommentReq, opts ...grpc.CallOption) (*SetSubmissionCommentResp, error)
		DelSubmissionComment(ctx context.Context, in *DelSubmissionCommentReq, opts ...grpc.CallOption) (*DelSubmissionCommentResp, error)
		IsMyComment(ctx context.Context, in *IsMyCommentReq, opts ...grpc.CallOption) (*IsMyCommentResp, error)
	}

	defaultCommentClient struct {
		cli zrpc.Client
	}
)

func NewCommentClient(cli zrpc.Client) CommentClient {
	return &defaultCommentClient{
		cli: cli,
	}
}

func (m *defaultCommentClient) GetSubmissionComment(ctx context.Context, in *GetSubmissionCommentReq, opts ...grpc.CallOption) (*GetSubmissionCommentResp, error) {
	client := pb.NewCommentClientClient(m.cli.Conn())
	return client.GetSubmissionComment(ctx, in, opts...)
}

func (m *defaultCommentClient) SetSubmissionComment(ctx context.Context, in *SetSubmissionCommentReq, opts ...grpc.CallOption) (*SetSubmissionCommentResp, error) {
	client := pb.NewCommentClientClient(m.cli.Conn())
	return client.SetSubmissionComment(ctx, in, opts...)
}

func (m *defaultCommentClient) DelSubmissionComment(ctx context.Context, in *DelSubmissionCommentReq, opts ...grpc.CallOption) (*DelSubmissionCommentResp, error) {
	client := pb.NewCommentClientClient(m.cli.Conn())
	return client.DelSubmissionComment(ctx, in, opts...)
}

func (m *defaultCommentClient) IsMyComment(ctx context.Context, in *IsMyCommentReq, opts ...grpc.CallOption) (*IsMyCommentResp, error) {
	client := pb.NewCommentClientClient(m.cli.Conn())
	return client.IsMyComment(ctx, in, opts...)
}
