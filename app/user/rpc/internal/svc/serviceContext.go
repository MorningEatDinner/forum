package svc

import (
	"forum/app/user/model"
	"forum/app/user/rpc/internal/config"
	"forum/common/sms"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	AsynqClient *asynq.Client

	UserModel      model.UsersModel
	RedisClient    *redis.Redis
	RabbitMqClient rabbitmq.Sender
	SMSClient      *sms.Aliyun
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := &ServiceContext{
		Config:      c,
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass}),

		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UserModel:      model.NewUsersModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		RabbitMqClient: rabbitmq.MustNewSender(c.RabbitSenderConf),
		SMSClient:      sms.NewSmsClient(c.Sms),
	}

	return ctx
}
