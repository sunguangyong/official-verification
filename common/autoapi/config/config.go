package config

import "github.com/zeromicro/go-zero/core/stores/redis"

type Config struct {
	SymbolConfigServiceAll struct {
		Host string
		Path string
	}
	SymbolConfigServiceSingle struct {
		Host string
		Path string
	}
	AccountTypeServiceAll struct {
		Host string
		Path string
	}
	AccountTypeServiceSingle struct {
		Host string
		Path string
	}
	CoinConfServiceAll struct {
		Host string
		Path string
	}
	CoinConfServiceSingle struct {
		Host string
		Path string
	}
	AutoConfRedis redis.RedisConf
}
