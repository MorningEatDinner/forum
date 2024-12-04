package listen

import (
	"context"
	"forum/app/vote/mq/internal/config"
	"forum/app/vote/mq/internal/mqs/rq"
	"forum/app/vote/mq/internal/svc"

	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/core/service"
)

// pub sub use kq (kafka)
func RabbitMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	// 绑定配置信息到handler函数
	return []service.Service{
		rabbitmq.MustNewListener(c.PostUpdateMQConf, rq.NewUpdatePostScoreMq(ctx, svcContext)),
	}

}
