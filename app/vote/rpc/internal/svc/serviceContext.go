package svc

import (
	"forum/app/vote/model"
	"forum/app/vote/rpc/internal/config"
	"forum/tmp/app/post/rpc/postservice"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	RedisClient     *redis.Redis
	VoteRecordModel model.VoteRecordModel
	PostRpc         postservice.PostService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		RedisClient:     redis.MustNewRedis(c.Redis.RedisConf),
		VoteRecordModel: model.NewVoteRecordModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		PostRpc:         postservice.NewPostService(zrpc.MustNewClient(c.PostRpcConf)),
	}
}
