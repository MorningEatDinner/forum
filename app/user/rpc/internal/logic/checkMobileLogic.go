package logic

import (
	"context"
	"database/sql"

	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckMobileLogic {
	return &CheckMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 这里就是检查手机号码是否存在
func (l *CheckMobileLogic) CheckMobile(in *pb.CheckMobileRequest) (*pb.CheckMobileResponse, error) {
	phone := sql.NullString{String: in.Phone, Valid: in.Phone != ""}

	exist := true
	_, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, phone)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("CheckMobile: %v", err)
		exist = false
	}

	return &pb.CheckMobileResponse{
		Exist: exist,
	}, nil
}
