package public

import (
	"context"

	"forum/app/post/api/internal/svc"
	"forum/app/post/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取帖子详情
func NewGetPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostDetailLogic {
	return &GetPostDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostDetailLogic) GetPostDetail(req *types.GetPostDetailReq) (resp *types.GetPostDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
