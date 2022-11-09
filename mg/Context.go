package mg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 为map[string]interface{}类型起了别名H
type H map[string]interface{}

// Context 创建上下文结构体, 封装整个请求过程中常用的属性
type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	Path       string
	Method     string
	StatusCode int
	Params     map[string]string
	// 中间件
	handlers []HandlerFunc
	// 记录当前执行到Context中的哪个中间件
	index int
}

// newContext 提供函数初始化一个 Context
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
		index:   -1,
	}
}

// Next 执行Context中的下一个中间件
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

// PostForm 表单获取参数
func (c *Context) PostForm(key string) string {
	return c.Request.PostFormValue(key)
}

// Query 路径获取?拼接参数
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// Status 设置响应码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置Header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 返回string
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 返回json
func (c *Context) JSON(code int, any interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(any); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 返回二进制文件
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "application/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) Fail(code int, errorMessage string) {
	c.String(code, "%v", errorMessage)
}
