package logic

import (
	"context"
	"forum/app/vote/rpc/voteservice"
	"forum/common/globalkey"
	"forum/common/xerr"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"

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
	logx.WithContext(l.ctx).Infof("UpdatePostScore 收到请求参数 - PostId: %d, Score: %d, Up: %v, Down: %v",
		in.PostId, in.Score, in.Up, in.Down)

	_, err := l.svcCtx.RedisClient.Zscore(globalkey.GetRedisKey(globalkey.PostScoreKey), strconv.FormatInt(in.PostId, 10))
	if err == redis.Nil {
		voteRecord, err := l.svcCtx.VoteRpc.GetPostVoteCounts(l.ctx, &voteservice.GetPostVoteCountsRequest{PostId: in.PostId})
		if err != nil {
			logx.WithContext(l.ctx).Errorf("获取 Vote rpc 数据失败, err: %v", err)
			return nil, errors.Wrapf(xerr.NewErrMsg("获取 Vote rpc 数据失败"), "err: %v", err)
		}

		pipe, err := l.svcCtx.RedisClient.TxPipeline()
		if err != nil {
			logx.WithContext(l.ctx).Errorf("初始化 Redis 事务失败, err: %v", err)
			return nil, errors.Wrapf(xerr.NewErrMsg("初始化 Redis 事务失败"), "err: %v", err)
		}
		defer pipe.Discard()

		// 可以考虑使用一个较早的初始时间或基于其他因素的初始分数
		initialScore := time.Now().Add(-24 * time.Hour).Unix()

		pipe.ZAdd(l.ctx, globalkey.GetRedisKey(globalkey.PostScoreKey), redis.Z{
			Member: strconv.FormatInt(in.PostId, 10),
			Score:  float64(initialScore),
		})
		pipe.ZAdd(l.ctx, globalkey.GetRedisKey(globalkey.PostUpCountKey), redis.Z{
			Member: strconv.FormatInt(in.PostId, 10),
			Score:  float64(voteRecord.Upvotes),
		})
		pipe.ZAdd(l.ctx, globalkey.GetRedisKey(globalkey.PostDownCountKey), redis.Z{
			Member: strconv.FormatInt(in.PostId, 10),
			Score:  float64(voteRecord.Downvotes),
		})

		_, err = pipe.Exec(l.ctx)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("执行 Redis 事务失败, err: %v", err)
			return nil, errors.Wrapf(xerr.NewErrMsg("执行 Redis 事务失败"), "err: %v", err)
		}
		return &pb.UpdatePostScoreResponse{Success: true}, nil
	}
	if err != nil {
		logx.WithContext(l.ctx).Errorf("获取 Redis 数据失败, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("获取 Redis 数据失败"), "err: %v", err)
	}

	pipe, err := l.svcCtx.RedisClient.TxPipeline()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("初始化 Redis 事务失败, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("初始化 Redis 事务失败"), "err: %v", err)
	}
	defer pipe.Discard()

	if in.Score != 0 {
		postKey := strconv.FormatInt(in.PostId, 10)
		pipe.ZIncrBy(l.ctx, globalkey.GetRedisKey(globalkey.PostScoreKey), float64(in.Score*HotBase), postKey)

		if in.Up {
			pipe.ZIncrBy(l.ctx, globalkey.GetRedisKey(globalkey.PostUpCountKey), float64(in.Score), postKey)
		}

		if in.Down {
			pipe.ZIncrBy(l.ctx, globalkey.GetRedisKey(globalkey.PostDownCountKey), float64(in.Score), postKey)
		}
	} else {
		logx.WithContext(l.ctx).Errorf("分数为 0, 不执行更新分数操作， 但是另外两个操作需要更新")
		if in.Up {
			pipe.ZIncrBy(l.ctx, globalkey.GetRedisKey(globalkey.PostUpCountKey), float64(1), strconv.FormatInt(in.PostId, 10))
		} else {
			pipe.ZIncrBy(l.ctx, globalkey.GetRedisKey(globalkey.PostUpCountKey), float64(-1), strconv.FormatInt(in.PostId, 10))
		}
		if in.Down {
			pipe.ZIncrBy(l.ctx, globalkey.GetRedisKey(globalkey.PostDownCountKey), float64(1), strconv.FormatInt(in.PostId, 10))
		} else {
			pipe.ZIncrBy(l.ctx, globalkey.GetRedisKey(globalkey.PostDownCountKey), float64(-1), strconv.FormatInt(in.PostId, 10))
		}
	}

	_, err = pipe.Exec(l.ctx)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("执行 Redis 事务失败, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("执行 Redis 事务失败"), "err: %v", err)
	}

	return &pb.UpdatePostScoreResponse{Success: true}, nil
}
