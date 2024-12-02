package logic

import (
	"context"
	"forum/app/mqueue/cmd/job/internal/svc"
	"forum/app/post/rpc/postservice"
	"forum/common/xerr"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

type DeletePostHandler struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostHandler(svcCtx *svc.ServiceContext) *DeletePostHandler {
	return &DeletePostHandler{
		ctx:    context.Background(),
		svcCtx: svcCtx,
	}
}

// 这里处理延迟的那个任务
func (l *DeletePostHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	_, err := l.svcCtx.PostRpc.DeletePostScheduler(l.ctx, &postservice.DeletePostSchedulerRequest{})
	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg("failed to delete post"), "failed to delete post")
	}

	return nil
}
