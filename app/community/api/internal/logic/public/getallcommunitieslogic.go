package public

import (
	"context"

	"forum/app/community/api/internal/svc"
	"forum/app/community/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllCommunitiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get all communities
func NewGetAllCommunitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllCommunitiesLogic {
	return &GetAllCommunitiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllCommunitiesLogic) GetAllCommunities(req *types.GetAllCommunitiesReq) (resp *types.GetAllCommunitiesResp, err error) {
	// todo: add your logic here and delete this line

	return
}
