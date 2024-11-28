package rq

import (
	"context"
	"encoding/json"
	"forum/app/user/mq/internal/svc"
	"forum/common/mq"
	"forum/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
*
Listening to the payment flow status change notification message queue
*/
type SendPhoneCodeMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendPhoneCodeMq(ctx context.Context, svcCtx *svc.ServiceContext) *SendPhoneCodeMq {
	return &SendPhoneCodeMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendPhoneCodeMq) Consume(val string) error {

	var message mq.SendPhoneCodeMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("SendPhoneCodeMq->Consume Unmarshal err : %v, val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("SendPhoneCodeMq->Consume execService err : %v, val : %s", err, val)
		return err
	}

	return nil
}

func (l *SendPhoneCodeMq) execService(message mq.SendPhoneCodeMessage) error {
	// 执行发送短信的功能
	res := l.svcCtx.SMSClient.Send(l.ctx, message.Phone, message.Code)
	if !res {
		logx.WithContext(l.ctx).Errorf("send code error")
		return errors.Wrapf(xerr.NewErrMsg("send code error"), "send code error")
	}
	return nil
}
