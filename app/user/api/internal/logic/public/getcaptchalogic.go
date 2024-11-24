package public

import (
	"context"

	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"forum/app/user/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get captcha
func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取图形验证码
func (l *GetCaptchaLogic) GetCaptcha(req *types.CaptchaReq) (resp *types.CaptchaResp, err error) {
	captchaResp, err := l.svcCtx.UserRpc.GetCaptcha(l.ctx, &pb.CaptchaRequest{})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GetCaptcha failed: %v", err)
		return
	}
	resp = &types.CaptchaResp{}
	copier.Copy(resp, captchaResp)

	return
}
