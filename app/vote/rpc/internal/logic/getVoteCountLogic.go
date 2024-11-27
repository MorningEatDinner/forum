package logic

import (
	"context"

	"forum/app/vote/rpc/internal/svc"
	"forum/app/vote/rpc/pb"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetVoteCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVoteCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVoteCountLogic {
	return &GetVoteCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVoteCountLogic) GetVoteCount(in *pb.GetVoteCountRequest) (*pb.GetVoteCountResponse, error) {
	voteCount, err := l.svcCtx.VoteCountModel.FindOneByPostId(l.ctx, uint64(in.PostId))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get vote count failed, postId: %d, err: %v", in.PostId, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get vote count failed, postId: %d", in.PostId)
	}

	voteCountResp := &pb.VoteCount{}
	copier.Copy(voteCountResp, voteCount)

	return &pb.GetVoteCountResponse{
		VoteCount: voteCountResp,
	}, nil
}
