package profile

import (
	"context"

	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"forum/app/user/rpc/userservice"
	"forum/common/ctxdata"

	"github.com/jinzhu/copier"
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
	userId := ctxdata.GetUidFromCtx(l.ctx)
	updatedResp, err := l.svcCtx.UserRpc.UpdateMobile(l.ctx, &userservice.UpdateMobileRequest{
		UserId:   userId,
		NewPhone: req.NewPhone,
		Code:     req.Code,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.UpdateMobileResp{}
	copier.Copy(resp, updatedResp)

	return
}
