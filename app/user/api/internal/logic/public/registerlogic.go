package public

import (
	"context"
	"forum/common/helpers"
	"forum/common/xerr"
	"github.com/pkg/errors"

	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"forum/app/user/rpc/userservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// register
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 检查各个必填字段
	switch {
	case req.Phone == "":
		return nil, errors.Wrapf(xerr.NewErrMsg("invalid parameter"), "email cannot be empty")
	case req.Code == "":
		return nil, errors.Wrapf(xerr.NewErrMsg("invalid parameter"), "verification code cannot be empty")
	case req.Password == "":
		return nil, errors.Wrapf(xerr.NewErrMsg("invalid parameter"), "password cannot be empty")
	case req.PasswordConfirm == "":
		return nil, errors.Wrapf(xerr.NewErrMsg("invalid parameter"), "password confirmation cannot be empty")
	}

	// 如果用户名为空，生成随机用户名
	if req.Name == "" {
		req.Name = helpers.GenerateRandomCode()
	}
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &userservice.RegisterRequest{
		Phone:           &req.Phone,
		Password:        req.Password,
		Code:            &req.Code,
		Name:            req.Name,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		logx.WithContext(l.ctx).Error("注册失败", err.Error())
		return
	}
	resp = &types.RegisterResp{}
	copier.Copy(resp, registerResp)
	return
}
