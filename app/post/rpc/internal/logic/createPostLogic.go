package logic

import (
	"context"
	"time"

	"forum/app/post/model"
	"forum/app/post/rpc/internal/svc"
	"forum/app/post/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePostLogic) CreatePost(in *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	post := &model.Posts{
		CommunityId: in.CommunityId,
		AuthorId:    in.AuthorId,
		Title:       in.Title,
		Content:     in.Content,
		CreateTime:  time.Now(),
		UpdatedTime: time.Now(),
	}
	insertRes, err := l.svcCtx.PostModel.Insert(l.ctx, post)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create post failed, err: %v", err)
		return nil, err
	}
	id, _ := insertRes.LastInsertId()

	return &pb.CreatePostResponse{
		PostId: id,
	}, nil
}
