package auth

import (
	"github.com/micro-in-cn/x-gateway/pkg/plugin/wrapper/util/request"
	"github.com/micro-in-cn/x-gateway/pkg/plugin/wrapper/util/response"
)

type Options struct {
	responseHandler response.Handler
	skipperFunc     request.SkipperFunc
}

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

func WithResponseHandler(handler response.Handler) Option {
	return func(o *Options) {
		o.responseHandler = handler
	}
}

func WithSkipperFunc(skipperFunc request.SkipperFunc) Option {
	return func(o *Options) {
		o.skipperFunc = skipperFunc
	}
}
