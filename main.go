package main

import (
	"goo"
	"net/http"
)

func main() {
	r := goo.Default()
	r.GET("/", func(c *goo.Context) {
		c.String(http.StatusOK, "Hello goo\n")
	})
	r.GET("/panic", func(c *goo.Context) {
		names := []string{"goo"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
