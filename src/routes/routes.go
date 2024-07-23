package routes

import (
	"html-aiccesible/controllers"

	"github.com/gin-gonic/gin"
	v "github.com/go-playground/validator/v10"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	validate := v.New()
	c := controllers.NewController(validate)
	r.GET("/", c.Ping)
	return r
}
