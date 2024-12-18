package public

import (
	"context"

	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"forum/app/user/rpc/userservice"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
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
	getMobilResp, err := l.svcCtx.UserRpc.GetMobileCode(l.ctx, &userservice.GetMobileCodeRequest{
		Phone:       req.Phone,
		CaptchaId:   req.CaptchaId,
		CaptchaCode: req.CaptchaCode,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GetMobileCode: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("获取手机验证码失败"), "GetMobileCode: %v", err)
	}
	resp = &types.GetMobileCodeResp{}
	copier.Copy(resp, getMobilResp)

	return
}
