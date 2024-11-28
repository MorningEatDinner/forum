package svc

import (
	"forum/app/mqueue/cmd/job/internal/config"
	"forum/common/mail"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	MailClient  *mail.Mailer
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		AsynqServer: newAsynqServer(c),
		MailClient:  mail.NewMailer(c.MailConf),
	}
}
