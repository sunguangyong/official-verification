package str

import (
	"strconv"
	"strings"
)

func SafeSql(str string) string {
	str = strings.ReplaceAll(str, "=", "")
	str = strings.ReplaceAll(str, "&", "")
	str = strings.ReplaceAll(str, "(", "")
	str = strings.ReplaceAll(str, ")", "")
	str = strings.ReplaceAll(str, ";", "")
	str = strings.Trim(str, " ")
	str = strings.Trim(str, "'")
	return str
}

func StrIn(str string, strs []string) bool {
	for _, v := range strs {
		if strings.Trim(strings.ToLower(str), "") == strings.Trim(strings.ToLower(v), "") {
			return true
		}
	}
	return false
}

func ConvertStrToFloat64(str string) float64 {
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return -1
	}
	return v
}
