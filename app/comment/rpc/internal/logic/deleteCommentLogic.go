package logic

import (
	"context"

	"forum/app/comment/rpc/internal/svc"
	"forum/app/comment/rpc/pb"

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
	// todo: add your logic here and delete this line

	return &pb.DeleteCommentResponse{}, nil
}
