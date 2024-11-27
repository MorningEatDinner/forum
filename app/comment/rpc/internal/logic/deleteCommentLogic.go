package logic

import (
	"context"

	"forum/app/comment/rpc/internal/svc"
	"forum/app/comment/rpc/pb"
	"forum/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除评论
func (l *DeleteCommentLogic) DeleteComment(in *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	// 1. 查看评论是否存在
	commentInfo, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.CommentId)
	if err != nil {
		// 更明确的错误日志，指明是查询评论失败
		logx.WithContext(l.ctx).Errorf("rpc delete comment error: failed to find comment with id %d: %v", in.CommentId, err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.COMMENT_NOT_FOUND), "failed to find comment with id %d: %v", in.CommentId, err.Error())
	}

	// 2. 查看评论是否是自己的
	if commentInfo.AuthorId != in.AuthorId {
		logx.WithContext(l.ctx).Errorf("rpc delete comment error: permission denied for user %d to delete comment %d", in.AuthorId, in.CommentId)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PERMISSION_DENIED), "user %d does not have permission to delete comment %d", in.AuthorId, in.CommentId)
	}

	// 3. 删除评论
	err = l.svcCtx.CommentModel.Delete(l.ctx, in.CommentId)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("rpc delete comment error: failed to delete comment with id %d: %v", in.CommentId, err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to delete comment with id %d: %v", in.CommentId, err.Error())
	}

	// 删除成功，返回响应
	return &pb.DeleteCommentResponse{}, nil
}
