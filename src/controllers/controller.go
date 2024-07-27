package controllers

import (
	"html-aiccesible/models"
	"html-aiccesible/repositories"
)

type Controller struct {
	UserRepo   repositories.UserRepository
	ConfigRepo repositories.ConfigRepository
}

func NewController() *Controller {
	db := models.GetDB()
	return &Controller{
		UserRepo:   repositories.UserRepo(db),
		ConfigRepo: repositories.ConfigRepo(db),
	}
}
