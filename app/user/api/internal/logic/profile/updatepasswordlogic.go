package profile

import (
	"context"

	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"forum/app/user/rpc/userservice"
	"forum/common/ctxdata"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// update password
func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UpdatePasswordReq) (resp *types.UpdatePasswordResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	// 确认两个密码是否匹配
	if req.NewPassword != req.ConfirmPassword {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_PASSWORD_DISMATCH_ERROR), "passwords do not match")
	}
	updateResp, err := l.svcCtx.UserRpc.UpdatePassword(l.ctx, &userservice.UpdatePasswordRequest{
		UserId:      userId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to update password: %v", err)
		return nil, err
	}
	resp = &types.UpdatePasswordResp{}

	copier.Copy(resp, updateResp)
	return
}
