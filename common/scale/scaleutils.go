package scale

import "fmt"

//目前策略法币保留两位小数，虚拟币保留八位
func ScaleBySymbol(v float64, symbol string) string {
	return fmt.Sprintf("%.8f", v)
}
