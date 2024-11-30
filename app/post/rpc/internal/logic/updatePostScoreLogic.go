package logic

import (
	"context"
	"forum/common/globalkey"
	"forum/common/xerr"
	"github.com/pkg/errors"
	"strconv"

	"forum/app/post/rpc/internal/svc"
	"forum/app/post/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	HotBase = 432
)

type UpdatePostScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostScoreLogic {
	return &UpdatePostScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePostScoreLogic) UpdatePostScore(in *pb.UpdatePostScoreRequest) (*pb.UpdatePostScoreResponse, error) {
	// 打印完整的入参信息
	logx.WithContext(l.ctx).Infof("UpdatePostScore 收到请求参数 - PostId: %d, Score: %d, Up: %v, Down: %v",
		in.PostId, in.Score, in.Up, in.Down)
	// 开启Redis事务
	pipe, err := l.svcCtx.RedisClient.TxPipeline()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("初始化 Redis 事务失败, err:%v", err)
		return nil, errors.Wrap(xerr.NewErrMsg("初始化 Redis 事务失败"), err.Error())
	}
	defer pipe.Discard()

	postKey := strconv.FormatInt(in.PostId, 10)
	// 更新主分数
	pipe.ZIncrBy(l.ctx, globalkey.GetRedisKey(globalkey.PostScoreKey), float64(in.Score*HotBase), postKey)

	// 更新投票分数
	if in.Up {
		pipe.ZIncrBy(l.ctx, globalkey.GetRedisKey(globalkey.PostUpCountKey), float64(in.Score), postKey)
	}

	if in.Down {
		pipe.ZIncrBy(l.ctx, globalkey.GetRedisKey(globalkey.PostDownCountKey), float64(in.Score), postKey)
	}

	// 执行事务
	_, err = pipe.Exec(l.ctx)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("执行 Redis 事务失败, err:%v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("执行 Redis 事务失败"), "err:%v", err)
	}

	return &pb.UpdatePostScoreResponse{Success: true}, nil
}
