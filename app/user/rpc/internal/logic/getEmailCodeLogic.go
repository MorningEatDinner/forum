package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"
	"forum/common/captcha"
	"forum/common/globalkey"
	"forum/common/helpers"
	"forum/common/mq"
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
	// 发送消息给RabbitMQ
	if err := l.SendCode2Email(in.Email, code); err != nil {
		logx.WithContext(l.ctx).Errorf("send code error, key is %s", key)
	}

	// 5. 返回响应
	return &pb.GetEmailCodeResponse{}, nil
}

// 发送邮件
func (l *GetEmailCodeLogic) SendCode2Email(email, code string) error {
	message := mq.SendEmailCodeMessage{
		Email: email,
		Code:  code,
	}
	body, err := json.Marshal(message)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("marshal message failed, err: %v", err.Error())
		return errors.Wrapf(xerr.NewErrMsg("marshal message failed"), "marshal message failed, err: %v", err.Error())
	}

	return l.svcCtx.RabbitMqClient.Send("verify_code", "send_code2email", body)
}
