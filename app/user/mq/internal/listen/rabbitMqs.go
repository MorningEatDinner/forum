package listen

import (
	"context"
	"forum/app/user/mq/internal/config"
	"forum/app/user/mq/internal/mqs/rq"
	"forum/app/user/mq/internal/svc"

	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/core/service"
)

// pub sub use kq (kafka)
func RabbitMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	// 绑定配置信息到handler函数
	return []service.Service{
		rabbitmq.MustNewListener(c.EmailMQConf, rq.NewSendEmailCodeMq(ctx, svcContext)),
		rabbitmq.MustNewListener(c.PhoneMQConf, rq.NewSendPhoneCodeMq(ctx, svcContext)),
	}

}
