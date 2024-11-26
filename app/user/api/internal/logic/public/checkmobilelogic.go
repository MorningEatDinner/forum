package public

import (
	"context"

	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"forum/app/user/rpc/userservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckMobileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// check if mobile exists
func NewCheckMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckMobileLogic {
	return &CheckMobileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 输入手机号码， 检查手机号码是否存在
func (l *CheckMobileLogic) CheckMobile(req *types.CheckMobileReq) (resp *types.CheckMobileResp, err error) {
	checkResp, err := l.svcCtx.UserRpc.CheckMobile(l.ctx, &userservice.CheckMobileRequest{
		Phone: req.Phone,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("CheckMobile failed: %v", err)
		return
	}

	resp = &types.CheckMobileResp{}
	copier.Copy(resp, checkResp)

	return
}
