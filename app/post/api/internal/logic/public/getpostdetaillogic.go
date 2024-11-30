package public

import (
	"context"
	"forum/tmp/app/post/rpc/postservice"

	"forum/app/community/rpc/communityservice"
	"forum/app/post/api/internal/svc"
	"forum/app/post/api/internal/types"
	"forum/app/user/rpc/userservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取帖子详情
func NewGetPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostDetailLogic {
	return &GetPostDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostDetailLogic) GetPostDetail(req *types.GetPostDetailReq) (resp *types.GetPostDetailResp, err error) {
	// 输入一个帖子的id， 获取帖子的详细信息
	postResp, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &postservice.GetPostDetailRequest{
		PostId: req.PostId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get post detail failed, err: %v", err)
		return nil, err
	}

	// 获取用户名
	userResp, err := l.svcCtx.UserRpc.GetUserDetail(l.ctx, &userservice.UserInfoRequest{
		UserId: postResp.Post.AuthorId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get user detail failed, err: %v", err)
		return nil, err
	}

	// 获取社区的信息
	communityResp, err := l.svcCtx.CommunityRpc.GetCommunityDetails(l.ctx, &communityservice.GetCommunityDetailsRequest{
		CommunityId: postResp.Post.CommunityId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get community detail failed, err: %v", err)
		return nil, err
	}

	// TODO: 获取当前的帖子的投票的数量
	postRet := &types.Post{}
	communityRet := &types.Community{}
	copier.Copy(postRet, postResp.Post)
	copier.Copy(communityRet, communityResp.Community)

	resp = &types.GetPostDetailResp{
		PostDetail: types.PostDetail{
			Post:       *postRet,
			Community:  *communityRet,
			AuthorName: userResp.User.Username,
		},
	}

	return
}
