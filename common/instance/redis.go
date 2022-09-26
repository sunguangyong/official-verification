package instance

import (
	"cointiger.com/verification/common/constant"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"sync"
)

var (
	redisOnce sync.Once
	rl        *redis.Redis
)

func GetRedis() (instance *redis.Redis) {
	return rl
}

func NewRedis(addr, pass string) (instance *redis.Redis) {
	redisOnce.Do(func() {
		rl = redis.New(addr, redis.WithPass(pass))
		rl.Ping()
	})
	return rl
}

func GetRedisPriRateLimitKey(path, apiKey string) string {
	return fmt.Sprintf(fmt.Sprintf(constant.REDIS_PRI_RATE_LIMIT_KEY, path, apiKey))
}
func GetRedisPubRateLimitKey(path, ip string) string {
	return fmt.Sprintf(fmt.Sprintf(constant.REDIS_PUB_RATE_LIMIT_KEY, path, ip))
}

func GetRedisUserKey(userId string) string {
	return fmt.Sprintf("%s%s", constant.LOGIN_TOKENS, userId)
}

func GetRedisApiKey(apiKey string) string {
	//return fmt.Sprintf(fmt.Sprintf(constant.REDIS_USER_API_KEY, apiKey))
	return fmt.Sprintf(fmt.Sprintf(constant.REDIS_USERAPIKEY_UID, apiKey))
}
