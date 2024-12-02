package logic

import (
	"context"
	"encoding/json"
	"forum/common/globalkey"
	"forum/common/mq"
	"forum/common/xerr"
	"github.com/pkg/errors"
	"strconv"
	"time"

	"forum/app/post/rpc/internal/svc"
	"forum/app/post/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	WindowSize = 7 * 1 * 1 * 1
)

type DeletePostSchedulerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostSchedulerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostSchedulerLogic {
	return &DeletePostSchedulerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostSchedulerLogic) DeletePostScheduler(in *pb.DeletePostSchedulerRequest) (*pb.DeletePostSchedulerResponse, error) {
	// 计算时间窗口
	stop := time.Now().Unix() - WindowSize
	key := globalkey.GetRedisKey(globalkey.PostScoreKey)

	// 1. 首先查询要处理的数据
	postList, err := l.svcCtx.RedisClient.ZrangebyscoreWithScores(key, 0, stop)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Redis ZRangeByScore error %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("Redis ZRangeByScore error"), "Redis ZRangeByScore error %v", err)
	}

	// 2. 如果有数据需要处理,再执行删除操作
	if len(postList) > 0 {
		removedCount, err := l.svcCtx.RedisClient.Zremrangebyscore(key, 0, stop)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("Redis ZRemRangeByScore error %v", err)
			return nil, errors.Wrapf(xerr.NewErrMsg("Redis ZRemRangeByScore error"), "Redis ZRemRangeByScore error %v", err)
		}

		logx.WithContext(l.ctx).Infof("Successfully processed %d posts, removed %d entries", len(postList), removedCount)

		// 3. 处理获取到的数据
		for _, post := range postList {
			pid, _ := strconv.ParseInt(post.Key, 10, 64)
			msg := mq.DeletePostMessage{
				PostId: pid,
			}
			body, _ := json.Marshal(msg)
			// 发送消息到消息队列
			err = l.svcCtx.RabbitMqClient.Send("delete_post", "up", body)
			if err != nil {
				logx.WithContext(l.ctx).Errorf("Send delete_post message error %v", err)
				return nil, errors.Wrapf(xerr.NewErrMsg("Send delete_post message error"), "Send delete_post message error %v", err)
			}
			err = l.svcCtx.RabbitMqClient.Send("delete_post", "down", body)
			if err != nil {
				logx.WithContext(l.ctx).Errorf("Send delete_post message error %v", err)
				return nil, errors.Wrapf(xerr.NewErrMsg("Send delete_post message error"), "Send delete_post message error %v", err)
			}
		}
	}

	return &pb.DeletePostSchedulerResponse{
		Success: true,
	}, nil
}
