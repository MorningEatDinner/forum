package public

import (
	"context"
	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"forum/app/user/rpc/userservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterByEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// register by email
func NewRegisterByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterByEmailLogic {
	return &RegisterByEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterByEmailLogic) RegisterByEmail(req *types.RegisterByEmailReq) (resp *types.RegisterByEmailResp, err error) {
	registerResp, err := l.svcCtx.UserRpc.RegisterByEmail(l.ctx, &userservice.RegisterByEmailRequest{
		Email:           req.Email,
		Code:            req.Code,
		Name:            req.Name,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		logx.WithContext(l.ctx).Error("注册失败", err.Error())
		return
	}
	resp = &types.RegisterByEmailResp{}
	copier.Copy(resp, registerResp)
	return
}
