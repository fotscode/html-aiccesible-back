package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string        `json:"username" gorm:"unique"`
	Password string        `json:"-"`
	Config   Configuration `json:"-"`
	// TODO: add more fields
	// Posts      Post[]
	// Comments   Comment[]
	// Likes      Post[]
}

type CreateUserBody struct {
	Username string `json:"username" binding:"required,min=4,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type UpdateUserBody struct {
	ID       uint   `json:"id" binding:"required"`
	Username string `json:"username" validate:"required,min=4,max=20"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type LoginUserBody struct {
	Username string `json:"username" binding:"required,min=4,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
