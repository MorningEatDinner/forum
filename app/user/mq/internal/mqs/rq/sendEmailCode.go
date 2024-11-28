package rq

import (
	"context"
	"encoding/json"
	"fmt"
	"forum/app/user/mq/internal/svc"
	"forum/common/mail"
	"forum/common/mq"
	"forum/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
*
Listening to the payment flow status change notification message queue
*/
type SendEmailCodeMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailCodeMq(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailCodeMq {
	return &SendEmailCodeMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailCodeMq) Consume(val string) error {
	var message mq.SendEmailCodeMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("SendEmailCodeMq->Consume json.Unmarshal err : %v, val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("SendEmailCodeMq->Consume execService err : %v, val : %s", err, val)
		return err
	}

	return nil
}

func (l *SendEmailCodeMq) execService(message mq.SendEmailCodeMessage) error {
	email := mail.Email{
		From: mail.From{
			l.svcCtx.Config.MailConf.FromConfig.Address,
			l.svcCtx.Config.MailConf.FromConfig.Name,
		},
		To:      []string{message.Email}, // 收件人地址
		Subject: "Forum 论坛验证码",           // 主题
		HTML:    []byte(fmt.Sprintf(mail.TemplateHTMLCode, message.Code)),
	}
	res := l.svcCtx.MailClient.Send(l.ctx, email)
	if !res {
		return errors.Wrapf(xerr.NewErrMsg("failed to send email"), "failed to send email")
	}
	return nil
}
