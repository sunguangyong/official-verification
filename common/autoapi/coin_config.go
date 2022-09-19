package autoapi

import (
	"cointiger.com/verification/common/client"
	"encoding/json"
	"errors"
	"github.com/valyala/fastjson"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type CoinConfResp struct {
	Code int64         `json:"code"`
	Msg  string        `json:"msg"`
	Data []*CoinConfig `json:"data"`
}

type CoinConfSingleResp struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data *CoinConfig `json:"data"`
}

func (a *autoConfig) getAllCoinConf() {
	pageIndex := int64(1)
	for {
		data := new(CoinConfResp)
		data.Data = make([]*CoinConfig, 0)
		err := client.Get(map[string]string{
			"pageNum":  strconv.FormatInt(pageIndex, 10),
			"pageSize": "100",
		}, &data, a.c.CoinConfServiceAll.Host, a.c.CoinConfServiceAll.Path)
		if err != nil {
			logx.Errorf("get all account type  config error:%v", err)
		}
		if err == nil && len(data.Data) > 0 {
			for _, coin := range data.Data {
				val, _ := json.Marshal(coin)
				a.redisCli.Set(a.getCoinConfigKey(coin.Coin), string(val))
			}
		}
		pageIndex++
		if err == nil && len(data.Data) == 0 {
			break
		}
	}
}

func (a *autoConfig) GetSingleCoinConf(coin string) (*CoinConfig, error) {
	valStr, err := a.redisCli.Get(a.getCoinConfigKey(coin))
	if err != nil {
		return nil, errors.New("not found")
	}
	conf := new(CoinConfig)
	var p fastjson.Parser
	v, err := p.Parse(valStr)
	//err = json.Unmarshal([]byte(valStr), conf)
	if err != nil {
		return nil, errors.New("not found")
	}
	conf.AddressType = v.Get("addressType").String()
	conf.AddressUrl = v.Get("addressUrl").String()
	conf.BigNum = v.Get("bigNum").String()
	conf.Coin = v.Get("coin").String()
	conf.CointTypeName = v.Get("cointTypeName").String()
	conf.Confirm = v.GetInt("confirm")
	conf.DepositCoin = v.Get("depositCoin").String()
	conf.DepositStatus = v.GetInt("depositStatus")
	conf.IsMain = v.GetInt("isMain")
	conf.LeasteDepositNum = v.Get("leasteDepositNum").String()
	conf.LegalCoin = v.GetInt("legalCoin")
	conf.Logo = v.Get("logo").String()
	conf.NeedMemo = v.GetInt("needMemo")
	conf.RealName = v.Get("realName").String()
	conf.RmbShowNum = v.GetInt("rmbShowNum")
	conf.ShowCoin = v.Get("showCoin").String()
	conf.ShowName = v.Get("showName").String()
	conf.Status = v.GetInt("status")
	conf.TxidUrl = v.Get("txidUrl").String()
	conf.Type = v.Get("type").String()
	conf.UsdtShowNum = v.GetInt("usdtShowNum")
	conf.ValuationCoin = v.GetInt("valuationCoin")
	conf.WithdrawFee = v.GetFloat64("withdrawFee")
	conf.WithdrawFeeMax = v.Get("withdrawFeeMax").String()
	conf.WithdrawFeeMin = v.Get("withdrawFeeMin").String()
	conf.WithdrawNumber = v.GetInt("WithdrawNumber")
	conf.WithdrawOneMax = v.Get("withdrawOneMax").String()
	conf.WithdrawOneMin = v.Get("withdrawOneMin").String()
	conf.WithdrawStatus = v.GetInt("withdrawStatus")
	return conf, nil
}

func (a *autoConfig) SyncCoinSingleConfig(coin string) error {
	data := new(CoinConfSingleResp)
	err := client.Get(map[string]string{
		"coin": coin,
	}, &data, a.c.CoinConfServiceSingle.Host, a.c.CoinConfServiceSingle.Path)
	if err != nil {
		logx.Errorf("get single coin  config error:%v", err)
	}
	if err == nil && data.Data != nil {
		val, _ := json.Marshal(data.Data)
		a.redisCli.Set(a.getCoinConfigKey(coin), string(val))
	}
	return err
}

type CoinConfig struct {
	AddressType      string  `json:"addressType"`      // 地址
	AddressUrl       string  `json:"addressUrl"`       //查地址的区块链接
	BigNum           string  `json:"bigNum"`           //大额上账
	Coin             string  `json:"coin"`             //币种
	CointTypeName    string  `json:"cointTypeName"`    // 链类型显示名
	Confirm          int     `json:"confirm"`          // 区块确认数
	DepositCoin      string  `json:"depositCoin"`      // 参考数据库
	DepositStatus    int     `json:"depositStatus"`    // 充币开关
	IsMain           int     `json:"isMain"`           //是否是主链0-否 1-是 目前只有提币接口使用
	LeasteDepositNum string  `json:"leasteDepositNum"` //最小充币数
	LegalCoin        int     `json:"legalCoin"`        // 是否是法币计价币，0不合法，1合法
	Logo             string  `json:"logo"`             //logo
	NeedMemo         int     `json:"needMemo"`         //0 不需要 1需要
	RealName         string  `json:"realName"`         //全称
	RmbShowNum       int     `json:"rmbShowNum"`       //rmb小数位
	ShowCoin         string  `json:"showCoin"`         //用于展示的coin名称
	ShowName         string  `json:"showName"`         //展示名
	Status           int     `json:"status"`           //状态
	TxidUrl          string  `json:"txidUrl"`          // 查txid
	Type             string  `json:"type"`             //erc20 trc20... 兼容老版本,如果不存在,应该取coin字段
	UsdtShowNum      int     `json:"usdtShowNum"`      // usdt小数位
	ValuationCoin    int     `json:"valuationCoin"`    // 是否是计价币，1是计价币，2不是计价币
	WithdrawFee      float64 `json:"withdrawFee"`      // 百分比手续费
	WithdrawFeeMax   string  `json:"withdrawFeeMax"`   //最大提币手续费
	WithdrawFeeMin   string  `json:"withdrawFeeMin"`   //最小提币手续费
	WithdrawNumber   int     `json:"withdrawNumber"`   //提币小数位
	WithdrawOneMax   string  `json:"withdrawOneMax"`   //最大提币数量
	WithdrawOneMin   string  `json:"withdrawOneMin"`   //最小提币数量
	WithdrawStatus   int     `json:"withdrawStatus"`   //提币开关
}
