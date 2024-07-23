package controllers

import (
	"html-aiccesible/httputil"

	"github.com/gin-gonic/gin"
)

func (b *Controller) Ping(c *gin.Context) {
	httputil.OK(c, "pong")
}
