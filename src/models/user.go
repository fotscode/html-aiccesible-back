package models

import (
	ct "html-aiccesible/constants"
	"strconv"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	cost, err := strconv.Atoi(ct.BCRYPT_COST)
	if err != nil {
		panic("BCRYPT_COST must be an integer")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
