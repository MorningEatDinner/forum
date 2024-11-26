package svc

import (
	"forum/app/post/model"
	"forum/app/post/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	PostModel model.PostsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		PostModel: model.NewPostsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
