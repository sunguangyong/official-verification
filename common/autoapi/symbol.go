package autoapi

import (
	"cointiger.com/verification/common/constant"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"math"
	"strconv"
	"strings"
)

/*
	判断symbol是否开启
	TCHBITCNY("tchbitcny", "tch", "bitcny", 1, "tchbitcny", 120010, "5-9999"),
	BTCBITCNY("btcbitcny", "btc", "bitcny", 1, "btcbitcny", 120020, "1-9999"),
	ETHBITCNY("ethbitcny", "eth", "bitcny", 1, "ethbitcny", 120030, "1-9999"),
	RXBITCNY("trxbitcny", "trx", "bitcny", 1, "trxbitcny", 120040, "1-9999"),
*/
func (a *autoConfig) JudgeSymbolExist(symbol string) bool {
	v, err := a.GetSingleSymbolConfig(strings.ToUpper(symbol))
	if err != nil {
		return false
	}
	if v.Status == constant.StatusOpen {
		return true
	}
	return false
}

//硬编码
var symbol_cant_buy = "saybtc,baieth,bchbtc,btabitcny"

func SymbolCantBuy(symbol string) bool {
	return strings.Contains(symbol_cant_buy, strings.ToLower(symbol))
}

//硬编码
var symbol_cant_sell = "saybtc,baieth,bchbtc,btabitcny"

func SymbolCantSell(symbol string) bool {
	return strings.Contains(symbol_cant_sell, strings.ToLower(symbol))
}

/*
	获取tradepropertyutils中使用的属性
	OPEN_MARKETS_LIMIT_MIN = "open_markets.limit_min.";
	OPEN_MARKETS_MARKET_BUY_MIN = "open_markets.market_buy_min.";
	OPEN_MARKETS_MARKET_SELL_MIN = "open_markets.market_sell_min.";
	OPEN_COINS_WITHDRAW_ONE_MIN = "open_coins.withdraw_one_min.";
	OPEN_COINS_WITHDRAW_ONE_MAX = "open_coins.withdraw_one_max.";
	OPEN_MARKETS_DEPTH_SELECT = "open_markets.depth.select.";
	OPEN_COINS_WITHDRAW_FEE_MIN = "open_coins.withdraw_fee_min.";
	OPEN_MARKETS_PARIC_DECIMAL_PLACES = "open_markets.paric.decimal.places.";
	TRADE_AREA_MIN_LIMIT = "trade.area.min.limit.";
	TRADE_SYMBOL_MIN_LIMIT = "trade.symbol.min.limit.";
	MORE_THAN_TRADE_PRICE_RANGE       = "more.than.trade.price.range"
*/
var more_than_trade_price_range_value = float64(5.0)

func (a *autoConfig) TradePropertyUtil(prefix, name string) (string, error) {
	if prefix == constant.MORE_THAN_TRADE_PRICE_RANGE {
		return strconv.FormatFloat(more_than_trade_price_range_value, 'E', 1, 64), nil
	}
	v, err := a.GetSingleSymbolConfig(strings.ToUpper(name))
	if err != nil {
		return "", err
	}
	if prefix == constant.OPEN_MARKETS_LIMIT_MIN || prefix == constant.OPEN_MARKETS_MARKET_SELL_MIN {
		return strconv.FormatFloat((1 / math.Pow10(int(v.AmountDigits))), 'E', int(v.AmountDigits), 64), nil

	}
	if prefix == constant.OPEN_MARKETS_DEPTH_SELECT || prefix == constant.OPEN_MARKETS_PARIC_DECIMAL_PLACES {
		return strconv.FormatInt(v.Depth0, 10), nil
	}
	if prefix == constant.OPEN_MARKETS_MARKET_BUY_MIN ||
		prefix == constant.TRADE_AREA_MIN_LIMIT ||
		prefix == constant.TRADE_SYMBOL_MIN_LIMIT {
		return strconv.FormatFloat(v.MinTradeAmount, 'E', 8, 64), nil
	}
	return "", nil
}

type Symbol struct {
	BaseCoin  string `json:"base_coin"`
	CountCoin string `json:"count_coin"`
}

