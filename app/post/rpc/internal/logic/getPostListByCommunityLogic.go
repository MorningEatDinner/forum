package logic

import (
	"context"

	"forum/app/post/rpc/internal/svc"
	"forum/app/post/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostListByCommunityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostListByCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostListByCommunityLogic {
	return &GetPostListByCommunityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostListByCommunityLogic) GetPostListByCommunity(in *pb.GetPostListByCommunityRequest) (*pb.GetPostListByCommunityResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetPostListByCommunityResponse{}, nil
}
