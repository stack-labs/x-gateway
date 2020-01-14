package auth

import (
	"github.com/micro-in-cn/x-gateway/internal/helper/request"
	"github.com/micro-in-cn/x-gateway/internal/helper/response"
)

//Options of auth
type Options struct {
	responseHandler response.Handler
	skipperFunc     request.SkipperFunc
}

//Option of auth
type Option func(o *Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		responseHandler: response.DefaultResponseHandler,
		skipperFunc:     request.DefaultSkipperFunc,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

//WithResponseHandler of auth
func WithResponseHandler(handler response.Handler) Option {
	return func(o *Options) {
		o.responseHandler = handler
	}
}

//WithSkipperFunc of auth
func WithSkipperFunc(skipperFunc request.SkipperFunc) Option {
	return func(o *Options) {
		o.skipperFunc = skipperFunc
	}
}
