# Micro Plugin

micro定制插件

- [auth](/plugin/auth)
- [cors](/pkg/plugin/micro/cors)
- [metrics](/pkg/plugin/micro/metrics)
- [trace](/pkg/plugin/micro/trace)
	- opentracing

## Ref

- [hb-go/micro-plugins](https://github.com/hb-go/micro-plugins)
	- 认证&鉴权`JWT`+`Casbin` [Auth](https://github.com/hb-go/micro-plugins/tree/master/micro/auth)
    - 跨域支持 [CORS](https://github.com/hb-go/micro-plugins/tree/master/micro/cors)    

### The plugin

Create a plugin.go file in the top level dir

```go
package main

import (
	"log"
	"github.com/micro/cli/v2"
	"github.com/micro/micro/v2/plugin"
)

func init() {
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("example"),
		plugin.WithFlag(&cli.StringFlag{
			Name:   "example_flag",
			Usage:  "This is an example plugin flag",
			EnvVars: []string{"EXAMPLE_FLAG"},
			Value: "avalue",
		}),
		plugin.WithInit(func(ctx *cli.Context) error {
			log.Println("Got value for example_flag", ctx.String("example_flag"))
			return nil
		}),
	))
}
```