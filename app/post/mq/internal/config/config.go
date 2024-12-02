package config

import (
	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	service.ServiceConf

	Redis redis.RedisConf

	DeleteUpMQConf   rabbitmq.RabbitListenerConf
	DeleteDownMQConf rabbitmq.RabbitListenerConf
}
