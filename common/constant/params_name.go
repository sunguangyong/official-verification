package constant

const (
	ApiKeyParamName     = "api-key"
	PathParamName       = "path"
	TimeSpanParamName   = "timespan"
	SignParamName       = "sign"
	SecretName          = "secret"
	SymbolBuyParamName  = "BUY"
	SymbolSellParamName = "SELL"
	StatusOpen          = int64(1)
	StatusClose         = int64(1)

	TRADE_BUY_STATUS          = "tradeBuyStatus"
	TRADE_SELL_STATUS         = "tradeSellStatus"
	MARKET_STATUS             = "marketStatus"
	JOB_STATUS                = "jobStatus"
	JOB_SECOND_STATUS         = "jobSecondStatus"
	API_BUY_STATUS            = "apiBuyStatus"
	API_SELL_STATUS           = "apiSellStatus"
	API_MARKET_STATUS         = "apiMarketStatus"
	API_ORDINARY_USERS_STATUS = "apiOrdinaryUsersStatus"
	MATCH_STATUS              = "matchStatus"
)

const (
	HttpHeaderUserId = "x-user-id"
	HttpHeaderSource = "x-source"
)

//api限流
const (
	ApiPrivateRateLimit = 1
	ApiPublicRateLimit  = 2
)

const (
	//trade property
	OPEN_MARKETS_LIMIT_MIN            = "open_markets.limit_min."
	OPEN_MARKETS_MARKET_BUY_MIN       = "open_markets.market_buy_min."
	OPEN_COINS_WITHDRAW_ONE_MIN       = "open_coins.withdraw_one_min."
	OPEN_COINS_WITHDRAW_ONE_MAX       = "open_coins.withdraw_one_max."
	OPEN_MARKETS_DEPTH_SELECT         = "open_markets.depth.select."
	OPEN_COINS_WITHDRAW_FEE_MIN       = "open_coins.withdraw_fee_min."
	OPEN_MARKETS_PARIC_DECIMAL_PLACES = "open_markets.paric.decimal.places."
	TRADE_AREA_MIN_LIMIT              = "trade.area.min.limit."
	TRADE_SYMBOL_MIN_LIMIT            = "trade.symbol.min.limit."
	OPEN_MARKETS_MARKET_SELL_MIN      = "open_markets.market_sell_min."
	MORE_THAN_TRADE_PRICE_RANGE       = "more.than.trade.price.range"
)

const (
	ORDER_TYPE_LIMIT  = "1" //限价单
	ORDER_TYPE_MAKERT = "2" //市价单
)

const (
	OrderLock   = 0 //冻结
	OrderNormal = 1 //正常
)

const (
	AccountNomal    = "01" //余额账户
	AccountLock     = "02" //冻结账户
	AccountWithdraw = "03" //提现中
)
const (
	//maker
	TradeMakerVip0 = int64(100) //挂单（maker）手续费
	TradeMakerVip1 = int64(101) //VIP1
	TradeMakerVip2 = int64(102) //VIP2
	TradeMakerVip3 = int64(103) //VIP3
	TradeMakerVip4 = int64(104) //VIP4
	TradeMakerVip5 = int64(105) //VIP5
	TradeMakerVip6 = int64(106) //VIP6
	TradeMakerVip7 = int64(107) //VIP7
	TradeMakerVip8 = int64(108) //VIP8
	TradeMakerVip9 = int64(109) //VIP9
	//taker
	TradeTakerVip0 = int64(200) //吃单（taker）手续费
	TradeTakerVip1 = int64(201) //VIP1
	TradeTakerVip2 = int64(202) //VIP2
	TradeTakerVip3 = int64(203) //VIP3
	TradeTakerVip4 = int64(204) //VIP4
	TradeTakerVip5 = int64(205) //VIP5
	TradeTakerVip6 = int64(206) //VIP6
	TradeTakerVip7 = int64(207) //VIP7
	TradeTakerVip8 = int64(208) //VIP8
	TradeTakerVip9 = int64(209) //VIP9
)

//做市商
const (
	MarketMakerOrderNew        = "new"         //新订单
	MarketMakerOrderPartFilled = "part_filled" //部分成交
	MarketMakerOrderFilled     = "filled"      //完全成交
	MarketMakerOrderCanceled   = "canceled"    //已撤销
	MarketMakerOrderExpired    = "expired"     //异常订单
	MarketMakerBuyMarket       = "buy-market"  //市价买
	MarketMakerSellMarket      = "sell-market" //市价卖
	MarketMakerBuyLimit        = "buy-limit"   //限价买
	MarketMakerSellLimit       = "sell-limit"  //限价卖
)

