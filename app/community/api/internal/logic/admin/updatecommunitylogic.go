package admin

import (
	"context"

	"forum/app/community/api/internal/svc"
	"forum/app/community/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// update community details
func NewUpdateCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommunityLogic {
	return &UpdateCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCommunityLogic) UpdateCommunity(req *types.UpdateCommunityReq) (resp *types.UpdateCommunityResp, err error) {
	// todo: add your logic here and delete this line

	return
}
