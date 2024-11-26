package logic

import (
	"context"
	"database/sql"

	"forum/app/user/model"
	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"
	"forum/common/tool"
	"forum/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user *model.Users
	var err error
	logx.WithContext(l.ctx).Infof("Login request - Email: %v, Phone: %v, Username: %v", *in.Email, *in.Phone, *in.Username)
	if in.Email != nil && *in.Email != "" {
		// 邮箱登录
		user, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, *in.Email)
	} else if in.Phone != nil && *in.Phone != "" {
		// 手机登录
		user, err = l.svcCtx.UserModel.FindOneByPhone(l.ctx, sql.NullString{String: *in.Phone, Valid: true})
	} else if in.Username != nil && *in.Username != "" {
		// 用户名登录
		user, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, *in.Username)
	} else {
		// 没有提供任何登录信息
		logx.WithContext(l.ctx).Errorf("no login information provided")
		return nil, errors.Wrapf(xerr.NewErrMsg("no login information provided"), "no login information provided")
	}
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to login: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to login: %v", err)
	}
	if user == nil {
		logx.WithContext(l.ctx).Errorf("user not found")
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_FOUND), "user not found")
	}

	// 验证密码
	if !tool.CheckPasswordHash(in.Password, user.Password) {
		logx.WithContext(l.ctx).Errorf("incorrect password")
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_PASSWORD_ERROR), "incorrect password")
	}

	// 如果都相等, 那么返回用户信息
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: user.UserId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", user.UserId)
	}

	return &pb.LoginResponse{
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
		ExpiresIn:    tokenResp.AccessExpire,
	}, nil
}
