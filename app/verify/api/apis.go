package main

import (
	"cointiger.com/verification/app/verify/api/internal/config"
	"cointiger.com/verification/app/verify/api/internal/handler"
	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/common/instance"
	"cointiger.com/verification/common/middleware"

	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	_ "go.uber.org/automaxprocs"
)

var configFile = flag.String("f", "etc/dev/apis.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	c.RestConf.Timeout = 600 * 1000
	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)

	defer server.Stop()
	logx.SetLevel(logx.InfoLevel)
	handler.RegisterHandlers(server, ctx)
	instance.NewRedis(c.VerifyRdb.Addr, c.VerifyRdb.Passwd)
	server.Use(middleware.PublicRateLimit)
	server.Use(middleware.JwtAuth)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
