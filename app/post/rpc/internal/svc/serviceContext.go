package svc

import (
	"forum/app/mqueue/cmd/job/jobtype"
	"forum/app/post/model"
	"forum/app/post/rpc/internal/config"
	"forum/app/user/rpc/userservice"
	"forum/app/vote/rpc/voteservice"
	"github.com/hibiken/asynq"
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
	AsynqScheduler    *asynq.Scheduler
	RabbitMqClient    rabbitmq.Sender
	VoteRpc           voteservice.VoteService
	UserRpc           userservice.UserService
	SingleFlightGroup singleflight.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := &ServiceContext{
		Config:         c,
		PostModel:      model.NewPostsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		RedisClient:    redis.MustNewRedis(c.Redis.RedisConf),
		RabbitMqClient: rabbitmq.MustNewSender(c.RabbitSenderConf),
		AsynqScheduler: asynq.NewScheduler(asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		}, nil),
		VoteRpc: voteservice.NewVoteService(zrpc.MustNewClient(c.VoteRpcConf)),
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
	}
	registerTask(ctx)
	return ctx
}

// 注册定时任务
func registerTask(ctx *ServiceContext) {
	registerCronHotPostPushing(ctx)
	registerCronDeletePost(ctx)
}

func registerCronHotPostPushing(ctx *ServiceContext) {
	ctx.AsynqScheduler.Register("0 0 10 ? * FRI", asynq.NewTask(jobtype.ScheduleHotPostPushing, nil))
}

func registerCronDeletePost(ctx *ServiceContext) {
	ctx.AsynqScheduler.Register("0 0 2 * * ?", asynq.NewTask(jobtype.ScheduleDeletePost, nil))
}
