package rq

import (
	"context"
	"encoding/json"
	"forum/app/post/mq/internal/svc"
	"forum/common/globalkey"
	"forum/common/mq"
	"forum/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
*
Listening to the payment flow status change notification message queue
*/
type DeletePostDownMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostDownMq(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostDownMq {
	return &DeletePostDownMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostDownMq) Consume(val string) error {
	var message mq.DeletePostMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("SendEmailCodeMq->Consume json.Unmarshal err : %v, val : %s", err, val)
		return err
	}

	if err := l.execService(message.PostId); err != nil {
		logx.WithContext(l.ctx).Error("SendEmailCodeMq->Consume execService err : %v, val : %s", err, val)
		return err
	}

	return nil
}

func (l *DeletePostDownMq) execService(pid int64) error {
	key := globalkey.GetRedisKey(globalkey.PostDownCountKey)
	_, err := l.svcCtx.RedisClient.Zrem(key, pid)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Redis Zrem error %v", err)
		return errors.Wrapf(xerr.NewErrMsg("Redis Zrem error"), "Redis Zrem error %v", err)
	}

	return nil
}
