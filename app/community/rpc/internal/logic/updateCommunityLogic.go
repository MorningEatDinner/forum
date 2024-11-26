package logic

import (
	"context"

	"forum/app/community/rpc/internal/svc"
	"forum/app/community/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommunityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommunityLogic {
	return &UpdateCommunityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Admin endpoints
func (l *UpdateCommunityLogic) UpdateCommunity(in *pb.UpdateCommunityRequest) (*pb.UpdateCommunityResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateCommunityResponse{}, nil
}
