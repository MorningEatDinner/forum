package logic

import (
	"context"
	"time"

	"forum/app/vote/model"
	"forum/app/vote/rpc/internal/svc"
	"forum/app/vote/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	VoteTypeOppose = -1 // 反对
	VoteTypeCancel = 0  // 取消
	VoteTypeAgree  = 1  // 赞成
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
	// 1. 查询是否已投票
	voteInfo, err := l.svcCtx.VoteRecordModel.FindOneByPostIdUserId(l.ctx, uint64(in.PostId), uint64(in.UserId))
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(err, "get user vote failed, userId: %d, postId: %d", in.UserId, in.PostId)
	}

	// 2. 已投票且方向相同则返回
	if voteInfo != nil && voteInfo.VoteType == int64(in.VoteType) {
		return nil, errors.New("already voted with same direction")
	}

	// 2. 处理已投票的情况
	if voteInfo != nil {
		voteCountInfo, err := l.svcCtx.VoteCountModel.FindOneByPostId(l.ctx, uint64(in.PostId))
		if err != nil {
			return nil, errors.Wrapf(err, "get vote count failed, postId: %d", in.PostId)
		}
		voteCountId := voteCountInfo.VoteCountId
		// 取消投票
		if in.VoteType == VoteTypeCancel {
			err := l.svcCtx.VoteRecordModel.Delete(l.ctx, voteInfo.VoteId)
			if err != nil {
				return nil, errors.Wrapf(err, "delete vote failed, voteId: %d", voteInfo.VoteId)
			}
			if voteInfo.VoteType == VoteTypeAgree {
				// 如果是赞成票
				err = l.svcCtx.VoteCountModel.DecrementAgreeCount(l.ctx, voteCountId)
				if err != nil {
					return nil, errors.Wrapf(err, "decrement agree count failed, voteId: %d", voteCountId)
				}
				// TODO: 需要修改帖子的分数
			} else {
				// 如果是反对票
				err = l.svcCtx.VoteCountModel.DecrementOpposeCount(l.ctx, voteCountId)
				if err != nil {
					return nil, errors.Wrapf(err, "decrement oppose count failed, voteId: %d", voteCountId)
				}
				// TODO: 需要修改帖子的分数
			}
			return &pb.VotePostResponse{Success: true}, nil
		}

		voteInfo.VoteType = int64(in.VoteType)
		err = l.svcCtx.VoteRecordModel.Update(l.ctx, voteInfo)
		if err != nil {
			return nil, errors.Wrapf(err, "update vote failed, voteId: %d", voteInfo.VoteId)
		}
		if in.VoteType == VoteTypeAgree {
			// 那么原来是反对票，现在改为赞成票
			err = l.svcCtx.VoteCountModel.DecrementOpposeCount(l.ctx, voteCountId)
			if err != nil {
				return nil, errors.Wrapf(err, "decrement oppose count failed, voteId: %d", voteCountId)
			}
			err = l.svcCtx.VoteCountModel.IncrementAgreeCount(l.ctx, voteCountId)
			if err != nil {
				return nil, errors.Wrapf(err, "increment agree count failed, voteId: %d", voteCountId)
			}
		} else {
			// 那么原来是赞成票，现在改为反对票
			err = l.svcCtx.VoteCountModel.DecrementAgreeCount(l.ctx, voteCountId)
			if err != nil {
				return nil, errors.Wrapf(err, "decrement agree count failed, voteId: %d", voteCountId)
			}
			err = l.svcCtx.VoteCountModel.IncrementOpposeCount(l.ctx, voteCountId)
			if err != nil {
				return nil, errors.Wrapf(err, "increment oppose count failed, voteId: %d", voteCountId)
			}
		}
		// TODO: 调用rpc请求， 修改帖子的分数

		return &pb.VotePostResponse{Success: true}, nil
	}

	// 3. 新增投票
	if in.VoteType != VoteTypeCancel {
		vote := &model.VoteRecord{
			PostId:      uint64(in.PostId),
			UserId:      uint64(in.UserId),
			VoteType:    int64(in.VoteType),
			CreateTime:  time.Now(),
			UpdatedTime: time.Now(),
		}
		_, err := l.svcCtx.VoteRecordModel.Insert(l.ctx, vote)
		if err != nil {
			return nil, errors.Wrapf(err, "insert vote failed, postId: %d, userId: %d", in.PostId, in.UserId)
		}

		voteCountInfo, err := l.svcCtx.VoteCountModel.FindOneByPostId(l.ctx, uint64(in.PostId))
		if err == model.ErrNotFound {
			// 不存在则创建
			voteCount := &model.VoteCount{
				PostId:      uint64(in.PostId),
				AgreeCount:  0,
				OpposeCount: 0,
				CreateTime:  time.Now(),
				UpdatedTime: time.Now(),
			}
			res, err := l.svcCtx.VoteCountModel.Insert(l.ctx, voteCount)
			if err != nil {
				return nil, errors.Wrapf(err, "insert vote count failed, postId: %d", in.PostId)
			}
			voteCountInfo = voteCount
			voteCountInfo.VoteCountId, _ = res.LastInsertId()
		} else if err != nil {
			return nil, errors.Wrapf(err, "get vote count failed, postId: %d", in.PostId)
		}

		if in.VoteType == VoteTypeAgree {
			err = l.svcCtx.VoteCountModel.IncrementAgreeCount(l.ctx, voteCountInfo.VoteCountId)
			if err != nil {
				return nil, errors.Wrapf(err, "increment agree count failed, voteId: %d", voteCountInfo.VoteCountId)
			}
		} else {
			err = l.svcCtx.VoteCountModel.IncrementOpposeCount(l.ctx, voteCountInfo.VoteCountId)
			if err != nil {
				return nil, errors.Wrapf(err, "increment oppose count failed, voteId: %d", voteCountInfo.VoteCountId)
			}
		}
		// TODO: 调用rpc方法修改帖子分数
	}

	return &pb.VotePostResponse{
		Success: true,
	}, nil
}
