package redismodel

type SymbolStatus struct {
	Id                     int64  `json:"id"`
	Symbol                 string `json:"symbol"`
	TradeBuyStatus         int64  `json:"tradeBuyStatus"`         //exchange买单开关
	TradeSellStatus        int64  `json:"tradeSellStatus"`        //exchange卖单开关
	MarketStatus           int64  `json:"marketStatus"`           //exchange交易页显示开关
	JobStatus              int64  `json:"jobStatus"`              //job开关
	JobSecondStatus        int64  `json:"jobSecondStatus"`        //job第二个开关
	ApiBuyStatus           int64  `json:"apiBuyStatus"`           //api买单开关
	ApiSellStatus          int64  `json:"apiSellStatus"`          //api卖单开关
	ApiMarketStatus        int64  `json:"apiMarketStatus"`        //api做市商开关
	MatchStatus            int64  `json:"matchStatus"`            //撮合开关
	ApiOrdinaryUsersStatus int64  `json:"apiOrdinaryUsersStatus"` //api普通用户开关
	WarningStatus          int64  `json:"warningStatus"`
	//Ctime                  time.Time `json:"ctime"`
	//Mtime                  time.Time `json:"mtime"`
}
