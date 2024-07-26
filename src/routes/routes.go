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
		api.GET("/protected", m.RequireAuth(), c.Ping)
		user := api.Group("/user")
		{
			user.POST("/add", gin.Bind(models.CreateUserBody{}), c.CreateUser)
			user.PUT("/update", gin.Bind(models.UpdateUserBody{}), c.UpdateUser)
			user.GET("/get/:id", c.GetUser)
			user.DELETE("/delete/:id", m.RequireAuth(), m.RequireAdmin(), c.DeleteUser)
			user.GET("/list", c.ListUsers)
			user.POST("/login", gin.Bind(models.LoginUserBody{}), c.Login)
		}
	}
	return r
}
