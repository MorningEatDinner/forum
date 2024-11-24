package logic

import (
	"context"

	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"

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
	// todo: add your logic here and delete this line

	return &pb.UpdateMobileResponse{}, nil
}
