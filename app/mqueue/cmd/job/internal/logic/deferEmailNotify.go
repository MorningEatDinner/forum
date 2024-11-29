package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"forum/app/mqueue/cmd/job/internal/svc"
	"forum/app/mqueue/cmd/job/jobtype"
	"forum/common/globalkey"
	"forum/common/mail"
	"forum/common/xerr"
	"strconv"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type NotifyUserUpdateHandler struct {
	svcCtx *svc.ServiceContext
}

func NewNotifyUserUpdateHandler(svcCtx *svc.ServiceContext) *NotifyUserUpdateHandler {
	return &NotifyUserUpdateHandler{
		svcCtx: svcCtx,
	}
}

// 这里处理延迟的那个任务
func (l *NotifyUserUpdateHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	var p jobtype.DeferNotifyUserPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(xerr.NewErrMsg("unmarshal payload failed"), "unmarshal payload failed")
	}

	// 1. 先看redis中， 这个用户是否已经更新完成个人信息了
	status, err := l.svcCtx.RedisClient.Get(fmt.Sprintf(globalkey.GetRedisKey(globalkey.UpdatedKey), strconv.FormatInt(p.UserId, 10)))
	if err != nil && err != redis.Nil {
		return errors.Wrapf(xerr.NewErrMsg("failed to get user updated info"), "failed to get user updated info")
	}
	if status == "1" {
		return nil
	}

	email := mail.Email{
		From: mail.From{
			l.svcCtx.Config.MailConf.FromConfig.Address,
			l.svcCtx.Config.MailConf.FromConfig.Name,
		},
		To:      []string{p.Email},              // 收件人地址
		Subject: "Forum 论坛邀请您更新个人信息", // 主题
		HTML:    []byte(fmt.Sprintf(mail.TemplateHTMLProfileUpdate, p.UserName)),
	}
	res := l.svcCtx.MailClient.Send(ctx, email)
	if !res {
		return errors.Wrapf(xerr.NewErrMsg("failed to send email"), "failed to send email")
	}

	// 如果发送成功了， 把状态数据从Redis中删除
	l.svcCtx.RedisClient.Hdel(globalkey.GetRedisKey(globalkey.UpdatedKey), strconv.FormatInt(p.UserId, 10))

	return nil
}
