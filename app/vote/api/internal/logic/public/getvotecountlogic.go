package public

import (
	"context"

	"forum/app/vote/api/internal/svc"
	"forum/app/vote/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVoteCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取帖子投票信息
func NewGetVoteCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVoteCountLogic {
	return &GetVoteCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// rpc实现就好
func (l *GetVoteCountLogic) GetVoteCount(req *types.GetVoteCountReq) (resp *types.GetVoteCountResp, err error) {
	// todo: add your logic here and delete this line

	return
}
