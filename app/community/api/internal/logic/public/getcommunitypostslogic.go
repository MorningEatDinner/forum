package public

import (
	"context"

	"forum/app/community/api/internal/svc"
	"forum/app/community/api/internal/types"
	"forum/app/community/rpc/communityservice"
	"forum/app/post/rpc/postservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommunityPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get all posts in a community
func NewGetCommunityPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommunityPostsLogic {
	return &GetCommunityPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommunityPostsLogic) GetCommunityPosts(req *types.GetCommunityPostsReq) (resp *types.GetCommunityPostsResp, err error) {
	// 获取某个社区下的所有帖子
	// 1. 查看社区是否存在
	_, err = l.svcCtx.CommunityRpc.GetCommunityDetails(l.ctx, &communityservice.GetCommunityDetailsRequest{
		CommunityId: req.CommunityId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get community details failed, err: %v", err)
		return nil, err
	}

	// 2. 获取社区下的所有帖子
	listResp, err := l.svcCtx.PostRpc.GetPostListByCommunity(l.ctx, &postservice.GetPostListByCommunityRequest{
		CommunityId: req.CommunityId,
		Page:        int64(req.Page),
		PageSize:    int64(req.PageSize),
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get post list by community failed, err: %v", err)
		return nil, err
	}

	resp = &types.GetCommunityPostsResp{}
	copier.Copy(resp, listResp)

	return
}
