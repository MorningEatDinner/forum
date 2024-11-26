package logic

import (
	"context"

	"forum/app/community/rpc/internal/svc"
	"forum/app/community/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllCommunitiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllCommunitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllCommunitiesLogic {
	return &GetAllCommunitiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Public endpoints
func (l *GetAllCommunitiesLogic) GetAllCommunities(in *pb.GetAllCommunitiesRequest) (*pb.GetAllCommunitiesResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetAllCommunitiesResponse{}, nil
}
