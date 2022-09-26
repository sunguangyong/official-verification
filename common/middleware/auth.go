package middleware

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func SortMap(m map[string]interface{}) string {
	keys := make([]string, 0, len(m))
	str := ""
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var builder strings.Builder
	for _, k := range keys {
		builder.WriteString(fmt.Sprintf("%s=%v&", k, m[k]))
	}
	str = builder.String()
	str = strings.Trim(str, "")
	return str
}

type UserApiKey struct {
	ApiKey      string    `db:"api-key"`      // api key
	ApiSecret   string    `db:"api_secret"`   // api secret
	TradeStatus int64     `db:"trade_status"` // 是否可以交易，0不可以，1可以
	Deleted     int64     `db:"deleted"`      // 是否已删除
	ApiLevel    int64     `db:"api_level"`    // 客户级别
	IsBindIp    int64     `db:"is_bind_ip"`   // 是否绑定ip，0未绑定，1绑定
	Ctime       time.Time `db:"ctime"`        // 第一次创建时间
	Mtime       time.Time `db:"mtime"`        // 重置后，新TOKEN生成时间
}
