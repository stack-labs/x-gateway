package main

import (
	"io"
	"net/http"
	"time"

	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro-in-cn/x-gateway/api"
	"github.com/micro-in-cn/x-gateway/plugin/auth"
	"github.com/micro-in-cn/x-gateway/plugin/metrics"
	"github.com/micro-in-cn/x-gateway/plugin/opentracing"
	"github.com/micro-in-cn/x-gateway/utils/response"
	tracer "github.com/micro-in-cn/x-gateway/plugin/trace"
	"golang.org/x/time/rate"
)

var (
	apiTracerCloser io.Closer
)

func cleanWork() error {
	// closer
	apiTracerCloser.Close()

	return nil
}

// 插件注册
func init() {
	// Auth
	initAuth()
	initMetrics()
	initTrace()
}

func initAuth() {

	casb := fileadapter.NewAdapter("./conf/casbin_policy.csv")
	auth.RegisterAdapter("default", casb)

	authPlugin := auth.NewPlugin(
		auth.WithResponseHandler(response.DefaultResponseHandler),
		auth.WithSkipperFunc(func(r *http.Request) bool {
			return false
		}),
	)
	api.Register(authPlugin)

}

//(b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_' || b == ':' || (b >= '0' && b <= '9' && i > 0)
func initMetrics() {
	api.Register(metrics.NewPlugin(
		metrics.WithNamespace("xgateway"), //only [a-zA-Z0-9:_]
		metrics.WithSubsystem(""),
		metrics.WithSkipperFunc(func(r *http.Request) bool {
			return false
		}),
	))
}

// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
func initTrace() {
	apiTracer, apiCloser, err := tracer.NewJaegerTracer("go.micro.x-gateway", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	limiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)
	apiTracerCloser = apiCloser
	api.Register(opentracing.NewPlugin(
		opentracing.WithTracer(apiTracer),
		opentracing.WithSkipperFunc(func(r *http.Request) bool {
			// 采样频率控制，根据需要细分Host、Path等分别控制
			if !limiter.Allow() {
				return true
			}
			return false
		}),
	))

}
