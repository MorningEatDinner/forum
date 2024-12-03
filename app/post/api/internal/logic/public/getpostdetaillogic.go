package public

import (
	"context"
	"forum/app/post/rpc/postservice"

	"forum/app/community/rpc/communityservice"
	"forum/app/post/api/internal/svc"
	"forum/app/post/api/internal/types"

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

	// 获取社区的信息
	communityResp, err := l.svcCtx.CommunityRpc.GetCommunityDetails(l.ctx, &communityservice.GetCommunityDetailsRequest{
		CommunityId: postResp.Post.CommunityId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get community detail failed, err: %v", err)
		return nil, err
	}

	communityRet := &types.Community{}
	copier.Copy(communityRet, communityResp.Community)
	postRet := &types.Post{}
	copier.Copy(postRet, postResp.Post)
	AuthorRet := &types.AuthorInfo{}
	logx.WithContext(l.ctx).Infof("postResp: %v", postResp.UserInfo)
	copier.Copy(AuthorRet, postResp.UserInfo)

	resp = &types.GetPostDetailResp{
		PostDetail: types.PostDetail{
			Post:       *postRet,
			Community:  *communityRet,
			AuthorInfo: *AuthorRet,
			VotedInfo: types.VotedInfo{
				UpCount:   postResp.UpCount,
				DownCount: postResp.DownCount,
			},
		},
	}

	return
}
