# MyGin

> https://geektutu.com/post/gee.html

对Web服务来说，无非是根据请求*http.Request，构造响应http.ResponseWriter, 当我们使用基础库时，需要很多重复且毫无意义的频繁手工处理的地方
这就是框架的价值所在, 能够方便的获取请求各种信息, 并且方便高效的构建响应, 提供常用工具对http中的header, cookies等进行方便处理,
并且预留扩展点, 让使用者可以对请求响应进行增强

## 目录结构
```
mg/
    |--mg.go
    |--go.mod
    |--Context.go
    |--router.go
    |--RouterGroup.go
    |--trie.go
    |--Logger.go
    |--Recovery.go
main.go
go.mod
```

## 职责划分
- mg.go

请求入口, 框架引擎. 实现了http包的 `ServeHTTP(w http.ResponseWriter, r *http.Request)` 接口, 统一接管所有请求

- Context

定义一次请求的上下文, 路由的处理函数，中间件，参数等在请求中常用的属性全都保存在里面

- tire

使用前缀树结构保存路由, 能够匹配动态路由

- router

路由, 控制路由和路由对应函数的执行

- RouterGroup

分组控制

- Logger & Recovery

默认实现的日志和错误处理机制的中间件


