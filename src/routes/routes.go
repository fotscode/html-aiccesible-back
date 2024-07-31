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
		user := api.Group("/user")
		{
			user.POST("/add", gin.Bind(models.CreateUserBody{}), c.CreateUser)
			user.PUT("/update", m.Auth(c), gin.Bind(models.UpdateUserBody{}), c.UpdateUser)
			user.GET("/get/:id", m.GetOptions(), c.GetUser)
			user.DELETE("/delete/:id", m.GetOptions(), m.Auth(c), m.Admin(), c.DeleteUser)
			user.GET("/list", m.ListOptions(), c.ListUsers)
			user.POST("/login", gin.Bind(models.LoginUserBody{}), c.Login)
		}
		config := api.Group("/config", m.Auth(c))
		{
			config.GET("/get", c.GetConfig)
			config.PUT("/update", gin.Bind(models.UpdateConfigBody{}), c.UpdateConfig)
		}
		post := api.Group("/post")
		{
			post.POST("/add", m.Auth(c), gin.Bind(models.CreatePostBody{}), c.CreatePost)
			post.PUT("/update", m.Auth(c), gin.Bind(models.UpdatePostBody{}), c.UpdatePost)
			post.GET("/get/:id", m.GetOptions(), c.GetPost)
			post.DELETE("/delete/:id", m.Auth(c), m.GetOptions(), c.DeletePost)
			post.GET("/list", m.ListOptions(), c.ListPosts)
			post.PATCH("/like/:id", m.Auth(c), m.GetOptions(), c.LikePost)
			post.GET("/likes/:id", m.GetOptions(), c.GetPostLikes)
		}
		comment := api.Group("/comment", m.Auth(c))
		{
			comment.POST("/add", m.Auth(c), gin.Bind(models.CreateCommentBody{}), c.CreateComment)
			comment.PUT("/update", m.Auth(c), gin.Bind(models.UpdateCommentBody{}), c.UpdateComment)
			comment.GET("/get/:id", m.GetOptions(), c.GetComment)
			comment.DELETE("/delete/:id", m.Auth(c), m.GetOptions(), c.DeleteComment)
			comment.GET("/list/:id", m.GetOptions(), m.ListOptions(), c.ListComments)
		}
		AI := api.Group("/models")
		{
			AI.GET("/list", c.ListModels)
			AI.POST("/accesibilize", gin.Bind(models.AccesibilizeBody{}), c.Accesibilize)
		}
	}
	return r
}
