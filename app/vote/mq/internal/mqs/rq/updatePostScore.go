package rq

import (
	"context"
	"encoding/json"
	"forum/app/post/rpc/postservice"
	"forum/app/vote/mq/internal/svc"
	"forum/common/mq"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
*
Listening to the payment flow status change notification message queue
*/
type UpdatePostScoreMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePostScoreMq(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostScoreMq {
	return &UpdatePostScoreMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePostScoreMq) Consume(val string) error {
	var message mq.UpdatePostScore
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.Errorf("unmarshal message failed, err: %v", err.Error())
		return errors.Wrapf(xerr.NewErrMsg("unmarshal message failed"), "unmarshal message failed, err: %v", err.Error())
	}

	if err := l.execService(message); err != nil {
		logx.Errorf("exec service failed, err: %v", err.Error())
		return errors.Wrapf(xerr.NewErrMsg("exec service failed"), "exec service failed, err: %v", err.Error())
	}

	return nil
}

func (l *UpdatePostScoreMq) execService(message mq.UpdatePostScore) error {
	var req postservice.UpdatePostScoreRequest
	copier.Copy(&req, &message)
	resp, err := l.svcCtx.PostRpc.UpdatePostScore(l.ctx, &req)
	if err != nil {
		logx.Errorf("update post score failed, err: %v", err.Error())
		return err
	}
	logx.Infof("update post score success, resp: %v", resp)

	return nil
}
