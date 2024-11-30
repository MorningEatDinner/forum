package logic

import (
	"context"
	"fmt"
	"forum/app/vote/rpc/internal/svc"
	"forum/app/vote/rpc/pb"
	"forum/common/globalkey"
	"forum/common/xerr"
	"forum/tmp/app/post/rpc/postservice"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
)

const (
	VoteTypeOppose    = -1               // 反对
	VoteTypeCancel    = 0                // 取消
	VoteTypeAgree     = 1                // 赞成
	VoteKeyExpiration = 7 * 24 * 60 * 60 // 设置7天的过期时间
	ScoreIncrHot      = 1
	ScoreDescHot      = -1
)

type VotePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVotePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VotePostLogic {
	return &VotePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VotePostLogic) VotePost(in *pb.VotePostRequest) (*pb.VotePostResponse, error) {
	// 1. 查询是否已投票-- Redis中查询是否已经存在投票记录了
	key := globalkey.GetRedisKey(fmt.Sprintf(globalkey.VoteRecordKey, in.PostId))
	votedType, err := l.svcCtx.RedisClient.Zscore(key, strconv.FormatInt(in.UserId, 10))
	if err == redis.Nil {
		// 3. 新增投票
		if in.VoteType != VoteTypeCancel {
			// 新增投票， 直接在Redis中记录
			_, err = l.svcCtx.RedisClient.Zadd(key, int64(in.VoteType), strconv.FormatInt(in.UserId, 10))
			if err != nil {
				return nil, errors.Wrapf(err, "zadd vote record failed, postId: %d, userId: %d", in.PostId, in.UserId)
			}
			// 续期
			if err = l.renewVoteKeyExpiration(key); err != nil {
				return nil, errors.Wrapf(err, "renew vote key expiration failed, key: %s", key)
			}

			// 无论是赞成还是反对， 都是增加热点分
			// TODO: 调用rpc方法修改post 服务中的分数
			_, err = l.svcCtx.PostRpc.UpdatePostScore(l.ctx, &postservice.UpdatePostScoreRequest{
				PostId: in.PostId,
				Score:  ScoreIncrHot,
				Up:     in.VoteType == VoteTypeAgree,
				Down:   in.VoteType == VoteTypeOppose,
			})
			if err != nil {
				return nil, errors.Wrapf(err, "update post score failed, postId: %d, userId: %d", in.PostId, in.UserId)
			}
		}
		return &pb.VotePostResponse{
			Success: true,
		}, nil
	}
	// 真的错误了
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "zscore vote record failed, postId: %d, userId: %d", in.PostId, in.UserId)
	}

	// 2. 已投票且方向相同则返回
	if votedType != 0 && votedType == int64(in.VoteType) {
		return nil, errors.New("already voted with same direction")
	}

	// 2. 处理已投票的情况
	if votedType != 0 {
		// 如果取消投票， 就是现在投票操作为0， 那么就删除， 并且减少post服务中的分数
		if in.VoteType == VoteTypeCancel {
			// 那么删除投票数据
			_, err = l.svcCtx.RedisClient.Zrem(key, strconv.FormatInt(in.UserId, 10))
			if err != nil {
				return nil, errors.Wrapf(err, "zrem vote record failed, postId: %d, userId: %d", in.PostId, in.UserId)
			}
			// TODO: 调用rpc方法减少post 服务中的分数， 这里可以改用消息队列来实现
			// 因为是撤销， 所以要减少分数
			_, err = l.svcCtx.PostRpc.UpdatePostScore(l.ctx, &postservice.UpdatePostScoreRequest{
				PostId: in.PostId,
				Score:  ScoreDescHot,
				Up:     votedType == VoteTypeAgree,
				Down:   votedType == VoteTypeOppose,
			})
			if err != nil {
				return nil, errors.Wrapf(err, "update post score failed, postId: %d, userId: %d", in.PostId, in.UserId)
			}
		}

		// 现在是要修改投票的方向
		changeScore := int64(in.VoteType) - int64(votedType)
		_, err = l.svcCtx.RedisClient.Zincrby(key, changeScore, strconv.FormatInt(in.UserId, 10))
		if err != nil {
			return nil, errors.Wrapf(err, "zincrby vote record failed, postId: %d, userId: %d", in.PostId, in.UserId)
		}

		// 续期
		if err = l.renewVoteKeyExpiration(key); err != nil {
			return nil, errors.Wrapf(err, "renew vote key expiration failed, key: %s", key)
		}

		// 转换投票方向不会对热度造成影响， 所以不需要调用rpc方法
		return &pb.VotePostResponse{Success: true}, nil
	}

	return nil, errors.Wrapf(xerr.NewErrMsg("server error"), "vote post failed, postId: %d, userId: %d", in.PostId, in.UserId)
}

// 添加续期方法
func (l *VotePostLogic) renewVoteKeyExpiration(key string) error {
	return l.svcCtx.RedisClient.Expire(key, VoteKeyExpiration)
}
