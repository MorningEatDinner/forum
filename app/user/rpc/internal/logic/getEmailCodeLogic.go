package logic

import (
	"context"
	"fmt"
	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"
	"forum/common/captcha"
	"forum/common/globalkey"
	"forum/common/helpers"
	"forum/common/mail"
	"forum/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailCodeLogic {
	return &GetEmailCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmailCodeLogic) GetEmailCode(in *pb.GetEmailCodeRequest) (*pb.GetEmailCodeResponse, error) {
	// 1. 检查验证码是否匹配
	res := captcha.NewCaptcha(l.ctx, l.svcCtx.RedisClient).VerifyCaptcha(in.CaptchaId, in.CaptchaCode)
	if !res {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_CAPTCHA_ERROR), "error captcha")
	}

	// 2. 随机生成验证码
	code := helpers.GenerateRandomCode()
	// 3. 保存到Redis中
	key := fmt.Sprintf(globalkey.GetRedisKey(globalkey.EmailCodeKey), in.Email)
	err := l.svcCtx.RedisClient.Setex(key, code, int(PhoneCodeExpireTime.Seconds()))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to save code to database, key is %s", key)
	}

	// 4. 发送短信给用户
	// TODO: 这里可以修改成为使用消息队列的形式
	// TODO: 这里确实需要使用消息队列的形式，  因为发送邮件需要时间
	email := mail.Email{
		From: mail.From{
			l.svcCtx.Config.MailConf.FromConfig.Address,
			l.svcCtx.Config.MailConf.FromConfig.Name,
		},
		To:      []string{in.Email}, // 收件人地址
		Subject: "Forum 论坛验证码",      // 主题
		HTML:    []byte(fmt.Sprintf(mail.TemplateHTML, code)),
	}
	res = l.svcCtx.MailClient.Send(l.ctx, email)
	if !res {
		logx.WithContext(l.ctx).Errorf("send code error, key is %s", key)
		return nil, errors.Wrapf(xerr.NewErrMsg("send code error"), "send code error, key is %s", key)
	}

	// 5. 返回响应
	return &pb.GetEmailCodeResponse{}, nil
}
