package main

import (
	"mg"
	"net/http"
)

func main() {

	// 创建myGin
	r := mg.New()
	r.GET("/hello", func(c *mg.Context) {
		name := c.Query("name")
		c.String(http.StatusOK, "hello %v", name)
	})

	r.POST("/hello", func(c *mg.Context) {
		name := c.PostForm("name")
		age := c.PostForm("age")
		c.JSON(http.StatusOK, mg.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/user/:id", func(c *mg.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, mg.H{
			"id":   id,
			"path": c.Path,
		})
	})

	r.POST("/upload/*filePath", func(c *mg.Context) {
		filePath := c.Param("filePath")
		c.JSON(http.StatusOK, mg.H{
			"filePath": filePath,
			"path":     c.Path,
		})
	})

	// 启动
	r.Run(":8081")
}
