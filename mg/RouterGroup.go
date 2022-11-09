package mg

// RouterGroup 路由组
type RouterGroup struct {
	prefix      string
	parent      *RouterGroup
	engine      *Engine
	middlewares []HandlerFunc
}
