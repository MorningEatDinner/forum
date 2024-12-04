// 包声明
package logic

// 导入所需的包
import (
	"context"
	"encoding/json"
	"forum/app/vote/model"
	"forum/app/vote/rpc/internal/svc"
	"forum/app/vote/rpc/pb"
	"forum/common/mq"
	"forum/common/xerr"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// 定义常量
const (
	VoteTypeOppose = -1 // 反对
	VoteTypeCancel = 0  // 取消
	VoteTypeAgree  = 1  // 赞成
	ScoreIncrHot   = 1  // 热度分数增加值
	ScoreDescHot   = -1 // 热度分数减少值
	ScoreNotChange = 0  // 不改变分数
)

// VotePostLogic 结构体定义
type VotePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewVotePostLogic 创建 VotePostLogic 实例的函数
func NewVotePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VotePostLogic {
	return &VotePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// VotePost 方法实现投票逻辑
func (l *VotePostLogic) VotePost(in *pb.VotePostRequest) (*pb.VotePostResponse, error) {
	// 将请求中的投票类型转换为 int64
	voteTypeReq := int64(in.VoteType)

	// 开启事务处理投票逻辑
	err := l.svcCtx.VoteRecordModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 查询用户是否已对该帖子投票
		votedInfo, err := l.svcCtx.VoteRecordModel.FindOneByPostIdUserId(l.ctx, uint64(in.PostId), uint64(in.UserId))
		if err == model.ErrNotFound {
			// 如果未找到投票记录且不是取消投票，则处理新投票
			if voteTypeReq != VoteTypeCancel {
				err = l.handleNewVote(ctx, session, in)
				if err != nil {
					return err
				}
			}
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "查询投票记录失败")
		}

		// 如果投票类型相同，则返回错误
		if votedInfo.VoteType != 0 && votedInfo.VoteType == int64(in.VoteType) {
			return l.handleSameVote(in)
		}

		// 处理取消投票或更改投票
		if votedInfo.VoteType != VoteTypeCancel {
			if in.VoteType == VoteTypeCancel {
				return l.handleCancelVote(ctx, votedInfo, in)
			}
			return l.handleChangeVote(ctx, votedInfo, in)
		}

		return nil
	})

	// 处理投票错误
	if err != nil {
		logx.WithContext(l.ctx).Errorf("投票失败: %v", err)
		return nil, errors.Wrap(xerr.NewErrMsg("投票失败"), err.Error())
	}

	// 返回投票成功响应
	return &pb.VotePostResponse{Success: true}, nil
}

// handleSameVote 处理重复投票
func (l *VotePostLogic) handleSameVote(in *pb.VotePostRequest) error {
	logx.WithContext(l.ctx).Errorf("投票类型相同, postId: %d, userId: %d", in.PostId, in.UserId)
	return errors.Wrapf(xerr.NewErrMsg("投票失败"), "投票类型相同, postId: %d, userId: %d", in.PostId, in.UserId)
}

// handleNewVote 处理新投票
func (l *VotePostLogic) handleNewVote(ctx context.Context, session sqlx.Session, in *pb.VotePostRequest) error {
	// 插入新的投票记录
	_, err := l.svcCtx.VoteRecordModel.Insert(ctx, session, &model.VoteRecord{
		PostId:      uint64(in.PostId),
		UserId:      uint64(in.UserId),
		VoteType:    int64(in.VoteType),
		CreateTime:  time.Now(),
		UpdatedTime: time.Now(),
	})
	if err != nil {
		return errors.Wrap(xerr.NewErrMsg("投票失败"), "插入投票记录失败")
	}

	// 更新帖子分数
	err = l.updatePostVoteScore(ctx, in.PostId, ScoreIncrHot, int64(in.VoteType))
	if err != nil {
		return err
	}

	return nil
}

// handleCancelVote 处理取消投票
func (l *VotePostLogic) handleCancelVote(ctx context.Context, votedInfo *model.VoteRecord, in *pb.VotePostRequest) error {
	// 删除投票记录
	err := l.svcCtx.VoteRecordModel.Delete(ctx, votedInfo.VoteId)
	if err != nil {
		return errors.Wrap(xerr.NewErrMsg("投票失败"), "删除投票记录失败")
	}

	// 更新帖子分数
	err = l.updatePostVoteScore(ctx, in.PostId, ScoreDescHot, votedInfo.VoteType)
	if err != nil {
		return err
	}

	return nil
}

// handleChangeVote 处理更改投票
func (l *VotePostLogic) handleChangeVote(ctx context.Context, votedInfo *model.VoteRecord, in *pb.VotePostRequest) error {
	// 计算分数变化
	changeScore := int64(in.VoteType) - int64(votedInfo.VoteType)
	votedInfo.VoteType += changeScore

	// 更新投票记录
	err := l.svcCtx.VoteRecordModel.Update(ctx, votedInfo)
	if err != nil {
		return errors.Wrap(xerr.NewErrMsg("投票失败"), "更新投票记录失败")
	}

	// 更新帖子分数
	err = l.updatePostVoteScore(ctx, in.PostId, ScoreNotChange, votedInfo.VoteType)
	if err != nil {
		return err
	}

	return nil
}

// updatePostVoteScore 更新帖子的投票分数
func (l *VotePostLogic) updatePostVoteScore(ctx context.Context, postId int64, score int64, voteType int64) error {
	// 根据投票类型确定是增加还是减少
	isUp := voteType == VoteTypeAgree
	isDown := voteType == VoteTypeOppose

	// 调用RPC更新帖子分数
	payload := mq.UpdatePostScore{
		PostId: postId,
		Score:  score,
		Up:     isUp,
		Down:   isDown,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		logx.WithContext(ctx).Errorf("marshal message failed, err: %v", err.Error())
		return errors.Wrapf(xerr.NewErrMsg("marshal message failed"), "marshal message failed, err: %v", err.Error())
	}
	// _, err := l.svcCtx.PostRpc.UpdatePostScore(ctx, req)
	// if err != nil {
	// 	return errors.Wrapf(err, "更新帖子分数失败, postId: %d", postId)
	// }

	return l.svcCtx.RabbitMqClient.Send("vote", "update_post_score", body)
}
