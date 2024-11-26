package logic

import (
	"context"

	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"
	"forum/common/tool"
	"forum/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(in *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	// 1. 获取当前用户信息
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to get user information: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to get user information: %v", err)
	}

	// 2. 确认密码是否正确
	if !tool.CheckPasswordHash(in.OldPassword, user.Password) {
		logx.WithContext(l.ctx).Errorf("failed to check password: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_PASSWORD_ERROR), "failed to check password: %v", err)
	}

	// 3. 修改密码
	user.Password = tool.EncryptPassword(in.NewPassword)

	// 5. 返回用户信息
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to update user information: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to update user information: %v", err)
	}

	return &pb.UpdatePasswordResponse{}, nil
}
