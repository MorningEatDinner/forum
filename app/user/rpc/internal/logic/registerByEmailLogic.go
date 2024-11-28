package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"forum/app/mqueue/cmd/job/jobtype"
	"forum/app/user/model"
	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"
	"forum/common/globalkey"
	"forum/common/tool"
	"forum/common/xerr"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterByEmailLogic {
	return &RegisterByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const NotifyUserUpdateTimeHours = 24

func (l *RegisterByEmailLogic) RegisterByEmail(in *pb.RegisterByEmailRequest) (*pb.RegisterByEmailResponse, error) {
	// 1. 验证验证码是否正确
	key := fmt.Sprintf(globalkey.GetRedisKey(globalkey.EmailCodeKey), in.Email)
	val, err := l.svcCtx.RedisClient.Get(key)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to get code from Redis: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to get code from Redis: %v", err)
	}
	if val != in.Code {
		logx.WithContext(l.ctx).Infof("Redis key used for verification: %s", key)
		logx.WithContext(l.ctx).Errorf("code mismatch: expected %v, got %v", val, in.Code)
		return nil, errors.Wrapf(xerr.NewErrMsg("code mismatch"), "verification code does not match")
	}

	// 2. 验证密码是否匹配
	if in.Password != in.PasswordConfirm {
		logx.WithContext(l.ctx).Error("passwords do not match")
		return nil, errors.Wrapf(xerr.NewErrMsg("password mismatch"), "passwords do not match")
	}

	// 3. 验证邮箱是否已经注册
	res, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil && err != sql.ErrNoRows {
		logx.WithContext(l.ctx).Errorf("failed to query email: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to query email: %v", err)
	}
	if res != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("email already registered"), "email %s already exists", in.Email)
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
		Email:       in.Email,
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

	// 启动延迟任务
	payload, err := json.Marshal(jobtype.DeferNotifyUserPayload{
		UserId:   id,
		Email:    user.Email,
		UserName: user.Username,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to marshal payload for userId: %d, err: %v", id, err)
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to marshal payload for userId"), "failed to marshal payload for userId: %d", id)
	}

	// 投递消息
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferEmailNotifyJob, payload), asynq.ProcessIn(time.Second*NotifyUserUpdateTimeHours))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to enqueue task for userId: %d, err: %v", id, err)
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to enqueue task for userId"), "failed to enqueue task for userId: %d", id)
	}

	return &pb.RegisterByEmailResponse{
		AccessToken:  tokenResp.AccessToken,
		ExpiresIn:    tokenResp.AccessExpire,
		RefreshToken: tokenResp.RefreshToken,
	}, nil
}
