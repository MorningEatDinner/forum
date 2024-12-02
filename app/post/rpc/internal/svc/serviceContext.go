package svc

import (
	"forum/app/post/model"
	"forum/app/post/rpc/internal/config"
	"forum/app/user/rpc/userservice"
	"forum/app/vote/rpc/voteservice"
	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"golang.org/x/sync/singleflight"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	PostModel         model.PostsModel
	RedisClient       *redis.Redis
	RabbitMqClient    rabbitmq.Sender
	VoteRpc           voteservice.VoteService
	UserRpc           userservice.UserService
	SingleFlightGroup singleflight.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		PostModel:      model.NewPostsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		RedisClient:    redis.MustNewRedis(c.Redis.RedisConf),
		RabbitMqClient: rabbitmq.MustNewSender(c.RabbitSenderConf),
		VoteRpc:        voteservice.NewVoteService(zrpc.MustNewClient(c.VoteRpcConf)),
		UserRpc:        userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
