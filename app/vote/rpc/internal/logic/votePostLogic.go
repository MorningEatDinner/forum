package logic

import (
	"context"
	"forum/app/post/rpc/postservice"
	"forum/app/vote/model"
	"forum/app/vote/rpc/internal/svc"
	"forum/app/vote/rpc/pb"
	"forum/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

const (
	VoteTypeOppose = -1 // 反对
	VoteTypeCancel = 0  // 取消
	VoteTypeAgree  = 1  // 赞成
	ScoreIncrHot   = 1
	ScoreDescHot   = -1
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
	// 1. 查询是否存在投票记录
	voteTypeReq := int64(in.VoteType)
	votedInfo, err := l.svcCtx.VoteRecordModel.FindOneByPostIdUserId(l.ctx, uint64(in.PostId), uint64(in.UserId))
	if err == model.ErrNotFound {
		// 如果没有投票记录
		if voteTypeReq != VoteTypeCancel {
			// 如果投票类型与之前的类型不一致，直接更新
			_, err = l.svcCtx.VoteRecordModel.Insert(l.ctx, &model.VoteRecord{
				PostId:      uint64(in.PostId),
				UserId:      uint64(in.UserId),
				VoteType:    int64(in.VoteType),
				CreateTime:  time.Now(),
				UpdatedTime: time.Now(),
			})
			if err != nil {
				return nil, errors.Wrapf(xerr.NewErrMsg("投票失败"), "insert vote record failed, err: %v", err)
			}

			// 调用rpc方法新增post的分数
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
		// 如果投 原来没有投票记录， 现在投0， 那么直接返回
		return &pb.VotePostResponse{
			Success: true,
		}, nil
	}
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("投票失败"), "find vote record failed, err: %v", err)
	}

	// 2. 已投票且方向相同则返回
	if votedInfo.VoteType != 0 && votedInfo.VoteType == int64(in.VoteType) {
		return nil, errors.New("already voted with same direction")
	}

	if votedInfo.VoteType != VoteTypeCancel {
		// 如果取消投票， 就是现在投票操作为0， 那么就删除， 并且减少post服务中的分数
		if in.VoteType == VoteTypeCancel {
			// 那么删除投票数据
			err = l.svcCtx.VoteRecordModel.Delete(l.ctx, votedInfo.VoteId)
			if err != nil {
				return nil, errors.Wrapf(xerr.NewErrMsg("投票失败"), "delete vote record failed, err: %v", err)
			}
			// TODO: 调用rpc方法减少post 服务中的分数， 这里可以改用消息队列来实现
			// 因为是撤销， 所以要减少分数
			_, err = l.svcCtx.PostRpc.UpdatePostScore(l.ctx, &postservice.UpdatePostScoreRequest{
				PostId: in.PostId,
				Score:  ScoreDescHot,
				Up:     votedInfo.VoteType == VoteTypeAgree,
				Down:   votedInfo.VoteType == VoteTypeOppose,
			})
			if err != nil {
				return nil, errors.Wrapf(err, "update post score failed, postId: %d, userId: %d", in.PostId, in.UserId)
			}
			return &pb.VotePostResponse{Success: true}, nil
		}

		// 现在是要修改投票的方向
		changeScore := int64(in.VoteType) - int64(votedInfo.VoteType)
		votedInfo.VoteType += changeScore
		err = l.svcCtx.VoteRecordModel.Update(l.ctx, votedInfo)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("投票失败"), "update vote record failed, err: %v", err)
		}
		// 转换投票方向不会对热度造成影响， 所以不需要调用rpc方法
		return &pb.VotePostResponse{Success: true}, nil
	}

	return nil, errors.Wrapf(xerr.NewErrMsg("server error"), "vote post failed, postId: %d, userId: %d", in.PostId, in.UserId)
}
