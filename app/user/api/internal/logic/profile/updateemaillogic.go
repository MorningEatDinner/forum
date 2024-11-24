package profile

import (
	"context"

	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// update email
func NewUpdateEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailLogic {
	return &UpdateEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEmailLogic) UpdateEmail(req *types.UpdateEmailReq) (resp *types.UpdateEmailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
