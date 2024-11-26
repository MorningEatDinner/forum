package logic

import (
	"context"

	"forum/app/post/rpc/internal/svc"
	"forum/app/post/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostListLogic {
	return &GetPostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostListLogic) GetPostList(in *pb.GetPostListRequest) (*pb.GetPostListResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetPostListResponse{}, nil
}
