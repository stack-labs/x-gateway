# API-Gateway

基于Go-micro的高性能网关，增加自定义micro工具，如`Auth`、`CORS`等.

[API​​网关](http://microservices.io/patterns/apigateway.html)为服务提供一个统一的公共流量入口。

![MICRO-API](docs/micro-api.png)

## API-Gateway的应用场景

+ 流量入口，统一用户访问的全部或者部分流量入口，包括移动端、PC Web端.
+ 允许不同系统的用户通过相同的网关访问后端的不同系统.
+ 对用户的请求进行统一认证(同时依赖后端认证服务).
+ 对用户操作进行鉴权(权限校验),确定用户是否具有具有操作或者访问资源的权限(同时依赖后端鉴权服务).
+ 记录用户的操作行为.
+ 追踪用户的操作行为链.

## 功能设计

+ 认证&鉴权`JWT`+`Casbin` [Auth](/pkg/plugin/micro/auth)
+ 跨域支持 [CORS](/pkg/plugin/micro/cors)
+ Metrics [Prometheus](/pkg/plugin/micro/metrics)
+ Trace [Opentracing](/pkg/plugin/micro/trace/opentracing)
+ REST to GRPC 转换REST调用到GRPC(HTTP[s] -> API-Gateway -> Micro Srv)[计划中...]
+ 提供服务发现
+ 动态路由公共
- 高性能
- 智能路由
- 流量控制管理
- 日志定制
- 链路追踪
- 认证
- 版本化
- 灰度，AB策略
- ...

# TODO

## 运行网关

```bash
# 编译
$ make build

# API
$ make run_api                                  # 默认mdns + http
$ make run_api registry=etcd transport=tcp      # 指定etcd + tcp

# Web
$ make run_web                                  # 默认mdns + http
$ make run_web registry=etcd transport=tcp      # 指定etcd + tcp
```

## Docker

```bash
# tag自定义
$ make docker tag=docker.pkg.github.com/micro-in-cn/api-gateway/api-gateway:v1.15.0
```
