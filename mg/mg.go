package mg

import (
	"net/http"
)

// HandlerFunc 自定义结构体, 供外界使用
type HandlerFunc func(*Context)

// Engine 引擎结构体
type Engine struct {
	// router 路由 路径和函数的映射关系. value是使用者的func
	router *router
	*RouterGroup
	groups []*RouterGroup
}

// ServerHTTP Engine实现http库中的ServeHTTP(ResponseWriter, *Request)方法, 这个方法会接管所有请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 构建上下文Context
	c := newContext(w, r)
	// 解析路由
	engine.router.handle(c)
}

// New 返回一个空的引擎
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// addRouter 添加路由 method 请求类型, pattern 路径, handler 要执行的函数
func (group *RouterGroup) addRouter(method string, pattern string, handler HandlerFunc) {
	group.engine.router.addRouter(method, group.prefix+pattern, handler)
}

// GET 绑定get方法
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRouter("GET", pattern, handler)
}

// POST 绑定post方法
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRouter("POST", pattern, handler)
}

// Run 启动
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
