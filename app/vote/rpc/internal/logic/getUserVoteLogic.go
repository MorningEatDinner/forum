package logic

import (
	"context"
	"forum/app/user/model"
	"forum/app/vote/rpc/internal/svc"
	"forum/app/vote/rpc/pb"
	"forum/common/xerr"
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
	votedInfo, err := l.svcCtx.VoteRecordModel.FindOneByPostIdUserId(l.ctx, uint64(in.PostId), uint64(in.UserId))
	if err == model.ErrNotFound {
		return &pb.GetUserVoteResponse{
			VoteRecord: 0,
		}, nil
	}
	if err != nil {
		logx.WithContext(l.ctx).Errorf("rpc GetUserVote FindOneByPostIdUserId err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("投票记录不存在"), "rpc GetUserVote FindOneByPostIdUserId err: %v", err)
	}

	return &pb.GetUserVoteResponse{
		VoteRecord: votedInfo.VoteType,
	}, nil
}
