package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	VerifyMysql struct {
		DataSource string
	}

	VerifyRdb struct {
		Addr   string
		Passwd string
	}
}
