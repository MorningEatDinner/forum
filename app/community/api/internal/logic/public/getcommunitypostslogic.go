package public

import (
	"context"

	"forum/app/community/api/internal/svc"
	"forum/app/community/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
