# 通用能力服务

接口文档：https://yapi.dustess.com/project/2476/interface/api

## 文档生成

```shell
make doc
```

## grpc 生成go文档

```shell
make grpc-proto
```

## grpc调试工具

```shell
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest

grpcui -plaintext localhost:50051

```

## rpc header 含义
```
Authorization = "authorization"

RequestID     = "x-request-id"
AuthAccountID = "x-auth-accountid"

B3TraceID      = "x-b3-traceid"
B3SpanID       = "x-b3-spanid"
B3ParentSpanID = "x-b3-parentspanid"
B3Sampled      = "x-b3-sampled"
B3Flags        = "x-b3-flags"

BizCtxAuthorization = "x-bizctx-authorization" // 前端传递的身份token，现阶段为了某些场景能兼容先透传
BizCtxCaller        = "x-bizctx-caller"        // 调用方法的应用名
BizCtxAccount       = "x-bizctx-account"       // 租户ID
BizCtxUID           = "x-bizctx-uid"           // 员工ID
BizCtxOpenID        = "x-bizctx-openid"        // 微信OpenID
BizCtxUnionID       = "x-bizctx-unionid"       // 微信UnionID
BizCtxCDB           = "x-bizctx-cdb"           // MongoDB库名（未来会废弃掉）
```

## 工程结构

```aidl
application 
|- service     // 应用服务
cmd         
|- server      // 服务启动入口
common
|- dto         // 应用服务、领域服务的入参定义
|- errors      // 业务错误的预定义
|- vo          // 应用服务、领域服务的返回值定义
|- utils       // 基础设施无关的工具类
domain
|- entity      // 领域模型（聚合根、实体、值对象）、防腐层模型、读模型
|- event       // 领域事件
|- interfaces  // 对基础实施依赖的接口约定
|- repository  // 对领域模型持久化的接口约定
|- service     // 领域服务（与应用服务的区别是领域服务具有业务不变性，如果识别不了可以都写到应用服务中）
infrastructure
|- config      // 配置的数据结构与加载
|- controller  // 请求入口（北向网关）
|- http        // HTTP路由与Controller的绑定（北向网关）
|- grpc        // GRPC接口与Controller的绑定（北向网关）
|- persistence // 领域模型持久化的实现，实现了领域模型到存储模型的转换与落库（南向网关）
|- service     // 对应用层、领域层的定义接口的实现，比如RPC请求
```
