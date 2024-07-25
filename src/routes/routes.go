package routes

import (
	"html-aiccesible/controllers"
	"html-aiccesible/middleware"
	"html-aiccesible/models"

	"github.com/gin-gonic/gin"
	v "github.com/go-playground/validator/v10"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	validate := v.New()
	c := controllers.NewController(validate)
	api := r.Group("/api", middleware.Errors(), middleware.JSONMiddleware())
	{
		api.GET("/", c.Ping)
		user := api.Group("/user")
		{
			user.POST("/add", gin.Bind(models.CreateUserBody{}), c.CreateUser)
			user.PUT("/update", gin.Bind(models.UpdateUserBody{}), c.UpdateUser)
			user.GET("/get/:id", c.GetUser)
			user.DELETE("/delete/:id", c.DeleteUser)
			user.GET("/list", c.ListUsers)
		}
	}
	return r
}
