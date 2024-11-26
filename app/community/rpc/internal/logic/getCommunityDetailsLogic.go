package logic

import (
	"context"

	"forum/app/community/rpc/internal/svc"
	"forum/app/community/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommunityDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommunityDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommunityDetailsLogic {
	return &GetCommunityDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommunityDetailsLogic) GetCommunityDetails(in *pb.GetCommunityDetailsRequest) (*pb.GetCommunityDetailsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetCommunityDetailsResponse{}, nil
}
