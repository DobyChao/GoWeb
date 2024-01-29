package main

import (
	"goo"
	"log"
	"net/http"
	"time"
)

func onlyForV2() goo.HandlerFunc {
	return func(c *goo.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := goo.New()
	// r.GET("/index", func(c *goo.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	// })

	// v1 := r.Group("/v1")
	// v1.GET("/", func(c *goo.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Hello Goo</h1>")
	// })
	// v1.GET("/hello", func(c *goo.Context) {
	// 	// expect /hello?name=goo
	// 	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	// })

	r.Use(goo.Logger()) // global midlleware
	r.GET("/", func(c *goo.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")

	// use middleware
	v2.Use(onlyForV2())

	v2.GET("/hello/:name", func(c *goo.Context) {
		// expect /hello/goo
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	// v2.POST("/login", func(c *goo.Context) {
	// 	c.JSON(http.StatusOK, goo.H{
	// 		"username": c.PostForm("username"),
	// 		"password": c.PostForm("password"),
	// 	})
	// })

	r.Run(":9999")
}
