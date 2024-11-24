package logic

import (
	"context"

	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMobileCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMobileCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMobileCodeLogic {
	return &GetMobileCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMobileCodeLogic) GetMobileCode(in *pb.GetMobileCodeRequest) (*pb.GetMobileCodeResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMobileCodeResponse{}, nil
}
