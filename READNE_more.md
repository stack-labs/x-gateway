# Overview

[micro api](https://micro.mu/docs/api.html)是micro中基于go-micro的API网关.

当前仓库的计划，是基于Micro精简到剩下Micro-API，再与x-gateway合并新的x-Gateway

![MICRO-API](https://github.com/micro-in-cn/x-gateway/raw/master/docs/micro-api.png)

## 注意事项

注意China go mod加速问题：[官方讨论](https://github.com/golang/go/issues/31755)

关联问题

[Go Modules Proxy](https://github.com/guanhui07/blog/issues/642)

[Go Module China 加速](https://github.com/developer-learning/night-reading-go/issues/468)

设置参考：

+ 设置China国内加)

```bash
go env -w GO111MODULE=on
//选一个代理
go env -w GOPROXY=https://goproxy.cn,direct   //go >= 1.13
go env -w GOPROXY=https://goproxy.io,direct
go env -w GOPROXY=https://proxy.golang.org,direct
go env -w GOPROXY=https://athens.azurefd.net,direct
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy,direct
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy,https://goproxy.cn,https://goproxy.io,https://athens.azurefd.net,direct
//选一个SUMDB
go env -w GOSUMDB=sum.golang.org //可选
go env -w GOSUMDB=sum.golang.google.cn //可选
```

## Usage

See all the options

```bash
micro --help
```
