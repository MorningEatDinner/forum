package public

import (
	"context"

	"forum/app/community/rpc/communityservice"
	"forum/app/post/api/internal/svc"
	"forum/app/post/api/internal/types"
	"forum/app/post/rpc/postservice"
	"forum/app/user/rpc/userservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取帖子列表
func NewGetPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostListLogic {
	return &GetPostListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostListLogic) GetPostList(req *types.GetPostListReq) (resp *types.GetPostListResp, err error) {
	var communityId, authorId *int64
	if req.CommunityId != 0 {
		// 需要验证communityId是否存在
		_, err := l.svcCtx.CommunityRpc.GetCommunityDetails(l.ctx, &communityservice.GetCommunityDetailsRequest{
			CommunityId: req.CommunityId,
		})
		if err != nil {
			logx.WithContext(l.ctx).Errorf("get community details failed, err: %v", err)
			return nil, err
		}
		communityIdVal := int64(req.CommunityId)
		communityId = &communityIdVal
	}
	if req.AuthorId != 0 {
		_, err := l.svcCtx.UserRpc.GetUserDetail(l.ctx, &userservice.UserInfoRequest{
			UserId: req.AuthorId,
		})
		if err != nil {
			logx.WithContext(l.ctx).Errorf("get user details failed, err: %v", err)
			return nil, err
		}
		authorIdVal := int64(req.AuthorId)
		authorId = &authorIdVal
	}
	listResp, err := l.svcCtx.PostRpc.GetPostList(l.ctx, &postservice.GetPostListRequest{
		Page:        req.Page,
		PageSize:    req.PageSize,
		CommunityId: communityId,
		AuthorId:    authorId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get post list failed, err: %v", err)
		return nil, err
	}

	resp = &types.GetPostListResp{}
	copier.Copy(resp, listResp)

	return
}
