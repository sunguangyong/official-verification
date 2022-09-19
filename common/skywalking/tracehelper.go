package skywalking

import (
	"sync"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
)

var tracer *go2sky.Tracer
var gcm GoroutineContextManager
var once sync.Once

func StartTracer(serviceAddr string, serviceName string) error {
	rp, err := reporter.NewGRPCReporter(serviceAddr, reporter.WithCheckInterval(time.Second))
	if err != nil {
		return err
	}
	tracer, err = go2sky.NewTracer(serviceName, go2sky.WithReporter(rp))

	once.Do(func() {
		gcm = GoroutineContextManager{}
	})

	return nil
}

func GetTracer() *go2sky.Tracer {
	return tracer
}

func GetGcm() *GoroutineContextManager {
	return &gcm
}
