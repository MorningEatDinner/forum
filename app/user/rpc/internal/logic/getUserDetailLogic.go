package logic

import (
	"context"

	"forum/app/user/model"
	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailLogic {
	return &GetUserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserDetailLogic) GetUserDetail(in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	//  根据userid获取用户的详细信息
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to get user detail: %v", err)
	}
	if user == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_FOUND), "user not found")
	}

	userResp := &pb.User{}
	copier.Copy(userResp, user)
	return &pb.UserInfoResponse{
		User: userResp,
	}, nil
}
