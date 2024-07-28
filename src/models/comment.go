package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	PostID  uint   `json:"post_id"`
	UserID  uint   `json:"user_id"`
}
