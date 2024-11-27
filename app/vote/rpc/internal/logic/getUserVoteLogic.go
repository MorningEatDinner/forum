package logic

import (
	"context"

	"forum/app/vote/model"
	"forum/app/vote/rpc/internal/svc"
	"forum/app/vote/rpc/pb"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
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
	voteInfo, err := l.svcCtx.VoteRecordModel.FindOneByPostIdUserId(l.ctx, uint64(in.PostId), uint64(in.UserId))
	if err != nil && err != model.ErrNotFound {
		logx.WithContext(l.ctx).Errorf("get user vote failed, err: %v, userId: %d, postId: %d", err, in.UserId, in.PostId)
		return nil, errors.Wrapf(xerr.NewErrMsg("get user vote failed"), "get user vote failed, err: %v", err)
	}

	voteResp := &pb.VoteRecord{}
	copier.Copy(voteResp, voteInfo)

	return &pb.GetUserVoteResponse{
		VoteRecord: voteResp,
	}, nil
}
