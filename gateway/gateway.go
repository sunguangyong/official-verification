package main

import (
	"flag"

	"cointiger.com/verification/common/middleware"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/gateway"

)

var configFile = flag.String("f", "etc/gateway.yaml", "config file")

func main() {
	flag.Parse()
	var c gateway.GatewayConf
	conf.MustLoad(*configFile, &c)
	gw := gateway.MustNewServer(c)
	// 鉴权
	gw.Use(middleware.JwtAuth)
	defer gw.Stop()
	gw.Start()
}
