<img align="right" width="150px" src="https://gitee.com/kevwan/static/raw/master/doc/images/go-zero.png">

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

![弹性设计](https://gitee.com/kevwan/static/raw/master/doc/images/resilience.jpg)
