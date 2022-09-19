package autoapi

/*
	FEE_BTCCNY_BTC(101001, "1", "01", "btc", "btccny", "公司：BTC-CNY交易对中BTC余额"),
	FEE_BTCCNY_CNY(101002, "1", "01", "cny", "btccny", "公司：BTC-CNY交易对中CNY余额"),
	FEE_ETHCNY_ETH(101003, "1", "01", "eth", "ethcny", "公司：ETH-CNY交易对中ETH余额"),
	FEE_ETHCNY_CNY(101004, "1", "01", "cny", "ethcny", "公司：ETH-CNY交易对中CNY余额"),
	FEE_LTCCNY_LTC(101005, "1", "01", "ltc", "ltccny", "公司：LTC-CNY交易对中LTC余额"),
	FEE_LTCCNY_CNY(101006, "1", "01", "cny", "ltccny", "公司：LTC-CNY交易对中CNY余额"),
	FEE_BCHCNY_BCH(101007, "1", "01", "bch", "bchcny", "公司：BCH-CNY交易对中BCH余额"),
	FEE_BCHCNY_CNY(101008, "1", "01", "cny", "bchcny", "公司：BCH-CNY交易对中CNY余额"),
	FEE_ETCCNY_ETC(101009, "1", "01", "etc", "etccny", "公司：ETC-CNY交易对中ETC余额"),
	FEE_ETCCNY_CNY(101010, "1", "01", "cny", "etccny", "公司：ETC-CNY交易对中CNY余额"),
*/
func (a *autoConfig) GetAccountType(typeId string) *AccountType {
	return a.GetSingleAccountTypeById(typeId)
}

/*
	AccType.properties
	FEE_BTCCNY_BTC=101001,101,btc,btccny
	FEE_BTCCNY_CNY=101002,101,cny,btccny
	FEE_ETHCNY_ETH=101003,101,eth,ethcny
	FEE_ETHCNY_CNY=101004,101,cny,ethcny
	FEE_LTCCNY_LTC=101005,101,ltc,ltccny
	FEE_LTCCNY_CNY=101006,101,cny,ltccny
	FEE_BCHCNY_BCH=101007,101,bch,bchcny
	FEE_BCHCNY_CNY=101008,101,cny,bchcny
	FEE_ETCCNY_ETC=101009,101,etc,etccny
	FEE_ETCCNY_CNY=101010,101,cny,etccny
*/
func (a *autoConfig) GetPersionType(accountType, coinSymbol string) string {
	v := a.GetSingleAccountTypeByabcCoin("2", accountType, coinSymbol)
	if v != nil {
		return v.Value
	}
	return ""
}

//硬编码 event.push.user
func GetEventPushUser() []string {
	return []string{"2816128"}
}
