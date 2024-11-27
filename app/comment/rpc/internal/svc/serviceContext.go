package svc

import (
	"forum/app/comment/model"
	"forum/app/comment/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
