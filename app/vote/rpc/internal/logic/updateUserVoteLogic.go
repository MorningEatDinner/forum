package logic

import (
	"context"

	"forum/app/vote/rpc/internal/svc"
	"forum/app/vote/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserVoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserVoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserVoteLogic {
	return &UpdateUserVoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserVoteLogic) UpdateUserVote(in *pb.UpdateUserVoteRequest) (*pb.UpdateUserVoteResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserVoteResponse{}, nil
}
