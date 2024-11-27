package public

import (
	"context"

	"forum/app/comment/rpc/commentservice"
	"forum/app/post/api/internal/svc"
	"forum/app/post/api/internal/types"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取评论列表
func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentListLogic) GetCommentList(req *types.GetCommentListReq) (resp *types.GetCommentListResp, err error) {
	// todo: add your logic here and delete this line
	getResp, err := l.svcCtx.CommentRpc.GetCommentsByPost(l.ctx, &commentservice.GetCommentsByPostRequest{
		PostId:   req.PostId,
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get comment list failed, err: %v", err.Error())
		return nil, errors.Wrapf(xerr.NewErrMsg("get comment list failed,"), "get comment list failed, err: %v", err.Error())
	}
	// TODO: 可以再查询作者名字和简介信息等
	resp = &types.GetCommentListResp{}
	copier.Copy(resp, getResp)
	return
}
