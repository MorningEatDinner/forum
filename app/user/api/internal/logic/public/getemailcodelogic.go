package public

import (
	"context"
	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"forum/app/user/rpc/userservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get email verification code
func NewGetEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailCodeLogic {
	return &GetEmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmailCodeLogic) GetEmailCode(req *types.GetEmailCodeReq) (resp *types.GetEmailCodeResp, err error) {
	getMobilResp, err := l.svcCtx.UserRpc.GetEmailCode(l.ctx, &userservice.GetEmailCodeRequest{
		Email:       req.Email,
		CaptchaId:   req.CaptchaId,
		CaptchaCode: req.CaptchaCode,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GetMobileCode: %v", err)
		return
	}
	resp = &types.GetEmailCodeResp{}
	copier.Copy(resp, getMobilResp)

	return
}
