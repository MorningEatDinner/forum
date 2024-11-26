package profile

import (
	"context"

	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"forum/app/user/rpc/userservice"
	"forum/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get user info
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取用户的信息
func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// 从jwt中可以获得当下的用户id
	userId := ctxdata.GetUidFromCtx(l.ctx)
	userResp, err := l.svcCtx.UserRpc.GetUserDetail(l.ctx, &userservice.UserInfoRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UserInfoResp{}
	copier.Copy(resp, userResp)

	return
}
