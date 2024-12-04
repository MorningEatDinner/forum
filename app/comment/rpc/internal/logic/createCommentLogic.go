package logic

import (
	"context"
	"time"

	"forum/app/comment/model"
	"forum/app/comment/rpc/internal/svc"
	"forum/app/comment/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建评论
func (l *CreateCommentLogic) CreateComment(in *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	insertComment := &model.Comments{
		PostId:      in.PostId,
		Content:     in.Content,
		AuthorId:    in.AuthorId,
		CreateTime:  time.Now(),
		UpdatedTime: time.Now(),
	}

	insertRes, err := l.svcCtx.CommentModel.Insert(l.ctx, insertComment)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("创建评论失败: %v", err)
		return nil, err
	}
	// TODO：创建帖子评论数表， 后面就需要增加了
	id, _ := insertRes.LastInsertId()
	insertComment.CommentId = id
	logx.WithContext(l.ctx).Infof("创建评论成功: %v", insertComment)
	pbComment := &pb.Comment{}
	copier.Copy(pbComment, insertComment)
	return &pb.CreateCommentResponse{
		Comment: pbComment,
	}, nil
}