const (
	UserApiKeyTradeStatusOpen  = 1
	UserApiKeyTradeStatusClose = 0
)

const (
	QueryPrev = "prev" //向前
	QueryNext = "next" //向后
)

//获取订单DB映射
const (
	DBExOrderActiveInit          = int64(0)
	DBExOrderActiveNew           = int64(1)
	DBExOrderActiveFilled        = int64(2)
	DBExOrderActivePartFilled    = int64(3)
	DBExOrderActiveCanceled      = int64(4)
	DBExOrderActivePendingCancel = int64(5)
	DBExOrderActiveExpired       = int64(6)
)

//const (
//	INIT = 0 //"trade.order.status.init","初始订单"),
//	NEW_ = 1 //"trade.order.status.new","新订单"),
//	FILLED = 2 //"trade.order.status.filled","完全成交"),
//	PART_FILLED = 3 //"trade.order.status.part.filled","部分成交"),
//	CANCELED = 4 //"trade.order.status.canceled","已撤单"),
//	PENDING_CANCEL = 5 //"trade.order.status.cancel","待撤单"),
//	EXPIRED = 6 //"trade.order.status.expired","异常订单");
//)
const (
	SenceCancelOrder = "cancel_order"
	SenceCreateOrder = "create_order"
)

const (
	Kline1Min   = "1min"
	Kline5Min   = "5min"
	Kline15Min  = "15min"
	Kline30Min  = "30min"
	Kline60Min  = "60min"
	Kline2Hour  = "2hour"
	Kline4Hour  = "4hour"
	Kline8Hour  = "8hour"
	Kline1Day   = "1day"
	Kline1Week  = "1week"
	Kline1Month = "1month"
)
const (
	//排除的symbol
	ExcludeSymbol = "SAYBTC,BPTNBTC,BAIETH"
)

const (
	DepthStep0 = "step0"
	DepthStep1 = "step1"
	DepthStep2 = "step2"
)
const (
	DBExOrderActiveLimit  = int64(1)
	DBExOrderActiveMarket = int64(2)
	DBExOrderActiveStop   = int64(3)
)

const (
	DBExOrderPrefix = "ex_order_"
	DBExTradePrefix = "ex_trade_"
)

var SIDE = map[string]struct{}{
	SymbolBuyParamName:  {},
	SymbolSellParamName: {},
}

var LimitMarkert = map[string]struct{}{
	ORDER_TYPE_LIMIT:  {},
	ORDER_TYPE_MAKERT: {},
}

//做市商订单分组类型
var MarketMakerOrderTypes = map[string]struct{}{
	MarketMakerBuyMarket:  {},
	MarketMakerSellMarket: {},
	MarketMakerBuyLimit:   {},
	MarketMakerSellLimit:  {},
}

//订单查询方向
var OrderQueryDirection = map[string]struct{}{
	QueryPrev: {},
	QueryNext: {},
}

//做市商订单状态类型
var MarketMakerStates = map[string]struct{}{
	MarketMakerOrderFilled:   {},
	MarketMakerOrderCanceled: {},
	MarketMakerOrderExpired:  {},
}

//委托单状态
var OrderStates = map[string]struct{}{
	MarketMakerOrderNew:        {},
	MarketMakerOrderPartFilled: {},
	MarketMakerOrderExpired:    {},
}

//K线刻度
var KlineScale = map[string]struct{}{
	Kline1Min:   {},
	Kline5Min:   {},
	Kline15Min:  {},
	Kline30Min:  {},
	Kline60Min:  {},
	Kline2Hour:  {},
	Kline4Hour:  {},
	Kline8Hour:  {},
	Kline1Day:   {},
	Kline1Week:  {},
	Kline1Month: {},
}

const (
	OrderSourceWeb          = "1" //web
	OrderSourceApp          = "2" //app
	OrderSourceApi          = "3" //api
	OrderSourceMarketMaking = "4" //对敲 marketMaking
)

//OrderSource
var OrderSource = map[string]struct{}{
	OrderSourceWeb:          {},
	OrderSourceApp:          {},
	OrderSourceApi:          {},
	OrderSourceMarketMaking: {},
}

var DepthStep = map[string]struct{}{
	DepthStep0: {},
	DepthStep1: {},
	DepthStep2: {},
}

// new cancel order开关
var CancelOrderSwitch = true

// new cancel order symbols
var CancelOrderSymbols = map[string]struct{}{
	"chhusdt":       {},
	"ethusdt":       {},
	"squidgrowusdt": {},
	"bosonusdt":     {},
	"gvrusdt":       {},
}
