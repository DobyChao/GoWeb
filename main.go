package main

import (
	"goo"
	"net/http"
)

func main() {
	r := goo.New()
	r.GET("/", func(c *goo.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Goo</h1>")
	})
	r.GET("/hello", func(c *goo.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *goo.Context) {
		c.JSON(http.StatusOK, goo.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.GET("/hello/:name", func(c *goo.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	
	r.GET("/assets/*filepath", func(c *goo.Context) {
		c.JSON(http.StatusOK, goo.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}