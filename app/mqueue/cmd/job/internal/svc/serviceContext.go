package svc

import (
	"forum/app/mqueue/cmd/job/internal/config"
	"forum/app/post/rpc/postservice"
	"forum/app/user/rpc/userservice"
	"forum/app/vote/rpc/voteservice"
	"forum/common/mail"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	MailClient  *mail.Mailer
	RedisClient *redis.Redis
	UserRpc     userservice.UserService
	PostRpc     postservice.PostService
	VoteRpc     voteservice.VoteService
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
		UserRpc:     userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
		PostRpc:     postservice.NewPostService(zrpc.MustNewClient(c.PostRpcConf)),
		VoteRpc:     voteservice.NewVoteService(zrpc.MustNewClient(c.VoteRpcConf)),
	}
}
