package mg

import (
	"net/http"
	"strings"
)

// router 路由结构体
type router struct {
	// roots 保存每种请求方式的根节点 key = 请求方式, value = 根节点指针
	roots map[string]*node
	// handlers 路径和要执行的handler映射
	handlers map[string]HandlerFunc
}

// 提供初始化路由的函数
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// parsePattern 解析路由
func parsePattern(pattern string) []string {
	patternSplit := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, part := range patternSplit {
		if part != "" {
			parts = append(parts, part)
			if strings.HasPrefix(part, "*") {
				break
			}
		}
	}
	return parts
}

// addRouter 增加路由
func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	_, ok := r.roots[method]
	if !ok {
		// 该请求方法没有根节点, 需新建出根节点
		r.roots[method] = &node{}
	}

	// 开始新建trie节点
	r.roots[method].insert(pattern, parsePattern(pattern), 0)

	// 用自己的格式构建key
	key := strings.Join([]string{method, pattern}, "-")
	// 记录key => func 映射
	r.handlers[key] = handler
}

// getRoute 获取路由
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	root, ok := r.roots[method]
	if !ok {
		// 没找到路由 404
		return nil, nil
	}
	params := make(map[string]string)
	// 查询路由
	n := root.search(parsePattern(path), 0)
	if n == nil {
		return nil, nil
	}
	// 解析参数
	parts := parsePattern(n.pattern)
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			params[part[1:]] = parsePattern(path)[i]
		}
		if strings.HasPrefix(part, "*") {
			params[part[1:]] = strings.Join(parsePattern(path)[i:], "/")
			break
		}
	}
	return n, params
}

// handle 处理路径和handler
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
