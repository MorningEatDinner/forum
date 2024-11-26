package logic

import (
	"context"

	"forum/app/community/rpc/internal/svc"
	"forum/app/community/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommunityPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommunityPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommunityPostsLogic {
	return &GetCommunityPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommunityPostsLogic) GetCommunityPosts(in *pb.GetCommunityPostsRequest) (*pb.GetCommunityPostsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetCommunityPostsResponse{}, nil
}
