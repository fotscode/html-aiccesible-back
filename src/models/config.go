package models

import "gorm.io/gorm"

type Configuration struct {
	gorm.Model
	ShowLikes    bool   `json:"show_likes" gorm:"default:true"`
	ShowComments bool   `json:"show_comments" gorm:"default:true"`
	Theme        string `json:"theme" gorm:"default:'light'"`
	Language     string `json:"language" gorm:"default:'es'"`
	SizeTitle    float32`json:"size_title" gorm:"default:1.0"`
	SizeText     float32`json:"size_text" gorm:"default:1.0"`
	UserID       uint   `json:"-"`
}

type UpdateConfigBody struct {
	ShowLikes    bool   `json:"show_likes"`
	ShowComments bool   `json:"show_comments"`
	Theme        string `json:"theme"`
	Language     string `json:"language"`
	SizeTitle    float32`json:"size_title"`
	SizeText     float32`json:"size_text"`
}

func FillConfigDefaults(config *Configuration) Configuration {
	config.ShowLikes = true
	config.ShowComments = true  
	config.Theme = "light"
	config.Language = "es"
	config.SizeTitle = 1.0
	config.SizeText = 1.0
	return *config
}
