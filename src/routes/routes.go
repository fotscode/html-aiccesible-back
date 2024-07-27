package routes

import (
	"html-aiccesible/controllers"
	m "html-aiccesible/middleware"
	"html-aiccesible/models"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	c := controllers.NewController()
	api := r.Group("/api", m.Errors(), m.JSONMiddleware())
	{
		api.GET("/", c.Ping)
		api.GET("/protected", m.Auth(), c.Ping)
		user := api.Group("/user")
		{
			user.POST("/add", gin.Bind(models.CreateUserBody{}), c.CreateUser)
			user.PUT("/update", gin.Bind(models.UpdateUserBody{}), c.UpdateUser)
			user.GET("/get/:id", m.GetOptions(), c.GetUser)
			user.DELETE("/delete/:id", m.GetOptions(), m.Auth(), m.Admin(), c.DeleteUser)
			user.GET("/list", m.ListOptions(), c.ListUsers)
			user.POST("/login", gin.Bind(models.LoginUserBody{}), c.Login)
		}
		config := api.Group("/config", m.Auth())
		{
			config.GET("/get", c.GetConfig)
			config.PUT("/update", gin.Bind(models.UpdateConfigBody{}), c.UpdateConfig)
		}
	}
	return r
}
