package logic

import (
	"context"
	"database/sql"
	"time"

	"forum/app/community/model"
	"forum/app/community/rpc/internal/svc"
	"forum/app/community/rpc/pb"
	"forum/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommunityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommunityLogic {
	return &CreateCommunityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommunityLogic) CreateCommunity(in *pb.CreateCommunityRequest) (*pb.CreateCommunityResponse, error) {
	// 1. 检查这个社区名称是否已经存在了
	community, err := l.svcCtx.CommunityModel.FindOneByCommunityName(l.ctx, in.CommunityName)
	if err != nil && err != model.ErrNotFound {
		logx.WithContext(l.ctx).Errorf("find community failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find community failed, err: %v", err)
	}
	if community != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.COMMUNITY_NAME_EXIST), "community name already exist")
	}

	// 2. 创建社区
	var introduction sql.NullString
	if in.Introduction != nil {
		introduction = sql.NullString{
			String: *in.Introduction,
			Valid:  true,
		}
	}
	community = &model.Communities{
		CommunityName: in.CommunityName,
		Introduction:  introduction,
		CreateTime:    time.Now(),
		UpdatedTime:   time.Now(),
	}
	insertRes, err := l.svcCtx.CommunityModel.Insert(l.ctx, community)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create community failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "create community failed, err: %v", err)
	}

	id, _ := insertRes.LastInsertId()
	// 3. 返回
	return &pb.CreateCommunityResponse{
		CommunityId: id,
	}, nil
}
