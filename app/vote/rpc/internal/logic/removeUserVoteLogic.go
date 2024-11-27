package logic

import (
	"context"

	"forum/app/vote/rpc/internal/svc"
	"forum/app/vote/rpc/pb"
	"forum/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveUserVoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveUserVoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserVoteLogic {
	return &RemoveUserVoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveUserVoteLogic) RemoveUserVote(in *pb.RemoveUserVoteRequest) (*pb.RemoveUserVoteResponse, error) {
	err := l.svcCtx.VoteCountModel.Delete(l.ctx, in.VoteId)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to delete vote: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to delete vote"), "failed to delete vote")
	}

	return &pb.RemoveUserVoteResponse{}, nil
}
