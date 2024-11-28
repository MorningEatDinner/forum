package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

type GetMobileCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const (
	PhoneCodeExpireTime = time.Minute * 15
)

func NewGetMobileCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMobileCodeLogic {
	return &GetMobileCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 向用户手机号码发送验证码
func (l *GetMobileCodeLogic) GetMobileCode(in *pb.GetMobileCodeRequest) (*pb.GetMobileCodeResponse, error) {
	// 1. 检查验证码是否匹配
	res := captcha.NewCaptcha(l.ctx, l.svcCtx.RedisClient).VerifyCaptcha(in.CaptchaId, in.CaptchaCode)
	if !res {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_CAPTCHA_ERROR), "error captcha")
	}

	// 2. 随机生成验证码
	code := helpers.GenerateRandomCode()
	// 3. 保存到Redis中
	key := fmt.Sprintf(globalkey.GetRedisKey(globalkey.PhoneCodeKey), in.Phone)
	err := l.svcCtx.RedisClient.Setex(key, code, int(PhoneCodeExpireTime.Seconds()))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to save code to database, key is %s", key)
	}
	// 4. 发送短信给用户
	// 发送消息给RabbitMQ
	if err := l.SendCode2Phone(in.Phone, code); err != nil {
		logx.WithContext(l.ctx).Errorf("send code error, key is %s", key)
		return nil, errors.Wrapf(xerr.NewErrMsg("send code error"), "send code error, key is %s", key)
	}

	// 5. 返回响应
	return &pb.GetMobileCodeResponse{}, nil
}

func (l *GetMobileCodeLogic) SendCode2Phone(phone, code string) error {
	message := mq.SendPhoneCodeMessage{
		Phone: phone,
		Code:  code,
	}
	body, err := json.Marshal(message)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("marshal message failed, err: %v", err.Error())
		return errors.Wrapf(xerr.NewErrMsg("marshal message failed"), "marshal message failed, err: %v", err.Error())
	}

	return l.svcCtx.RabbitMqClient.Send("verify_code", "send_code2phone", body)
}
