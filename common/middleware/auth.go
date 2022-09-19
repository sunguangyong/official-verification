package middleware

import (
	"bytes"
	"cointiger.com/verification/common/constant"
	"cointiger.com/verification/common/instance"
	"cointiger.com/verification/common/req"
	"cointiger.com/verification/common/result"
	"cointiger.com/verification/common/sign"
	"cointiger.com/verification/common/xerr"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
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

			//sign
			signParam := req.NewStr(constant.SignParamName, r.Header.Get(constant.SignParamName))
			if signParam.Len(1, 1024*1024).Empty().Err != nil {
				result.HttpResult(r, w, nil, xerr.NewErr(constant.REQUEST_PARAM_ERROR, signParam.Err))
				return
			}

			// check apikey in redis
			rl := instance.GetRedis()
			uidStr, err := rl.GetCtx(context.Background(), instance.GetRedisApiKey(apiKey.GetValue()))
			uidStr = strings.ReplaceAll(uidStr, "\"", "")
			if err != nil || uidStr == "" {
				result.HttpResult(r, w, nil, xerr.NewNotFountErr(constant.REQUEST_PARAM_ERROR, constant.ApiKeyParamName, apiKey.GetValue()))
				return
			}
			r.Header.Add(constant.HttpHeaderUserId, uidStr)

			isExist, err := rl.SismemberCtx(context.Background(), instance.GetRedisBlacklistUserKey(), uidStr)
			if err != nil || isExist {
				result.HttpResult(r, w, nil, xerr.NewNoPermission())
				return
			}

			// read body
			var bodyBytes []byte
			bodyBytes, _ = ioutil.ReadAll(r.Body)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			m := make(map[string]interface{})
			err = json.Unmarshal(bodyBytes, &m)
			if err != nil {
				result.HttpResult(r, w, nil, xerr.NewErrCode(constant.SERVER_COMMON_ERROR))
				return
			}

			secretStr, err := rl.GetCtx(context.Background(), instance.GetRedisUserApiKeyEntity(apiKey.GetValue()))
			if err != nil || secretStr == "" {
				result.HttpResult(r, w, nil, xerr.NewNotFountErr(constant.REQUEST_PARAM_ERROR, constant.SecretName, ""))
				return
			}
			v := new(UserApiKey)
			if err := json.Unmarshal([]byte(secretStr), v); err != nil {
				result.HttpResult(r, w, nil, xerr.NewNotFountErr(constant.REQUEST_PARAM_ERROR, constant.SecretName, ""))
				return
			}
			//assembly parameters
			//assemblyParamStr := SortMap(m)

			assemblyParamStr := string(bodyBytes) //修改签名直接签署body
			assemblyParamStr += v.ApiSecret

			signAfterStr := sign.Sign(apiKey.GetValue(), assemblyParamStr)
			fmt.Println("auth:apikey:", apiKey.GetValue())
			fmt.Println("auth:waitSign:", assemblyParamStr)
			fmt.Println("auth:signstr:", signAfterStr)
			if signAfterStr != signParam.GetValue() {
				w.WriteHeader(http.StatusUnauthorized)
				result.HttpResult(r, w, nil, xerr.NewErrCode(constant.REQUEST_SIGN_ERROR))
				return
			}
		}
		next(w, r)
	}
}

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
