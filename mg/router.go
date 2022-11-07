package mg

import (
	"log"
	"net/http"
	"strings"
)

// router 路由结构体
type router struct {
	// handlers 路径和要执行的handler映射
	handlers map[string]HandlerFunc
}

// 提供初始化路由的函数
func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

// addRouter 增加路由
func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	// 用自己的格式构建key
	key := strings.Join([]string{method, pattern}, "-")
	r.handlers[key] = handler
}

// handle 处理路径和handler
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
