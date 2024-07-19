package main

import (
	"html-aiccesible/httputil"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		httputil.OK(c, "Hello, World!")
	})
	r.Run()
}
