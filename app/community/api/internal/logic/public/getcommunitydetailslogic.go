package public

import (
	"context"

	"forum/app/community/api/internal/svc"
	"forum/app/community/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommunityDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get community details
func NewGetCommunityDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommunityDetailsLogic {
	return &GetCommunityDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 不需要写
func (l *GetCommunityDetailsLogic) GetCommunityDetails(req *types.GetCommunityDetailsReq) (resp *types.GetCommunityDetailsResp, err error) {
	// todo: add your logic here and delete this line
	logx.WithContext(l.ctx).Infof("req: %v", req)
	return
}
