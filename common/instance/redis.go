package instance

import (
	"cointiger.com/verification/common/constant"
	"cointiger.com/verification/common/redismodel"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strings"
	"sync"
)

var (
	redisOnce sync.Once
	rl        *redis.Redis //ratelimit redis instance
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

func GetRedisBlacklistUserKey() string {
	return constant.REDIS_BLACKLIST_USER_KEY
}

func GetRedisApiKey(apiKey string) string {
	//return fmt.Sprintf(fmt.Sprintf(constant.REDIS_USER_API_KEY, apiKey))
	return fmt.Sprintf(fmt.Sprintf(constant.REDIS_USERAPIKEY_UID, apiKey))
}

func GetRedisUserApiKeyEntity(apiKey string) string {
	return fmt.Sprintf(fmt.Sprintf(constant.REDIS_USER_API_KEY_ENTITY_KEY, apiKey))
}

func GetRedisPriRateLimitKey(path, apiKey string) string {
	return fmt.Sprintf(fmt.Sprintf(constant.REDIS_PRI_RATE_LIMIT_KEY, path, apiKey))
}
func GetRedisPubRateLimitKey(path, ip string) string {
	return fmt.Sprintf(fmt.Sprintf(constant.REDIS_PUB_RATE_LIMIT_KEY, path, ip))
}

func UserIsMarker(uid string) (bool, error) {
	return rl.Sismember(constant.REDIS_MARKET_USER_KEY, uid)
}

func CheckStatus(symbol, projectName string) (bool, error) {
	projectName = strings.ToLower(strings.Trim(projectName, " "))
	key := fmt.Sprintf("%s%s", constant.REDIS_SYMBOL_STATUS_KEY, strings.ToLower(symbol))
	val, err := rl.GetCtx(context.Background(), key)
	if err != nil {
		return false, err
	}
	symbolStatus := new(redismodel.SymbolStatus)
	err = json.Unmarshal([]byte(val), symbolStatus)
	if err != nil {
		return false, err
	}
	if projectName == strings.ToLower(constant.TRADE_BUY_STATUS) {
		return symbolStatus.TradeBuyStatus == constant.StatusOpen, nil
	}
	if projectName == strings.ToLower(constant.TRADE_SELL_STATUS) {
		return symbolStatus.TradeSellStatus == constant.StatusOpen, nil
	}
	if projectName == strings.ToLower(constant.MARKET_STATUS) {
		return symbolStatus.MarketStatus == constant.StatusOpen, nil
	}
	if projectName == strings.ToLower(constant.JOB_STATUS) {
		return symbolStatus.JobStatus == constant.StatusOpen, nil
	}
	if projectName == strings.ToLower(constant.JOB_SECOND_STATUS) {
		return symbolStatus.JobSecondStatus == constant.StatusOpen, nil
	}
	if projectName == strings.ToLower(constant.API_BUY_STATUS) {
		return symbolStatus.ApiBuyStatus == constant.StatusOpen, nil
	}
	if projectName == strings.ToLower(constant.API_SELL_STATUS) {
		return symbolStatus.ApiSellStatus == constant.StatusOpen, nil
	}
	if projectName == strings.ToLower(constant.API_MARKET_STATUS) {
		return symbolStatus.ApiMarketStatus == constant.StatusOpen, nil
	}
	if projectName == strings.ToLower(constant.API_ORDINARY_USERS_STATUS) {
		return symbolStatus.ApiOrdinaryUsersStatus == constant.StatusOpen, nil
	}
	return false, errors.New("not match")
}
