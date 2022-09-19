package middleware

import (
	"cointiger.com/verification/common/constant"
	"cointiger.com/verification/common/instance"
	"cointiger.com/verification/common/req"
	"cointiger.com/verification/common/result"
	"cointiger.com/verification/common/xerr"
	"errors"
	"github.com/zeromicro/go-zero/core/limit"
	"net"
	"net/http"
	"strings"
)

func PublicRateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//path
		path := req.NewStr(constant.PathParamName, r.RequestURI)
		if path.Len(1, 200).Empty().Err != nil {
			result.HttpResult(r, w, nil, xerr.NewErr(constant.REQUEST_PARAM_ERROR, path.Err))
			return
		}
		v, isOk := apiRateLimit[path.GetValue()]
		if !isOk {
			result.HttpResult(r, w, nil, xerr.NewErr(constant.REQUEST_PARAM_ERROR, path.Err))
			return
		}
		if isOk && v.RateType == constant.ApiPublicRateLimit && !v.SkipLimit {
			rl := instance.GetRedis()
			rate, isOK := apiRateLimit[path.GetValue()]
			if !isOK {
				result.HttpResult(r, w, nil, xerr.NewParamsErr(constant.ApiKeyParamName, path.GetValue()))
				return
			}

			ip, _ := GetIP(r)
			rlKey := instance.GetRedisPubRateLimitKey(path.GetValue(), ip)
			limiter := limit.NewTokenLimiter(rate.QPS, rate.QPS, rl, rlKey)
			if !limiter.Allow() {
				result.HttpResult(r, w, nil, xerr.NewRateLimitErr())
				return
			}
		}
		next(w, r)
	}
}
func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}
	return "0.0.0.0", errors.New("no valid ip found")
}
