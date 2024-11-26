package admin

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

type CreateCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// create a new community
func NewCreateCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommunityLogic {
	return &CreateCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommunityLogic) CreateCommunity(req *types.CreateCommunityReq) (resp *types.CreateCommunityResp, err error) {
	communityResp, err := l.svcCtx.CommunityRpc.CreateCommunity(l.ctx, &communityservice.CreateCommunityRequest{
		CommunityName: req.CommunityName,
		Introduction:  &req.Introduction,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create community failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "create community failed, err: %v", err)
	}
	resp = &types.CreateCommunityResp{}
	copier.Copy(resp, communityResp)

	return
}
