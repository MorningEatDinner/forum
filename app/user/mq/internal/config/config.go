package config

import (
	"forum/common/mail"
	"forum/common/sms"

	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	service.ServiceConf

	Redis redis.RedisConf

	EmailMQConf rabbitmq.RabbitListenerConf
	PhoneMQConf rabbitmq.RabbitListenerConf
	Sms         sms.SMSConfig
	MailConf    mail.Config
}
