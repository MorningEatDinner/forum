package logic

import (
	"context"

	"forum/app/community/rpc/internal/svc"
	"forum/app/community/rpc/pb"
	"forum/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
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
	var page, pageSize int32
	if in.Page != nil {
		page = *in.Page
	} else {
		page = 1
	}
	if in.PageSize != nil {
		pageSize = *in.PageSize
	} else {
		size, err := l.svcCtx.CommunityModel.Count(l.ctx)
		if err != nil {
			return nil, errors.New("get community count failed")
		}
		pageSize = int32(size)
	}
	list, err := l.svcCtx.CommunityModel.GetCommunityList(l.ctx, page, pageSize)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get community list failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get community list failed")
	}

	var respList []*pb.Community
	if len(list) > 0 {
		for _, com := range list {
			var pbCommunity pb.Community
			copier.Copy(&pbCommunity, com)

			respList = append(respList, &pbCommunity)
		}
	}

	return &pb.GetAllCommunitiesResponse{
		Total:       int64(len(respList)),
		Communities: respList,
	}, nil
}
