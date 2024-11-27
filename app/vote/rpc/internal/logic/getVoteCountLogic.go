package logic

import (
	"context"

	"forum/app/vote/rpc/internal/svc"
	"forum/app/vote/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVoteCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVoteCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVoteCountLogic {
	return &GetVoteCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVoteCountLogic) GetVoteCount(in *pb.GetVoteCountRequest) (*pb.GetVoteCountResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetVoteCountResponse{}, nil
}
