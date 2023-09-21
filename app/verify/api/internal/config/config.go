package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	VerifyRpcConf zrpc.RpcClientConf

	VerifyMysql struct {
		DataSource string
	}

	VerifyRdb struct {
		Addr   string
		Passwd string
	}
}
