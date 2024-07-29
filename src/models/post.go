package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Before      string    `json:"before"`
	After       string    `json:"after"`
	Likes       []*User   `json:"likes" gorm:"many2many:post_likes;"`
	Comments    []Comment `json:"comments"`
	UserID      uint      `json:"user_id"`
}

type CreatePostBody struct {
	Title       string `json:"title" binding:"required,min=4,max=100"`
	Description string `json:"description" binding:"required,min=4,max=100"`
	Before      string `json:"before" binding:"required,min=4,max=100"`
	After       string `json:"after" binding:"required,min=4,max=100"`
}

type UpdatePostBody struct {
	ID          uint   `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required,min=4,max=100"`
	Description string `json:"description" binding:"required,min=4,max=100"`
	Before      string `json:"before" binding:"required,min=4,max=100"`
	After       string `json:"after" binding:"required,min=4,max=100"`
}
