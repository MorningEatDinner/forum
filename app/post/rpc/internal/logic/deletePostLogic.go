package logic

import (
	"context"

	"forum/app/post/rpc/internal/svc"
	"forum/app/post/rpc/pb"
	"forum/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostLogic) DeletePost(in *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	// 1. 获取帖子的信息
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, in.PostId)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get post failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get post failed")
	}

	// 2. 如果帖子的作者是当前这个用户则可以删除， 否则不可以
	if in.AuthorId != post.AuthorId {
		logx.WithContext(l.ctx).Errorf("delete post failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PERMISSION_DENIED), "delete post failed")
	}

	// 3. 删除帖子
	err = l.svcCtx.PostModel.Delete(l.ctx, in.PostId)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("delete post failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "delete post failed")
	}

	// TODO: 这里不需要返回东西
	return &pb.DeletePostResponse{
		Success: true,
	}, nil
}
