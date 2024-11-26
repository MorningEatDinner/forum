package logic

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"forum/app/user/model"
	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"
	"forum/common/globalkey"
	"forum/common/tool"
	"forum/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// 1. 验证验证码是否正确
	key := fmt.Sprintf(globalkey.GetRedisKey(globalkey.PhoneCodeKey), *in.Phone)
	val, err := l.svcCtx.RedisClient.Get(key)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to get code from database: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to get code from database: %v", err)
	}
	if val != *in.Code {
		logx.WithContext(l.ctx).Infof("Redis key used for verification: %s", key)
		logx.WithContext(l.ctx).Errorf("code mismatch: expected %v, got %v", val, *in.Code)
		return nil, errors.Wrapf(xerr.NewErrMsg("code mismatch"), "verification code does not match")
	}

	// 2. 验证密码是否匹配
	if in.Password != in.PasswordConfirm {
		logx.WithContext(l.ctx).Error("passwords do not match")
		return nil, errors.Wrapf(xerr.NewErrMsg("password mismatch"), "passwords do not match")
	}

	// 3. 验证手机号码是否存在
	phone := sql.NullString{
		String: *in.Phone,
		Valid:  *in.Phone != "",
	}
	res, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, phone)
	if err != nil && err != sql.ErrNoRows {
		logx.WithContext(l.ctx).Errorf("failed to query phone number: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to query phone number: %v", err)
	}
	if res != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("phone already registered"), "phone number %s already exists", phone.String)
	}

	// 4. 验证用户名是否已经存在
	res, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Name)
	if err != nil && err != sql.ErrNoRows {
		logx.WithContext(l.ctx).Errorf("failed to query username: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to query username: %v", err)
	}
	if res != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("username already in use"), "username %s already exists", in.Name)
	}

	// 5. 注册用户
	user := &model.Users{
		Username:    in.Name,
		Password:    tool.EncryptPassword(in.Password),
		Phone:       phone,
		CreateTime:  time.Now(),
		UpdatedTime: time.Now(),
	}
	insertRes, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to insert user: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to insert user: %v", err)
	}

	// 注册成功之后需要生成jwt token
	// 6. 生成token
	id, _ := insertRes.LastInsertId()
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: id,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to generate token for userId: %d, err: %v", id, err)
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to generate token for userId"), "failed to generate token for userId: %d", id)
	}

	return &pb.RegisterResponse{
		AccessToken:  tokenResp.AccessToken,
		ExpiresIn:    tokenResp.AccessExpire,
		RefreshToken: tokenResp.RefreshToken,
	}, nil
}
