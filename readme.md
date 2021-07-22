# tao

[English]() | 简体中文


## 0. tao 介绍

tao 是一个集成了各种工程实践的 web 和 rpc 框架。通过弹性设计保障了大并发服务端的稳定性，经受了充分的实战检验。

tao 包含极简的 API 定义和生成工具 taoctl，可以根据定义的 api 文件一键生成 Go, iOS, Android, Kotlin, Dart, TypeScript, JavaScript 代码，并可直接运行。

使用 tao 的好处：

* 轻松获得支撑千万日活服务的稳定性
* 内建级联超时控制、限流、自适应熔断、自适应降载等微服务治理能力，无需配置和额外代码
* 微服务治理中间件可无缝集成到其它现有框架使用
* 极简的 API 描述，一键生成各端代码
* 自动校验客户端请求参数合法性
* 大量微服务治理和并发工具包

<img src="https://gitee.com/kevwan/static/raw/master/doc/images/architecture.png" alt="架构图" width="1500" />


## 1. tao 框架设计

对于微服务框架的设计，我们期望保障微服务稳定性的同时，也要特别注重研发效率。所以设计之初，我们就有如下一些准则：

* 保持简单，第一原则
* 弹性设计，面向故障编程
* 工具大于约定和文档
* 高可用
* 高并发
* 易扩展
* 对业务开发友好，封装复杂度
* 约束做一件事只有一种方式


## 2. tao 项目实现和特点

tao 是一个集成了各种工程实践的包含 web 和 rpc 框架，有如下主要特点：

* 强大的工具支持，尽可能少的代码编写
* 极简的接口
* 完全兼容 net/http
* 支持中间件，方便扩展
* 高性能
* 面向故障编程，弹性设计
* 内建服务发现、负载均衡
* 内建限流、熔断、降载，且自动触发，自动恢复
* API 参数自动校验
* 超时级联控制
* 自动缓存控制
* 链路跟踪、统计报警等
* 高并发支撑，稳定保障了疫情期间每天的流量洪峰

如下图，我们从多个层面保障了整体服务的高可用：

<<<<<<< HEAD
![弹性设计](https://gitee.com/kevwan/static/raw/master/doc/images/resilience.jpg)
=======
go-zero is a web and rpc framework that integrates lots of engineering practices. The features are mainly listed below:

* powerful tool included, less code to write
* simple interfaces
* fully compatible with net/http
* middlewares are supported, easy to extend
* high performance
* failure-oriented programming, resilience design
* builtin service discovery, load balancing
* builtin concurrency control, adaptive circuit breaker, adaptive load shedding, auto trigger, auto recover
* auto validation of API request parameters
* chained timeout control
* auto management of data caching
* call tracing, metrics and monitoring
* high concurrency protected

As below, go-zero protects the system with couple layers and mechanisms:

![Resilience](https://raw.githubusercontent.com/tal-tech/zero-doc/main/doc/images/resilience-en.png)

## 4. Future development plans of go-zero

* auto generate API mock server, make the client debugging easier
* auto generate the simple integration test for the server side just from the .api files

## 5. Installation

Run the following command under your project:

```shell
go get -u github.com/tal-tech/go-zero
```

## 6. Quick Start

0. full examples can be checked out from below:

     [Rapid development of microservice systems](https://github.com/tal-tech/zero-doc/blob/main/doc/shorturl-en.md)

     [Rapid development of microservice systems - multiple RPCs](https://github.com/tal-tech/zero-doc/blob/main/docs/zero/bookstore-en.md)

1. install goctl

   `goctl`can be read as `go control`. `goctl` means not to be controlled by code, instead, we control it. The inside `go` is not `golang`. At the very beginning, I was expecting it to help us improve the productivity, and make our lives easier.

   ```shell
   GO111MODULE=on go get -u github.com/tal-tech/go-zero/tools/goctl
   ```

   make sure goctl is executable.

2. create the API file, like greet.api, you can install the plugin of goctl in vs code, api syntax is supported.

   ```go
   type Request struct {
     Name string `path:"name,options=you|me"` // parameters are auto validated
   }
   
   type Response struct {
     Message string `json:"message"`
   }
   
   service greet-api {
     @handler GreetHandler
     get /greet/from/:name(Request) returns (Response);
   }
   ```
   
   the .api files also can be generate by goctl, like below:

   ```shell
   goctl api -o greet.api
   ```
   
3. generate the go server side code

   ```shell
   goctl api go -api greet.api -dir greet
   ```

   the generated files look like:

   ```Plain Text
   ├── greet
   │   ├── etc
   │   │   └── greet-api.yaml        // configuration file
   │   ├── greet.go                  // main file
   │   └── internal
   │       ├── config
   │       │   └── config.go         // configuration definition
   │       ├── handler
   │       │   ├── greethandler.go   // get/put/post/delete routes are defined here
   │       │   └── routes.go         // routes list
   │       ├── logic
   │       │   └── greetlogic.go     // request logic can be written here
   │       ├── svc
   │       │   └── servicecontext.go // service context, mysql/redis can be passed in here
   │       └── types
   │           └── types.go          // request/response defined here
   └── greet.api                     // api description file
   ```

   the generated code can be run directly:

   ```shell
   cd greet
   go mod init
   go mod tidy
   go run greet.go -f etc/greet-api.yaml
   ```

   by default, it’s listening on port 8888, while it can be changed in configuration file.

   you can check it by curl:

   ```shell
   curl -i http://localhost:8888/greet/from/you
   ```

   the response looks like:

   ```http
   HTTP/1.1 200 OK
   Date: Sun, 30 Aug 2020 15:32:35 GMT
   Content-Length: 0
   ```

4. Write the business logic code

    * the dependencies can be passed into the logic within servicecontext.go, like mysql, reds etc.
    * add the logic code in logic package according to .api file

5. Generate code like Java, TypeScript, Dart, JavaScript etc. just from the api file

   ```shell
   goctl api java -api greet.api -dir greet
   goctl api dart -api greet.api -dir greet
   ...
   ```

## 7. Benchmark

![benchmark](https://raw.githubusercontent.com/tal-tech/zero-doc/main/doc/images/benchmark.png)

[Checkout the test code](https://github.com/smallnest/go-web-framework-benchmark)

## 8. Documents (adding)

* [Documents](https://go-zero.dev/en/)
* [Rapid development of microservice systems](https://github.com/tal-tech/zero-doc/blob/main/doc/shorturl-en.md)
* [Rapid development of microservice systems - multiple RPCs](https://github.com/tal-tech/zero-doc/blob/main/docs/zero/bookstore-en.md)
* [Examples](https://github.com/zeromicro/zero-examples)

## 9. Chat group

Join the chat via https://join.slack.com/t/go-zero/shared_invite/zt-thyennhc-_fNXFpeUJcGE_tQNZFpsdA

## Give a Star! ⭐

If you like or are using this project to learn or start your solution, please give it a star. Thanks!
>>>>>>> 75952308f9223d64ff31d3013263b97a1f96a3d6
