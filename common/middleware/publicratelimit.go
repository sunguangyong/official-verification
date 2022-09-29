package middleware

import (
	"cointiger.com/verification/common/constant"
	"cointiger.com/verification/common/instance"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"cointiger.com/verification/common/req"
	"cointiger.com/verification/common/result"
	"cointiger.com/verification/common/xerr"
	"github.com/zeromicro/go-zero/core/limit"

	"errors"
	"net"
	"net/http"
	"strings"
)

type RateLimt struct {
	Seconds   int  // 秒
	Quota     int  // 定额
	RateType  int  //1:private,2:public
	SkipLimit bool //是否跳过限流
}

var apiRateLimit = map[string]*RateLimt{
	// 60 秒 放行一个
	"/api/offical/verify/informadd": &RateLimt{Seconds: 60, Quota: 1, RateType: constant.ApiPublicRateLimit, SkipLimit: false},
}

func PublicRateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := req.NewStr(constant.PathParamName, r.RequestURI)
		if path.Len(1, 200).Empty().Err != nil {
			result.HttpResult(r, w, nil, nil, xerr.NewErr(constant.REQUEST_PARAM_ERROR, path.Err))
			return
		}
		v, isOk := apiRateLimit[path.GetValue()]
		if isOk && v.RateType == constant.ApiPublicRateLimit && !v.SkipLimit {
			rl := instance.GetRedis()
			rate, isOK := apiRateLimit[path.GetValue()]
			fmt.Println(rate)
			if !isOK {
				result.HttpResult(r, w, nil, nil, xerr.NewParamsErr(constant.ApiKeyParamName, path.GetValue()))
				return
			}

			ip, _ := GetIP(r)
			rlKey := instance.GetRedisPubRateLimitKey(path.GetValue(), ip)
			logx.Infof("rlKey ====== ", rlKey)

			l := limit.NewPeriodLimit(rate.Seconds, rate.Quota, rl, rlKey)
			code, err := l.Take("")
			if err != nil {
				logx.Error(err)
			}

			switch code {
			case limit.OverQuota: // 超过限制
				logx.Errorf("OverQuota key: %v", rlKey)
				result.HttpResult(r, w, nil, nil, xerr.NewRateLimitErr())
				return

			case limit.Allowed: // 通过
				logx.Infof("AllowedQuota key: %v", rlKey)

			case limit.HitQuota: // 达到限额
				logx.Errorf("HitQuota key: %v", rlKey)
			default:
				logx.Errorf("DefaultQuota key: %v", rlKey)
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
