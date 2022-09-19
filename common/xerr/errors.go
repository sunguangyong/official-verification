package xerr

import (
	"cointiger.com/verification/common/constant"
	"fmt"
)

/**
常用通用固定错误
*/

type CodeError struct {
	errCode uint32
	errMsg  string
}

//返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

//返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}
func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: constant.SERVER_COMMON_ERROR, errMsg: errMsg}
}

func NewErr(code uint32, err error) *CodeError {
	return &CodeError{errCode: code, errMsg: fmt.Sprintf("%v", err)}
}

func NewGetSymbolConfigErr() *CodeError {
	return &CodeError{errCode: constant.SERVER_COMMON_ERROR, errMsg: "get symbol config error"}
}

func NewNotFountErr(code uint32, name, value string) *CodeError {
	return &CodeError{errCode: code, errMsg: fmt.Sprintf("field:%s,value:%s,is not found", name, value)}
}

func NewRateLimitErr() *CodeError {
	return &CodeError{errCode: constant.REQUEST_RATE_LIMIT_ERROR, errMsg: " touch rate limit"}

}

func NewParamsErr(name, value string) *CodeError {
	return &CodeError{errCode: constant.REQUEST_PARAM_ERROR, errMsg: fmt.Sprintf("field:%s,value:%s error", name, value)}
}

func NewReqParamsErr() *CodeError {
	return &CodeError{errCode: constant.REQUEST_PARAM_ERROR, errMsg: "req error"}
}

func NewNoPermission() *CodeError {
	return &CodeError{errCode: constant.REQUEST_NO_PERMISSION, errMsg: "no permission"}
}

func NewInvaildSymbolErr(symbol string) *CodeError {
	return &CodeError{errCode: constant.REQUEST_SYMBOL_ERROR, errMsg: fmt.Sprintf("symbol:%s is invaild", symbol)}
}

func NewSymbolNotOpenErr(symbol string) *CodeError {
	return &CodeError{errCode: constant.REQUEST_SYMBOL_ERROR, errMsg: fmt.Sprintf("symbol:%s is not open", symbol)}
}

func NewLimitExceeded(field string) *CodeError {
	return &CodeError{errCode: constant.REQUEST_PARAM_ERROR, errMsg: fmt.Sprintf("field:%s, limit exceeded", field)}
}

func NewServerError() *CodeError {
	return &CodeError{errCode: constant.SERVER_COMMON_ERROR, errMsg: "server error"}
}
