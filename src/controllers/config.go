package controllers

import (
	"html-aiccesible/httputil"
	"html-aiccesible/models"
	r "html-aiccesible/repositories"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetConfig(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	config, err := r.ConfigRepo().GetConfig(int(user.ID))
	if err != nil {
		httputil.NotFound(ctx, err)
		return
	}
	httputil.OK(ctx, config)
}

func (c *Controller) UpdateConfig(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	configBody := ctx.MustGet(gin.BindKey).(*models.UpdateConfigBody)
	config, err := r.ConfigRepo().UpdateConfig(int(user.ID), configBody)
	if err != nil {
		httputil.InternalServerError(ctx, err)
		return
	}
	httputil.OK(ctx, config)
}
