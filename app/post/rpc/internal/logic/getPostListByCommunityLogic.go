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
	logic := NewGetPostListLogic(l.ctx, l.svcCtx)
	getResp, err := logic.GetPostList(&pb.GetPostListRequest{
		CommunityId: &in.CommunityId,
		Page:        in.Page,
		PageSize:    in.PageSize,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get post list by community failed, err: %v", err)
		return nil, err
	}

	return &pb.GetPostListByCommunityResponse{
		Posts: getResp.Posts,
		Total: getResp.Total,
	}, nil
}
