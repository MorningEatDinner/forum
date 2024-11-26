package logic

import (
	"context"

	"forum/app/community/rpc/internal/svc"
	"forum/app/community/rpc/pb"

	"github.com/jinzhu/copier"
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
	// 1. 验证社区是否存在
	community, err := l.svcCtx.CommunityModel.FindOne(l.ctx, uint64(in.CommunityId))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get community details failed, err: %v", err)
		return nil, err
	}
	// 2. 获取社区详情

	// 3. 返回社区详情
	var modelCommunity pb.Community
	copier.Copy(&modelCommunity, community) // 修改了这里
	return &pb.GetCommunityDetailsResponse{
		Community: &modelCommunity,
	}, nil
}
