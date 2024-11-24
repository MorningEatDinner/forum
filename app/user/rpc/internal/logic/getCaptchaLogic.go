package logic

import (
	"context"

	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"
	"forum/common/captcha"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCaptchaLogic) GetCaptcha(in *pb.CaptchaRequest) (*pb.CaptchaResponse, error) {
	// 生成验证码
	id, b64s, _, err := captcha.NewCaptcha(l.ctx, l.svcCtx.RedisClient).GenerateCaptcha()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GenerateCaptcha failed: %v", err)
		return nil, err
	}

	return &pb.CaptchaResponse{
		CaptchaId:   id,
		ImageBase64: b64s,
	}, nil
}
