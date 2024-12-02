package config

import (
	"forum/common/mail"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	service.ServiceConf
	Redis       redis.RedisConf
	MailConf    mail.Config
	UserRpcConf zrpc.RpcClientConf
	PostRpcConf zrpc.RpcClientConf
	VoteRpcConf zrpc.RpcClientConf
}
