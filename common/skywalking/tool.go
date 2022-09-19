package skywalking

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/SkyAPM/go2sky"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

const ComponentIDGINHttpServer = 5006

func Get(link string) (response string, err error) {
	client := http.Client{Timeout: time.Second * 10}
	var reqest *http.Request
	reqest, err = http.NewRequest("GET", link, nil)
	if err != nil {
		return
	}

	tracer = GetTracer()
	// perr 的作用就是Tag信息
	// operationName 就是调用名称，意思要明确
	url := reqest.URL
	operationName := url.Scheme + "://" + url.Host + url.Path
	ctx, _ := GetGcm().GetContext()
	span, err := tracer.CreateExitSpan(*ctx, operationName, url.Host, func(key, value string) error {
		reqest.Header.Set(key, value)
		return nil
	})
	if err != nil {
		return
	}
	span.SetComponent(ComponentIDGINHttpServer)
	span.Tag(go2sky.TagHTTPMethod, reqest.Method)
	span.Tag(go2sky.TagURL, link)
	span.SetSpanLayer(agentv3.SpanLayer_Http)

	resp, err := client.Do(reqest)
	if err != nil {
		span.Error(time.Now(), err.Error())
	} else {
		span.Tag(go2sky.TagStatusCode, strconv.Itoa(resp.StatusCode))
	}

	span.End()

	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

type GoroutineContextManager struct {
	gls GoroutineLocalStorage
}

func (gcm *GoroutineContextManager) SetContext(ctx *context.Context) {
	key := gcm.gls.GetGoroutineId()
	gcm.gls.Set(key, ctx)
}

func (gcm *GoroutineContextManager) DelContext() {
	key := gcm.gls.GetGoroutineId()
	gcm.gls.Del(key)
}

func (gcm *GoroutineContextManager) GetContext() (*context.Context, bool) {
	key := gcm.gls.GetGoroutineId()
	ctx, ok := gcm.gls.Get(key).(*context.Context)
	return ctx, ok
}

type GoroutineLocalStorage struct {
	sync.Map
}

func (gls *GoroutineLocalStorage) Get(key uint64) interface{} {
	value, ok := gls.Load(key)
	if !ok {
		return nil
	}
	return value
}

func (gls *GoroutineLocalStorage) Set(key uint64, v interface{}) {
	gls.Store(key, v)
}

func (gls *GoroutineLocalStorage) Del(key uint64) {
	gls.Delete(key)
}

func (gls *GoroutineLocalStorage) GetGoroutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
