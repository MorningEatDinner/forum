package svc

import (
	"forum/app/vote/model"
	"forum/app/vote/rpc/internal/config"

	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	RabbitMqClient rabbitmq.Sender

	RedisClient     *redis.Redis
	VoteRecordModel model.VoteRecordModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		RedisClient:     redis.MustNewRedis(c.Redis.RedisConf),
		VoteRecordModel: model.NewVoteRecordModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		RabbitMqClient:  rabbitmq.MustNewSender(c.RabbitSenderConf),
	}
}
