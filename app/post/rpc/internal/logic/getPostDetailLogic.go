package logic

import (
	"context"

	"forum/app/post/rpc/internal/svc"
	"forum/app/post/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostDetailLogic {
	return &GetPostDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostDetailLogic) GetPostDetail(in *pb.GetPostDetailRequest) (*pb.GetPostDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetPostDetailResponse{}, nil
}
