package vote

import (
	"context"

	"forum/app/vote/api/internal/svc"
	"forum/app/vote/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RevokeVoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 撤销用户对帖子的投票
func NewRevokeVoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RevokeVoteLogic {
	return &RevokeVoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RevokeVoteLogic) RevokeVote(req *types.RevokeVoteReq) (resp *types.RevokeVoteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
