module github.com/micro-in-cn/x-gateway

go 1.13

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d

require (
	github.com/casbin/casbin/v2 v2.1.1
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.16.1-0.20191121173113-8dc3fb964eaa
	github.com/micro/go-plugins v1.5.1
	github.com/micro/micro v1.16.1-0.20191121175420-186c72c1941d
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/client_golang v1.2.1
	github.com/rs/cors v1.7.0
	github.com/stretchr/testify v1.4.0
	github.com/uber/jaeger-client-go v2.20.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
)
