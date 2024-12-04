package main

import (
	"flag"
	"fmt"

	"forum/app/post/api/internal/config"
	"forum/app/post/api/internal/handler"
	"forum/app/post/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/postcenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	fmt.Println("Starting server at %s:%d...", c.Host, c.Port)
	server.Start()
}
