package logic

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"forum/app/user/model"
	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"
	"forum/common/globalkey"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoResponse, error) {
	// 修改用户信息
	// 1. 获取当前用户信息
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db error: %v", err)
	}

	if *in.Username != user.Username {
		// 2. 判断用户名是否重复
		newUser, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, *in.Username)
		if newUser != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NAME_EXISTS_ERROR), "username already exist")
		}
		if err != nil && err != model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "username already exist")
		}
	}

	// 3. 更新
	if in.Username != nil && *in.Username != "" {
		user.Username = *in.Username
	}
	user.City = sql.NullString{String: *in.City, Valid: *in.City != ""}
	user.Introduction = sql.NullString{String: *in.Introduction, Valid: *in.Introduction != ""}

	// 4. 写回数据库
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db error: %v", err)
	}

	// 5. 更新用户状态为已经更新信息
	err = l.svcCtx.RedisClient.Setex(fmt.Sprintf(globalkey.GetRedisKey(globalkey.UpdatedKey), strconv.FormatInt(in.UserId, 10)), "1", 24*int(time.Hour.Seconds()))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db error: %v", err)
	}

	userResp := &pb.User{}
	copier.Copy(userResp, user)
	return &pb.UpdateUserInfoResponse{
		User: userResp,
	}, nil
}
