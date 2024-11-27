package post

import (
	"context"

	"forum/app/post/api/internal/svc"
	"forum/app/post/api/internal/types"
	"forum/app/post/rpc/postservice"
	"forum/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除帖子
func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(req *types.DeletePostReq) (resp *types.DeletePostResp, err error) {
	// 只有帖子的作者才可以删除帖子
	userid := ctxdata.GetUidFromCtx(l.ctx)
	delResp, err := l.svcCtx.PostRpc.DeletePost(l.ctx, &postservice.DeletePostRequest{
		PostId:   req.PostId,
		AuthorId: userid,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("delete post failed, err: %v", err)
		return nil, err
	}
	resp = &types.DeletePostResp{}
	copier.Copy(resp, delResp)

	// 删除帖子的时候， 评分表也需要删除对应的数据

	return
}
