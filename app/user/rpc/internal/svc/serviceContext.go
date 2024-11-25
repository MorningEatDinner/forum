package svc

import (
	"forum/app/user/model"
	"forum/app/user/rpc/internal/config"
	"forum/common/sms"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel   model.UsersModel
	RedisClient *redis.Redis
	SMSClient   *sms.Aliyun
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UserModel: model.NewUsersModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		SMSClient: sms.NewSmsClient(c.Sms),
	}
}
