package models

import "gorm.io/gorm"

type Configuration struct {
	gorm.Model
	ShowLikes    bool   `json:"show_likes" gorm:"default:true"`
	ShowComments bool   `json:"show_comments" gorm:"default:true"`
	Theme        string `json:"theme" gorm:"default:'light'"`
	Language     string `json:"language" gorm:"default:'en'"`
	SizeTitle    int    `json:"size_title" gorm:"default:20"`
	SizeText     int    `json:"size_text" gorm:"default:14"`
	UserID       uint   `json:"-"`
}

type UpdateConfigBody struct {
	ShowLikes    bool   `json:"show_likes"`
	ShowComments bool   `json:"show_comments"`
	Theme        string `json:"theme"`
	Language     string `json:"language"`
	SizeTitle    int    `json:"size_title"`
	SizeText     int    `json:"size_text"`
}

func FillConfigDefaults(config *Configuration) Configuration {
	config.ShowLikes = true
	config.ShowComments = true  
	config.Theme = "light"
	config.Language = "en"
	config.SizeTitle = 20
	config.SizeText = 14
	return *config
}
