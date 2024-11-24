package profile

import (
	"context"

	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMobileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// update mobile
func NewUpdateMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMobileLogic {
	return &UpdateMobileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMobileLogic) UpdateMobile(req *types.UpdateMobileReq) (resp *types.UpdateMobileResp, err error) {
	// todo: add your logic here and delete this line

	return
}
