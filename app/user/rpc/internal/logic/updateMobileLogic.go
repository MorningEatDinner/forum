package logic

import (
	"context"
	"database/sql"
	"fmt"

	"forum/app/user/model"
	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"
	"forum/common/globalkey"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMobileLogic {
	return &UpdateMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMobileLogic) UpdateMobile(in *pb.UpdateMobileRequest) (*pb.UpdateMobileResponse, error) {
	// 1. 验证手机验证码是否正确
	key := fmt.Sprintf(globalkey.GetRedisKey(globalkey.PhoneCodeKey), in.NewPhone)
	val, err := l.svcCtx.RedisClient.Get(key)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to get code from database: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to get code from database: %v", err)
	}
	if val != in.Code {
		logx.WithContext(l.ctx).Infof("Redis key used for verification: %s", key)
		logx.WithContext(l.ctx).Errorf("code mismatch: expected %v, got %v", val, in.Code)
		return nil, errors.Wrapf(xerr.NewErrMsg("code mismatch"), "verification code does not match")
	}

	// 2. 验证手机号码是否存在
	phone := sql.NullString{
		String: in.NewPhone,
		Valid:  in.NewPhone != "",
	}
	res, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, phone)
	if err != nil && err != model.ErrNotFound {
		logx.WithContext(l.ctx).Errorf("failed to query phone number: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to query phone number: %v", err)
	}
	if res != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("phone already registered"), "phone number %s already exists", phone.String)
	}

	// 3. 获取当前用户信息
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to get user information: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to get user information: %v", err)
	}

	// 4. 修改手机号码
	user.Phone = phone

	// 5. 更新用户信息
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to update user information: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to update user information: %v", err)
	}

	userResp := &pb.User{}
	copier.Copy(userResp, user)

	return &pb.UpdateMobileResponse{
		User: userResp,
	}, nil
}
