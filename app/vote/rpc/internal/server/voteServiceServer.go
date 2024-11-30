// Code generated by goctl. DO NOT EDIT.
// Source: vote.proto

package server

import (
	"context"

	"forum/app/vote/rpc/internal/logic"
	"forum/app/vote/rpc/internal/svc"
	"forum/app/vote/rpc/pb"
)

type VoteServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedVoteServiceServer
}

func NewVoteServiceServer(svcCtx *svc.ServiceContext) *VoteServiceServer {
	return &VoteServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *VoteServiceServer) VotePost(ctx context.Context, in *pb.VotePostRequest) (*pb.VotePostResponse, error) {
	l := logic.NewVotePostLogic(ctx, s.svcCtx)
	return l.VotePost(in)
}


func (s *VoteServiceServer) GetUserVote(ctx context.Context, in *pb.GetUserVoteRequest) (*pb.GetUserVoteResponse, error) {
	l := logic.NewGetUserVoteLogic(ctx, s.svcCtx)
	return l.GetUserVote(in)
}
