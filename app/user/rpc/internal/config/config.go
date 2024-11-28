package config

import (
	"forum/common/mail"
	"forum/common/sms"

	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Cache            cache.CacheConf
	RabbitSenderConf rabbitmq.RabbitSenderConf
	Sms              sms.SMSConfig
	MailConf         mail.Config
}
