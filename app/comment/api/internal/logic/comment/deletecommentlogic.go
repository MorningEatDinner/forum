package comment

import (
	"context"

	"forum/app/comment/api/internal/svc"
	"forum/app/comment/api/internal/types"
	"forum/app/comment/rpc/commentservice"
	"forum/common/ctxdata"
	"forum/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除评论
func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentReq) (resp *types.DeleteCommentResp, err error) {
	// 删除评论， 只有评论的用户自己才可以删除
	userId := ctxdata.GetUidFromCtx(l.ctx)

	// 删除评论
	_, err = l.svcCtx.CommentRpc.DeleteComment(l.ctx, &commentservice.DeleteCommentRequest{
		CommentId: req.CommentId,
		AuthorId:  userId,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("rpc delete comment error: %v", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "rpc delete comment error: %v", err.Error())
	}

	resp = &types.DeleteCommentResp{}

	return
}
