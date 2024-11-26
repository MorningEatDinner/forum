package public

import (
	"context"

	"forum/app/community/api/internal/svc"
	"forum/app/community/api/internal/types"
	"forum/app/community/rpc/communityservice"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
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
	var page, pageSize *int32
	if req.Page != 0 {
		pageVal := int32(req.Page)
		page = &pageVal
	}
	if req.PageSize != 0 {
		pageSizeVal := int32(req.PageSize)
		pageSize = &pageSizeVal
	}
	getResp, err := l.svcCtx.CommunityRpc.GetAllCommunities(l.ctx, &communityservice.GetAllCommunitiesRequest{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get all communities failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("get all communities failed"), "get all communities failed")
	}

	resp = &types.GetAllCommunitiesResp{}
	copier.Copy(resp, getResp)
	return
}
