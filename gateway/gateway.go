package main

import (
	"context"
	"flag"

	"cointiger.com/verification/common/middleware"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

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

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md := metadata.New(map[string]string{"name": "lsz"})
	ctx = metadata.NewOutgoingContext(ctx, md)
	logx.Info("调用rpc服务前")

	err := invoker(ctx, method, req, reply, cc)
	if err != nil {
		return err
	}
	logx.Info("调用rpc服务后")
	return nil
}
