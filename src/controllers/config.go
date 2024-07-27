package controllers

import (
	"html-aiccesible/httputil"
	"html-aiccesible/models"

	"github.com/gin-gonic/gin"
)

func (b *Controller) GetConfig(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	config, err := b.ConfigRepo.GetConfig(int(user.ID))
	if err != nil {
		httputil.NotFound(c, err)
		return
	}
	httputil.OK(c, config)
}

func (b *Controller) UpdateConfig(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	configBody := c.MustGet(gin.BindKey).(*models.UpdateConfigBody)
	config, err := b.ConfigRepo.UpdateConfig(int(user.ID), configBody)
	if err != nil {
		httputil.InternalServerError(c, err)
		return
	}
	httputil.OK(c, config)
}
