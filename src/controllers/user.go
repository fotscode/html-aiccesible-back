package controllers

import (
	"html-aiccesible/httputil"
	"html-aiccesible/models"
	"strconv"

	"html-aiccesible/repositories"

	"github.com/gin-gonic/gin"
)

func (b *Controller) CreateUser(c *gin.Context) {
	body := c.MustGet(gin.BindKey).(*models.CreateUserBody)
	user, err := repositories.UserRepo().CreateUser(body)
	if err != nil {
		httputil.InternalServerError[string](c, err.Error())
		return
	}
	httputil.Created[*models.User](c, user)
}

func (b *Controller) UpdateUser(c *gin.Context) {
	body := c.MustGet(gin.BindKey).(*models.UpdateUserBody)
	user, err := repositories.UserRepo().UpdateUser(body)
	if err != nil {
		httputil.InternalServerError[string](c, err.Error())
		return
	}
	httputil.OK[*models.User](c, user)
}

func (b *Controller) GetUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		httputil.BadRequest[string](c, "Invalid ID")
		return
	}
	user, err := repositories.UserRepo().GetUser(idInt)
	if err != nil {
		httputil.InternalServerError[string](c, err.Error())
		return
	}
	httputil.OK[*models.User](c, user)
}

func (b *Controller) DeleteUser(c *gin.Context) {
	id := c.Param("id") // TODO: middleware for all ID params
	idInt, err := strconv.Atoi(id)
	if err != nil {
		httputil.BadRequest[string](c, "Invalid ID")
		return
	}
	err = repositories.UserRepo().DeleteUser(idInt)
	if err != nil {
		httputil.InternalServerError[string](c, err.Error())
		return
	}
	httputil.NoContent[string](c, "Deleted user successfully")
}

func (b *Controller) ListUsers(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1")) // TODO: middleware for all lists
	if err != nil {
		httputil.BadRequest[string](c, "Invalid page")
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		httputil.BadRequest[string](c, "Invalid size")
		return
	}
	users, err := repositories.UserRepo().ListUsers(page, size)
	if err != nil {
		httputil.InternalServerError[string](c, err.Error())
		return
	}
	httputil.OK[*[]models.User](c, &users)
}
