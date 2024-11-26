package logic

import (
	"context"

	"forum/app/post/rpc/internal/svc"
	"forum/app/post/rpc/pb"

	"github.com/jinzhu/copier"
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
	// 输入一个帖子id， 获取帖子的详细信息
	postResp, err := l.svcCtx.PostModel.FindOne(l.ctx, in.PostId)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get post detail failed, err: %v", err)
		return nil, err
	}

	postRet := &pb.Post{}
	copier.Copy(postRet, postResp)

	return &pb.GetPostDetailResponse{
		Post: postRet,
	}, nil
}
