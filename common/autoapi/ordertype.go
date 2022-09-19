package autoapi

import (
	"cointiger.com/verification/common/constant"
	"strconv"
)

func OrderType(orderType int64, side string) string {
	if side == constant.SymbolBuyParamName && strconv.FormatInt(orderType, 10) == constant.ORDER_TYPE_LIMIT {
		return constant.MarketMakerBuyLimit
	}
	if side == constant.SymbolSellParamName && strconv.FormatInt(orderType, 10) == constant.ORDER_TYPE_LIMIT {
		return constant.MarketMakerSellLimit
	}
	if side == constant.SymbolBuyParamName && strconv.FormatInt(orderType, 10) == constant.ORDER_TYPE_MAKERT {
		return constant.MarketMakerBuyMarket
	}
	if side == constant.SymbolSellParamName && strconv.FormatInt(orderType, 10) == constant.ORDER_TYPE_MAKERT {
		return constant.MarketMakerSellMarket
	}
	return "stop"
}
