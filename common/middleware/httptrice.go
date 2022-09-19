package middleware

import (
	"cointiger.com/verification/common/skywalking"
	"github.com/SkyAPM/go2sky"
	"net/http"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
	"time"
)

func HttpTrace(next http.HandlerFunc) http.HandlerFunc {
	tracerobj := skywalking.GetTracer()
	if tracerobj == nil {
		return func(w http.ResponseWriter, r *http.Request) {
			next(w, r)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		operationName := r.RequestURI
		span, ctx, err := tracerobj.CreateEntrySpan(r.Context(), operationName, func(key string) (string, error) {
			return r.Header.Get(key), nil
		})
		if err != nil {
			span.Error(time.Now(), "create entry span error:", err.Error())
			next(w, r)
			return
		}
		span.SetComponent(skywalking.ComponentIDGINHttpServer)
		span.Tag(go2sky.TagHTTPMethod, r.Method)
		span.Tag(go2sky.TagURL, r.Host+r.RequestURI)
		span.SetSpanLayer(agentv3.SpanLayer_Http)

		skywalking.GetGcm().SetContext(&ctx)
		defer skywalking.GetGcm().DelContext()

		next(w, r)

		span.Tag(go2sky.TagStatusCode, w.Header().Get("Status Code"))
		span.End()
	}
}
