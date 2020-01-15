package metrics

import (
	"github.com/micro/micro/plugin"
)

//NewPlugin of metrics
func NewPlugin(opts ...Option) plugin.Plugin {
	return newPrometheus(opts...)
}
