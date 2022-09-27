package middleware

import (
	"cointiger.com/verification/common/constant"
	"cointiger.com/verification/common/instance"
	"cointiger.com/verification/common/req"
	"cointiger.com/verification/common/result"
	"cointiger.com/verification/common/xerr"
	"context"
	"github.com/zeromicro/go-zero/core/limit"
	"net/http"
)

type RateLimt struct {
	QPS       int  //qps/s
	RateType  int  //1:private,2:public
	SkipLimit bool //是否跳过限流
}

var apiRateLimit = map[string]*RateLimt{
	"/api/offical/verify/informadd": &RateLimt{QPS: 20, RateType: constant.ApiPublicRateLimit, SkipLimit: false},
}

func RateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//path
		path := req.NewStr(constant.PathParamName, r.RequestURI)
		if path.Len(1, 200).Empty().Err != nil {
			result.HttpResult(r, w, nil, xerr.NewErr(constant.REQUEST_PARAM_ERROR, path.Err))
			return
		}
		if v, isOk := apiRateLimit[path.GetValue()]; isOk && v.RateType == constant.ApiPrivateRateLimit && !v.SkipLimit {
			//apikey
			apiKey := req.NewStr(constant.ApiKeyParamName, r.Header.Get(constant.ApiKeyParamName))
			if apiKey.Len(1, 100).Empty().Err != nil {
				result.HttpResult(r, w, nil, xerr.NewErr(constant.REQUEST_PARAM_ERROR, apiKey.Err))
				return
			}
			rl := instance.GetRedis()
			isExist, err := rl.ExistsCtx(context.Background(), instance.GetRedisApiKey(apiKey.GetValue()))
			if err != nil || !isExist {
				result.HttpResult(r, w, nil, xerr.NewNotFountErr(constant.REQUEST_PARAM_ERROR, constant.ApiKeyParamName, apiKey.GetValue()))
				return
			}

			rate, isOK := apiRateLimit[path.GetValue()]
			if !isOK {
				result.HttpResult(r, w, nil, xerr.NewParamsErr(constant.ApiKeyParamName, path.GetValue()))
				return
			}

			//rate limit
			rlKey := instance.GetRedisPriRateLimitKey(path.GetValue(), apiKey.GetValue())
			limiter := limit.NewTokenLimiter(rate.QPS, rate.QPS, rl, rlKey)
			if !limiter.Allow() {
				result.HttpResult(r, w, nil, xerr.NewRateLimitErr())
				return
			}
		}
		next(w, r)
	}
}
