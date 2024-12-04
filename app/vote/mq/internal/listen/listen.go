package listen

import (
	"context"
	"forum/app/vote/mq/internal/config"
	"forum/app/vote/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

// back to all consumers
func Mqs(c config.Config) []service.Service {

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	//kq ：pub sub
	services = append(services, RabbitMqs(c, ctx, svcContext)...)

	return services
}
