package vote

import (
	"context"

	"forum/app/vote/api/internal/svc"
	"forum/app/vote/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserVoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户对帖子投票的记录
func NewGetUserVoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserVoteLogic {
	return &GetUserVoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 接口无需实现， rpc实现
func (l *GetUserVoteLogic) GetUserVote(req *types.GetUserVoteReq) (resp *types.GetUserVoteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
