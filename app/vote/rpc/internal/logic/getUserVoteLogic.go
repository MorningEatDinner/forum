package logic

import (
	"context"
	"fmt"
	"forum/common/globalkey"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"

	"forum/app/vote/rpc/internal/svc"
	"forum/app/vote/rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserVoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserVoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserVoteLogic {
	return &GetUserVoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserVoteLogic) GetUserVote(in *pb.GetUserVoteRequest) (*pb.GetUserVoteResponse, error) {
	// 获取用户的投票信息
	key := fmt.Sprintf(globalkey.VoteRecordKey, in.PostId)
	voteInfo, err := l.svcCtx.RedisClient.Zscore(key, strconv.FormatInt(in.UserId, 10))
	if err == redis.Nil {
		return &pb.GetUserVoteResponse{VoteRecord: 0}, nil
	}
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get user vote failed, postId: %d, userId: %d", in.PostId, in.UserId)
		return nil, errors.Wrapf(err, "get user vote failed, postId: %d, userId: %d", in.PostId, in.UserId)
	}

	return &pb.GetUserVoteResponse{
		VoteRecord: voteInfo,
	}, nil
}
