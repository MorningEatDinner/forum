package vote

import (
	"context"

	"forum/app/vote/api/internal/svc"
	"forum/app/vote/api/internal/types"
	"forum/app/vote/rpc/voteservice"
	"forum/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type VotePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 为帖子投票
func NewVotePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VotePostLogic {
	return &VotePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VotePostLogic) VotePost(req *types.VotePostReq) (*types.VotePostResp, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.VoteRpc.VotePost(l.ctx, &voteservice.VotePostRequest{
		UserId:   userId,
		PostId:   req.PostId,
		VoteType: req.VoteType,
	})
	if err != nil {
		return nil, err
	}

	return &types.VotePostResp{Success: resp.Success}, nil
}