/*
	根据symbol返回base/count coin
	BTCCNY("btccny", "btc", "cny", 0, "比特币人民币", 0, "0-1"),
	ETHCNY("ethcny", "eth", "cny", 0, "以太坊人民币", 0, "0-2"),
	LTCCNY("ltccny", "ltc", "cny", 0, "莱特币人民币", 0, "0-3"),
	BCHCNY("bchcny", "bch", "cny", 0, "bch人民币", 0, "0-4"),
	ETCCNY("etccny", "etc", "cny", 0, "以太经典人民币", 0, "0-5"),
	BTCKRW("btckrw", "btc", "krw", 0, "btckrw", 0, "0-6"),
	ETHKRW("ethkrw", "eth", "krw", 0, "ethkrw", 0, "0-7"),
	LTCKRW("ltckrw", "ltc", "krw", 0, "ltckrw", 0, "0-8"),
	BCCKRW("bchkrw", "bch", "krw", 0, "bchkrw", 0, "0-9"),
	ETCKRW("etckrw", "etc", "krw", 0, "etckrw", 0, "0-10"),
	LTCETH("ltceth", "ltc", "ETH", 0, "ltceth", 0, "0-11"),
	BCHETH("bcheth", "bch", "ETH", 0, "bcheth", 0, "0-12"),
	ETCETH("etceth", "etc", "ETH", 0, "etceth", 0, "0-13"),
*/
func (a *autoConfig) GetBaseCountCoinBySymbol(symbol string) (*Symbol, error) {
	v, err := a.GetSingleSymbolConfig(strings.ToLower(symbol))
	if err != nil {
		return nil, err
	}
	return &Symbol{
		BaseCoin:  v.BaseCoin,
		CountCoin: v.CountCoin,
	}, nil
}

//硬编码:ult.bitunivers.discount
func UltBitUniversDiscount() float64 {
	return 0.5
}

//硬编码:event.push.symbol
func GetEventPushSymbol() []string {
	return []string{"btcusdt"}
}

func XemReplaceNex(symbol string) string {
	if strings.Contains(symbol, "xem") {
		return strings.Replace(symbol, "xem", "nem", -1)
	}
	return symbol
}

type CurrencysResponse struct {
	BaseCurrency    string        `json:"baseCurrency"`    // 基础币种
	QuoteCurrency   string        `json:"quoteCurrency"`   // 计价币种
	PricePrecision  string        `json:"pricePrecision"`  // 价格精度位数（0为个位）
	AmountPrecision string        `json:"amountPrecision"` // 数量精度位数（0为个位）
	AmountMin       string        `json:"amountMin"`       // 最小交易数量
	WithdrawFeeMin  string        `json:"withdrawFeeMin"`  // 提币手续费最小值
	WithdrawFeeMax  string        `json:"withdrawFeeMax"`  // 提币手续费最小值
	WithdrawOneMin  string        `json:"withdrawOneMin"`  // 单笔提现最小限制
	WithdrawOneMax  string        `json:"withdrawOneMax"`  // 单笔提现最大限制
	DepthSelect     *SymbolSelect `json:"depthSelect"`     //深度选择配置
	minTurnover     float64       `json:"minTurnover"`
}
type SymbolSelect struct {
	Step0 string `json:"step0"`
	Step1 string `json:"step1"`
	Step2 string `json:"step2"`
}

/*
 	获取指定交易区的币种配置
	symbol/inner/coinSymbolConfigApi/queryCoinSymbolConfigByMarket
*/
func (a *autoConfig) GetSymbolConfigByCoin(coin string) []*CurrencysResponse {
	resp := make([]*CurrencysResponse, 0)
	if vals, err := a.GetTradeAreaSymbol(coin); err == nil {
		for _, v := range vals {
			coinConf, err := a.GetSingleCoinConf(v.BaseCoin)
			if err != nil {
				logx.Errorf("GetSymbolConfigByCoin val:%v,err:%v", v.CoinSymbol, err)
				continue
			}
			resp = append(resp, &CurrencysResponse{
				BaseCurrency:    v.BaseCoin,
				QuoteCurrency:   v.CountCoin,
				PricePrecision:  strconv.FormatInt(v.PriceDigits, 10),
				AmountPrecision: strconv.FormatInt(v.AmountDigits, 10),
				AmountMin:       fmt.Sprintf("%v", v.MinTradeQuantity),
				DepthSelect: &SymbolSelect{ //深度选择配置
					Step0: strconv.FormatInt(v.Depth0, 10),
					Step1: strconv.FormatInt(v.Depth1, 10),
					Step2: strconv.FormatInt(v.Depth2, 10),
				},
				WithdrawFeeMin: coinConf.WithdrawFeeMin, //提币手续费最小值
				WithdrawFeeMax: coinConf.WithdrawFeeMax, // 提币手续费最小值
				WithdrawOneMin: coinConf.WithdrawOneMin, // 单笔提现最小限制
				WithdrawOneMax: coinConf.WithdrawOneMax, // 单笔提现最大限制
			})
		}
	}
	return resp
}
