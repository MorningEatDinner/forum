package public

import (
	"context"

	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMobileCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get mobile verification code
func NewGetMobileCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMobileCodeLogic {
	return &GetMobileCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMobileCodeLogic) GetMobileCode(req *types.GetMobileCodeReq) (resp *types.GetMobileCodeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
