package svc

import (
	"forum/app/vote/model"
	"forum/app/vote/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	VoteRecordModel model.VoteRecordModel
	VoteCountModel  model.VoteCountModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		VoteRecordModel: model.NewVoteRecordModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		VoteCountModel:  model.NewVoteCountModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
