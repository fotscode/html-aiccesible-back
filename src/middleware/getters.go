package middleware

import (
	"html-aiccesible/httputil"
	"html-aiccesible/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListOptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil {
			httputil.BadRequest[string](c, "Invalid page")
			return
		}
		size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
		if err != nil {
			httputil.BadRequest[string](c, "Invalid size")
			return
		}
		lo := &models.ListOptions{
			Page: page,
			Size: size,
		}
		c.Set("lo", lo)
	}
}

func GetOptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			httputil.BadRequest[string](c, "Invalid ID")
			return
		}
		getOpt := &models.GetOptions{
			Id: id,
		}
		c.Set("getOpt", getOpt)
	}
}
