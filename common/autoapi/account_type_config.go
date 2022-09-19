package autoapi

import (
	"cointiger.com/verification/common/client"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type AccountTypeConf struct {
	AccountDesc    string `json:"accountDesc"`    //描述
	AccountType    int64  `json:"accountType"`    //账户类型, 101: 余额
	AccountTypeSub int64  `json:"accountTypeSub"` //子类型，如1347
	AccountValue   int64  `json:"accountValue"`   //交易对中coin余额对应的account type值,type+typesub
	BusinessKey    string `json:"businessKey"`    //业务名称
	BusinessType   int64  `json:"businessType"`   //业务类型 1公司 2个人
	CoinName       string `json:"coinName"`       //币种名称
}

type AccountType struct {
	Name        string `json:"name"`
	Value       string `json:"value"`       //账户type值
	AccountA    string `json:"accountA"`    //type的第1位
	AccountBC   string `json:"accountBC"`   //type的23位
	CoinSymbol  string `json:"coinSymbol"`  //货币符号
	Symbol      string `json:"symbol"`      //交易对   公司账户才有这个字段
	Description string `json:"description"` //账户类型描述
}
type AccountTypeResp struct {
	Code int64              `json:"code"`
	Msg  string             `json:"msg"`
	Data []*AccountTypeConf `json:"data"`
}
type AccountTypeSingleResp struct {
	Code int64            `json:"code"`
	Msg  string           `json:"msg"`
	Data *AccountTypeConf `json:"data"`
}

func (a *autoConfig) getAllAccountTypes() {
	pageIndex := int64(1)
	for {
		data := new(AccountTypeResp)
		data.Data = make([]*AccountTypeConf, 0)
		err := client.Get(map[string]string{
			"pageNum":  strconv.FormatInt(pageIndex, 10),
			"pageSize": "100",
		}, &data, a.c.AccountTypeServiceAll.Host, a.c.AccountTypeServiceAll.Path)
		if err != nil {
			logx.Errorf("get all account type  config error:%v", err)
		}
		if err == nil && len(data.Data) > 0 {
			for _, accType := range data.Data {
				//typeid规整
				val, _ := json.Marshal(&AccountType{
					Value:       strconv.FormatInt(accType.AccountValue, 10),
					AccountA:    getAccountTypeA(accType.AccountValue),
					AccountBC:   getAccountTypeBC(accType.AccountValue),
					CoinSymbol:  accType.CoinName,
					Symbol:      "",
					Description: accType.AccountDesc,
				})
				a.redisCli.Set(a.getAccountTypeKey(strconv.FormatInt(accType.AccountValue, 10)), string(val))
				//abccoin规整
				a.redisCli.Set(a.getABCCOinKey(getAccountTypeA(accType.AccountValue), getAccountTypeBC(accType.AccountValue), accType.CoinName), string(val))
			}
		}
		pageIndex++
		if err == nil && len(data.Data) == 0 {
			break
		}
	}
}

func (a *autoConfig) SyncSingleAccountTypeConfig(symbol string) error {
	data := new(AccountTypeSingleResp)
	err := client.Get(map[string]string{
		"coinSymbol": symbol,
	}, &data, a.c.AccountTypeServiceSingle.Host, a.c.AccountTypeServiceSingle.Path)
	if err != nil {
		logx.Errorf("get all account type  config error:%v", err)
	}
	if err == nil && data.Data != nil {
		accType := data.Data
		//typeid规整
		val, _ := json.Marshal(&AccountType{
			Value:       strconv.FormatInt(accType.AccountValue, 10),
			AccountA:    getAccountTypeA(accType.AccountValue),
			AccountBC:   getAccountTypeBC(accType.AccountValue),
			CoinSymbol:  accType.CoinName,
			Symbol:      "",
			Description: accType.AccountDesc,
		})
		a.redisCli.Set(a.getAccountTypeKey(strconv.FormatInt(accType.AccountValue, 10)), string(val))
		//abccoin规整
		a.redisCli.Set(a.getABCCOinKey(getAccountTypeA(accType.AccountValue), getAccountTypeBC(accType.AccountValue), accType.CoinName), string(val))
	}
	return err
}

func (a *autoConfig) GetSingleAccountTypeById(typeid string) *AccountType {
	val, err := a.redisCli.Get(a.getAccountTypeKey(typeid))
	if err != nil {
		return nil
	}
	v := new(AccountType)
	err = json.Unmarshal([]byte(val), v)
	if err != nil {
		return nil
	}
	return v
}
func (c *autoConfig) GetSingleAccountTypeByabcCoin(a, bc, coin string) *AccountType {
	val, err := c.redisCli.Get(c.getABCCOinKey(a, bc, coin))
	if err != nil {
		return nil
	}
	v := new(AccountType)
	err = json.Unmarshal([]byte(val), v)
	if err != nil {
		return nil
	}
	return v
}

func getAccountTypeA(i int64) string {
	if i < 100 {
		return ""
	}
	return strconv.FormatInt(i, 10)[0:1]
}
func getAccountTypeBC(i int64) string {
	if i < 100 {
		return ""
	}
	return strconv.FormatInt(i, 10)[1:3]
}
