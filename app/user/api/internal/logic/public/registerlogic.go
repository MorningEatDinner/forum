package public

import (
	"context"

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
