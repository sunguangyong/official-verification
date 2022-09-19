package coin

import "strings"

func NEXtoXEM(symbol string) string {
	if strings.Contains(symbol, "nem") {
		return strings.Replace(symbol, "nem", "xem", -1)
	}
	if strings.Contains(symbol, "NEM") {
		return strings.Replace(symbol, "nem", "xem", -1)
	}
	return symbol
}

func XEMtoNEM(symbol string) string {
	if strings.Contains(symbol, "xemx") || strings.Contains(symbol, "XEMX") {
		return symbol
	}
	if strings.Contains(symbol, "xem") {
		return strings.Replace(symbol, "xem", "nem", -1)
	}
	if strings.Contains(symbol, "XEM") {
		return strings.Replace(symbol, "XEM", "NEM", -1)
	}
	return symbol
}
