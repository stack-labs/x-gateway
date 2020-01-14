package metrics

import (
	"github.com/micro-in-cn/x-gateway/plugin"
)

//NewPlugin of metrics
func NewPlugin(opts ...Option) plugin.Plugin {
	return newPrometheus(opts...)
}
