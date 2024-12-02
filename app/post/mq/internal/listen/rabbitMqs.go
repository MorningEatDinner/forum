package listen

import (
	"context"
	"forum/app/post/mq/internal/config"
	"forum/app/post/mq/internal/mqs/rq"
	"forum/app/post/mq/internal/svc"
	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/core/service"
)

// pub sub use kq (kafka)
func RabbitMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	// 绑定配置信息到handler函数
	return []service.Service{
		rabbitmq.MustNewListener(c.DeleteUpMQConf, rq.NewDeletePostUpMq(ctx, svcContext)),
		rabbitmq.MustNewListener(c.DeleteDownMQConf, rq.NewDeletePostDownMq(ctx, svcContext)),
	}
}
