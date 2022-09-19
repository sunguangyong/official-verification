package autoapi

import (
	"cointiger.com/verification/common/client"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type SymbolListResponse struct {
	Code int64           `json:"code"`
	Msg  string          `json:"msg"`
	Data []*SymbolConfig `json:"data"`
}
type ReqParams struct {
	PageNum  int64 `json:"pageNum"`
	PageSize int64 `json:"pageSize"`
}

func (a *autoConfig) getAllSymbolsConfig() {
	pageIndex := int64(1)
	for {
		data := new(SymbolListResponse)
		data.Data = make([]*SymbolConfig, 0)
		err := client.Get(map[string]string{
			"pageNum":  strconv.FormatInt(pageIndex, 10),
			"pageSize": "100",
		}, &data, a.c.SymbolConfigServiceAll.Host, a.c.SymbolConfigServiceAll.Path)
		if err != nil {
			logx.Errorf("get all symbols config error:%v", err)
		}
		if err == nil && len(data.Data) > 0 {
			for _, symbol := range data.Data {
				//symbols规整
				bs, _ := json.Marshal(symbol)
				a.redisCli.Set(a.getSymbolConfigKey(symbol.CoinSymbol), string(bs))

				//交易区数据规整
				a.redisCli.Hset(a.getSymbolTradeConfigKey(symbol.CountCoin), symbol.CoinSymbol, string(bs))
			}
		}
		pageIndex++
		if err == nil && len(data.Data) == 0 {
			break
		}
	}
}

func (a *autoConfig) GetSingleSymbolConfig(symbol string) (*SymbolConfig, error) {
	valStr, err := a.redisCli.Get(a.getSymbolConfigKey(symbol))
	if err != nil {
		return nil, err
	}
	v := new(SymbolConfig)
	err = json.Unmarshal([]byte(valStr), v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

type SymbolSingleResponse struct {
	Code int64         `json:"code"`
	Msg  string        `json:"msg"`
	Data *SymbolConfig `json:"data"`
}

func (a *autoConfig) SyncSingleSymbolConfig(symbol string) error {
	symbolResult := new(SymbolSingleResponse)
	symbolResult.Data = new(SymbolConfig)
	err := client.Get(map[string]string{
		"coinSymbol": symbol,
	}, symbolResult, a.c.SymbolConfigServiceSingle.Host, a.c.SymbolConfigServiceSingle.Path)
	if err != nil {
		logx.Errorf("get single symbol config error:%v", err)
		return err
	}
	if err == nil && symbolResult != nil {
		symbol := symbolResult.Data
		//symbols规整
		bs, _ := json.Marshal(symbol)
		a.redisCli.Set(a.getSymbolConfigKey(symbol.CoinSymbol), string(bs))

		//交易区数据规整
		a.redisCli.Hset(a.getSymbolTradeConfigKey(symbol.CountCoin), symbol.CoinSymbol, string(bs))
	}
	return nil
}

type SymbolConfig struct {
	AmountDigits           int64   `json:"amountDigits"`           //数量小数位 //OPEN_MARKETS_LIMIT_MIN = "open_markets.limit_min.";//OPEN_MARKETS_MARKET_SELL_MIN = "open_markets.market_sell_min.";
	ApiBuyStatus           int64   `json:"apiBuyStatus"`           //api买单开关
	ApiMarketStatus        int64   `json:"apiMarketStatus"`        //api做市商开关
	ApiOrdinaryUsersStatus int64   `json:"apiOrdinaryUsersStatus"` //api普通用户开关
	ApiSellStatus          int64   `json:"apiSellStatus"`          //api卖单开关
	AppShow                int64   `json:"appShow"`                //app展示名
	BaseCoin               string  `json:"baseCoin"`               //基准货币 --
	CoinOfflineTime        string  `json:"coinOfflineTime"`        //币对下线时间,每次下线更新时间
	CoinOnlineTime         string  `json:"coinOnlineTime"`         //币对上线时间
	CoinSymbol             string  `json:"coinSymbol"`             //币对
	CountCoin              string  `json:"countCoin"`              //计价货币 --
	DealDeep               float64 `json:"dealDeep"`               //盘口阴影深度（数量）
	Depth0                 int64   `json:"depth0"`                 //价格小数位 //OPEN_MARKETS_DEPTH_SELECT = "open_markets.depth.select."; //OPEN_MARKETS_PARIC_DECIMAL_PLACES = "open_markets.paric.decimal.places."; ==== 直接取depth0
	Depth1                 int64   `json:"depth1"`                 //价格小数位 //OPEN_MARKETS_DEPTH_SELECT = "open_markets.depth.select.";
	Depth2                 int64   `json:"depth2"`                 //价格小数位 //OPEN_MARKETS_DEPTH_SELECT = "open_markets.depth.select.";
	ExecuteJobIp           string  `json:"executeJobIp"`           //执行job服务器 --
	Fiat                   int64   `json:"fiat"`                   //参考数据库
	JobSecondStatus        int64   `json:"jobSecondStatus"`        //job第二个开关
	JobStatus              int64   `json:"jobStatus"`              //job开关
	Market                 string  `json:"market"`                 //所属区显示名
	MarketStatus           int64   `json:"marketStatus"`           //exchange交易页显示开关
	MatchStatus            int64   `json:"matchStatus"`            //撮合开关
	MinTradeAmount         float64 `json:"minTradeAmount"`         //最小交易额（限价买入&卖出、市价买入&卖出）(USDT) //OPEN_MARKETS_MARKET_BUY_MIN = "open_markets.market_buy_min."; //TRADE_AREA_MIN_LIMIT = "trade.area.min.limit."; //TRADE_SYMBOL_MIN_LIMIT = "trade.symbol.min.limit.";
	MinTradeQuantity       float64 `json:"minTradeQuantity"`       //最小交易量（小数位保持一致与数量小数位限价买入&卖出、市价卖出）
	NewCoin                int64   `json:"newCoin"`                //是否为新币0新币，1非新币
	PriceDigits            int64   `json:"priceDigits"`            //价格小数位
	Seq                    int64   `json:"seq"`                    //序号 --
	Status                 int64   `json:"status"`                 //1代表启用，0代表停用 --
	TradeBuyStatus         int64   `json:"tradeBuyStatus"`         //exchange买单开关
	TradeSellStatus        int64   `json:"tradeSellStatus"`        //exchange卖单开关
	WarningStatus          int64   `json:"warningStatus"`          //报警状态
	WebShow                int64   `json:"webShow"`                //web展示名
	Zone                   int64   `json:"zone"`                   //所属区ID
}

func (a *autoConfig) GetTradeAreaSymbol(areaName string) (map[string]*SymbolConfig, error) {
	m, err := a.redisCli.Hgetall(a.getSymbolTradeConfigKey(areaName))
	if err != nil {
		return nil, errors.New("not found")
	}
	r := make(map[string]*SymbolConfig)
	for k, v := range m {
		s := new(SymbolConfig)
		if err = json.Unmarshal([]byte(v), s); err == nil {
			r[k] = s
		}
	}
	return r, nil
}
