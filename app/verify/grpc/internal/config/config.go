package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	VerifyMysql struct {
		DataSource string
	}

	VerifyRdb struct {
		Addr   string
		Passwd string
	}
}
