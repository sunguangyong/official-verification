package constant

const (
	OK                       uint32 = 200    //成功返回
	SERVER_COMMON_ERROR      uint32 = 100001 //全局通用错误
	REQUEST_PARAM_ERROR      uint32 = 100002 //请求参数错误
	REQUEST_RATE_LIMIT_ERROR uint32 = 100003 //触发限流
	REQUEST_SIGN_ERROR       uint32 = 100004 //签名错误
	REQUEST_NO_PERMISSION    uint32 = 100005 //没有权限
	REQUEST_SYMBOL_ERROR     uint32 = 100006 //无效symbol
	REQUEST_DATADUPLICATION  uint32 = 100007 //数据重复
)
