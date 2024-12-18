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

type UpdateInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// update user info
func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInfoLogic) UpdateInfo(req *types.UpdateUserInfoReq) (resp *types.UpdateUserInfoResp, err error) {
	// 获取当前userid
	userId := ctxdata.GetUidFromCtx(l.ctx)
	updatedResp, err := l.svcCtx.UserRpc.UpdateUserInfo(l.ctx, &userservice.UpdateUserInfoRequest{
		UserId:       userId,
		Username:     &req.Username,
		City:         &req.City,
		Introduction: &req.Introduction,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.UpdateUserInfoResp{}
	copier.Copy(resp, updatedResp)

	return
}
