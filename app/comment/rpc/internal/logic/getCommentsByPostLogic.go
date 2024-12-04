package logic

import (
	"context"

	"forum/app/comment/rpc/internal/svc"
	"forum/app/comment/rpc/pb"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsByPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsByPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsByPostLogic {
	return &GetCommentsByPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据帖子ID获取评论
func (l *GetCommentsByPostLogic) GetCommentsByPost(in *pb.GetCommentsByPostRequest) (*pb.GetCommentsByPostResponse, error) {
	listResp, err := l.svcCtx.CommentModel.FindCommentListByPostId(l.ctx, in.PostId, int64(in.Page), int64(in.PageSize))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get comments by post failed, err: %v", err.Error())
		return nil, errors.Wrapf(xerr.NewErrMsg("get comments by post failed"), "get comments by post failed, err: %v", err.Error())
	}

	commentListResp := []*pb.Comment{}
	copier.Copy(&commentListResp, listResp)
	resp := &pb.GetCommentsByPostResponse{
		Comments: commentListResp,
		Total:    int64(len(listResp)),
	}
	logx.WithContext(l.ctx).Infof("get comments by post success, resp: %v", resp)

	return resp, nil
}
