package public

import (
	"context"

	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"forum/app/user/rpc/pb"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// login
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	logResp, err := l.svcCtx.UserRpc.Login(l.ctx, &pb.LoginRequest{
		Username: &req.Username,
		Email:    &req.Email,
		Phone:    &req.Phone,
		Password: req.Password,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to login: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to login"), "failed to login: %v", err)
	}

	resp = &types.LoginResp{}
	copier.Copy(resp, logResp)
	return
}
