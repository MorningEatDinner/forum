package main

import (
	"context"
	"flag"
	"fmt"

	"forum/app/user/api/internal/config"
	"forum/app/user/api/internal/handler"
	"forum/app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	// 配置日志信息
	logx.MustSetup(logx.LogConf{
		ServiceName: "user-api",
		Mode:        "file",
		Path:        "/var/log/user-api", // 改为与 volumeMount 中配置的 mountPath 一致
		Compress:    true,                // 是否压缩
		KeepDays:    7,                   // 日志保留天数
	})

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	logx.WithContext(context.Background()).Info("容器启动成功")

	server.Start()
}
