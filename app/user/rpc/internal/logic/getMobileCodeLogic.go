package logic

import (
	"context"
	"fmt"
	"time"

	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"
	"forum/common/captcha"
	"forum/common/globalkey"
	"forum/common/helpers"
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
	PhoneCodeExpireTime = time.Minute * 5
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
	_, err := l.svcCtx.RedisClient.SetnxExCtx(l.ctx, key, code, int(PhoneCodeExpireTime.Seconds()))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to save code to database, key is %s", key)
	}
	// 4. 发送短信给用户
	// TODO: 这里可以修改成为使用消息队列的形式
	res = l.svcCtx.SMSClient.Send(l.ctx, in.Phone, code)
	if !res {
		return nil, errors.Wrapf(xerr.NewErrMsg("send code error"), "send code error, key is %s", key)
	}

	// 5. 返回响应
	return &pb.GetMobileCodeResponse{}, nil
}
