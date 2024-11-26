package svc

import (
	"forum/app/community/model"
	"forum/app/community/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	CommunityModel model.CommunitiesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CommunityModel: model.NewCommunitiesModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
