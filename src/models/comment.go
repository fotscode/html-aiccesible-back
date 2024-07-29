package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	PostID  uint   `json:"post_id"`
	UserID  uint   `json:"user_id"`
}

type CreateCommentBody struct {
	Title   string `json:"title" binding:"required,min=5,max=100"`
	Content string `json:"content" binding:"required,min=5,max=1000"`
	PostID  uint   `json:"post_id" binding:"required"`
}

type UpdateCommentBody struct {
	ID      uint   `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required,min=5,max=100"`
	Content string `json:"content" binding:"required,min=5,max=1000"`
}
