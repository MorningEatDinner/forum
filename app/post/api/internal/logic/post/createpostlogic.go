package post

import (
	"context"

	"forum/app/community/rpc/communityservice"
	"forum/app/post/api/internal/svc"
	"forum/app/post/api/internal/types"
	"forum/app/post/rpc/postservice"
	"forum/common/ctxdata"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建帖子
func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(req *types.CreatePostReq) (resp *types.CreatePostResp, err error) {
	// 获取当前的用户id
	userid := ctxdata.GetUidFromCtx(l.ctx)

	// 验证当前给定社区是否存在
	_, err = l.svcCtx.CommunityRpc.GetCommunityDetails(l.ctx, &communityservice.GetCommunityDetailsRequest{
		CommunityId: req.CommunityId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get community details failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.COMMUNITY_NAME_EXIST), "community not exist")
	}

	// 调用rpc接口
	createResp, err := l.svcCtx.PostRpc.CreatePost(l.ctx, &postservice.CreatePostRequest{
		AuthorId:    userid,
		CommunityId: req.CommunityId,
		Title:       req.Title,
		Content:     req.Content,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create post failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("create post failed"), "create post failed")
	}

	resp = &types.CreatePostResp{}
	copier.Copy(resp, createResp)

	return
}
