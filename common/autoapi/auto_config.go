package autoapi

import (
	"cointiger.com/verification/common/autoapi/config"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strings"
)

const (
	accountTypeIdKeyPrefix     = "openapi:account:config:typeid:%s"
	accountABCCoinKeyPrefix    = "openapi:account:config:abccoin:%s_%s_%s"
	coinConfigKeyPrefix        = "openapi:coin:config:%s"
	symbolConfigKeyPrefix      = "openapi:symbol:config:%s"
	symbolTradeConfigKeyPrefix = "openapi:trade:hash:countcoin:%s"
)

func NewAutoConfig(c config.Config) AutoConfig {
	return &autoConfig{
		c:        c,
		redisCli: redis.New(c.AutoConfRedis.Host, redis.WithPass(c.AutoConfRedis.Pass)),
	}
}

type AutoConfig interface {
	//account
	GetSingleAccountTypeById(typeid string) *AccountType
	GetSingleAccountTypeByabcCoin(a, bc, coin string) *AccountType
	SyncSingleAccountTypeConfig(symbol string) error
	GetPersionType(accountType, coinSymbol string) string
	TradePropertyUtil(prefix, name string) (string, error)
	GetAccountType(typeId string) *AccountType
	//symbol
	JudgeSymbolExist(symbol string) bool
	GetSingleSymbolConfig(symbol string) (*SymbolConfig, error)
	SyncSingleSymbolConfig(symbol string) error
	GetTradeAreaSymbol(areaName string) (map[string]*SymbolConfig, error)
	GetBaseCountCoinBySymbol(symbol string) (*Symbol, error)
	//coin
	GetSingleCoinConf(coin string) (*CoinConfig, error)
	GetSymbolConfigByCoin(coin string) []*CurrencysResponse
	SyncCoinSingleConfig(coin string) error
	Sync() error
}

type autoConfig struct {
	c        config.Config
	redisCli *redis.Redis
}

func (a *autoConfig) Sync() error {
	a.getAllCoinConf()
	a.getAllSymbolsConfig()
	a.getAllAccountTypes()
	return nil
}

func (a *autoConfig) getAccountTypeKey(typeid string) string {
	return fmt.Sprintf(accountTypeIdKeyPrefix, strings.ToLower(typeid))
}
func (c *autoConfig) getABCCOinKey(a, bc, coin string) string {
	return fmt.Sprintf(accountABCCoinKeyPrefix, strings.ToLower(a), strings.ToLower(bc), strings.ToLower(coin))
}
func (a *autoConfig) getCoinConfigKey(coin string) string {
	return fmt.Sprintf(coinConfigKeyPrefix, strings.ToLower(coin))
}
func (a *autoConfig) getSymbolConfigKey(symbol string) string {
	return fmt.Sprintf(symbolConfigKeyPrefix, strings.ToLower(symbol))
}
func (a *autoConfig) getSymbolTradeConfigKey(countCoin string) string {
	return fmt.Sprintf(symbolTradeConfigKeyPrefix, strings.ToLower(countCoin))
}
