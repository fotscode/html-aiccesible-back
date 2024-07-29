package controllers

import (
	"html-aiccesible/models"
	"html-aiccesible/repositories"
)

type Controller struct {
	UserRepo    repositories.UserRepository
	ConfigRepo  repositories.ConfigRepository
	PostRepo    repositories.PostRepository
	CommentRepo repositories.CommentRepository
}

func NewController() *Controller {
	db := models.GetDB()
	return &Controller{
		UserRepo:    repositories.UserRepo(db),
		ConfigRepo:  repositories.ConfigRepo(db),
		PostRepo:    repositories.PostRepo(db),
		CommentRepo: repositories.CommentRepo(db),
	}
}
