package comment

import (
	"context"

	"forum/app/comment/api/internal/svc"
	"forum/app/comment/api/internal/types"
	"forum/app/comment/rpc/commentservice"
	"forum/app/post/rpc/postservice"
	"forum/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建评论
func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentReq) (resp *types.CreateCommentResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.PostRpc.GetPostDetail(l.ctx, &postservice.GetPostDetailRequest{
		PostId: req.PostId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("获取帖子详情失败: %v", err)
		return nil, err
	}

	// 创建评论
	commentResp, err := l.svcCtx.CommentRpc.CreateComment(l.ctx, &commentservice.CreateCommentRequest{
		PostId:   req.PostId,
		AuthorId: userId,
		Content:  req.Content,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("创建评论失败: %v", err)
		return nil, err
	}
	logx.WithContext(l.ctx).Infof("创建评论成功: %v", commentResp)
	resp = &types.CreateCommentResp{}
	copier.Copy(resp, commentResp)

	return
}
