package logic

import (
	"context"
	"forum/common/xerr"
	"github.com/pkg/errors"

	"forum/app/vote/rpc/internal/svc"
	"forum/app/vote/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostVoteCountsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostVoteCountsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostVoteCountsLogic {
	return &GetPostVoteCountsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostVoteCountsLogic) GetPostVoteCounts(in *pb.GetPostVoteCountsRequest) (*pb.GetPostVoteCountsResponse, error) {
	// 计算帖子的投票数和踩数
	upCount, err := l.svcCtx.VoteRecordModel.CountUpvotes(l.ctx, uint64(in.PostId))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to count upvotes, postId: %d, err: %v", in.PostId, err)
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to count upvotes"), "failed to count upvotes, postId: %d, err: %v", in.PostId, err)
	}
	downCount, err := l.svcCtx.VoteRecordModel.CountDownvotes(l.ctx, uint64(in.PostId))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to count downvotes, postId: %d, err: %v", in.PostId, err)
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to count downvotes"), "failed to count downvotes, postId: %d, err: %v", in.PostId, err)
	}

	return &pb.GetPostVoteCountsResponse{
		Upvotes:   upCount,
		Downvotes: downCount,
	}, nil
}
