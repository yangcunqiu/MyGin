package main

import (
	"fmt"
	"log"
	"mg"
	"net/http"
	"time"
)

func main() {

	// 创建myGin
	r := mg.Default()
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

	bookGroup := r.Group("/book")
	bookGroup.Use(onlyBook())
	{
		bookGroup.GET("/info", func(c *mg.Context) {
			log.Printf("/book/info---invoke")
			c.String(http.StatusOK, fmt.Sprintf("path:%v", c.Path))
		})
	}

	r.GET("/panic", func(c *mg.Context) {
		a := []int{1, 2}
		b := a[4:]
		fmt.Println(b)
	})

	// 启动
	r.Run(":8081")
}

func calCost() mg.HandlerFunc {
	return func(c *mg.Context) {
		start := time.Now()
		c.Next()
		log.Printf("%v cost %v ns", c.Path, time.Since(start))
	}
}

func onlyBook() mg.HandlerFunc {
	return func(c *mg.Context) {
		log.Printf("onlyBookMiddleware---start")
		c.Next()
		log.Printf("onlyBookMiddleware---end")
	}
}
