package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		gormDB, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		gormDB.AutoMigrate(
			&User{},
			&Configuration{},
		)
		db = gormDB
	}
	return db
}

func getDsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	return dsn
}

func CreateDefaultUser() {
	db := GetDB()
	username := os.Getenv("ADMIN_USERNAME")
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		user.Username = username
		pwd := os.Getenv("ADMIN_PASSWORD")
		hash, err := HashPassword(pwd)
		if err != nil {
			panic(err)
		}
		user.Password = hash
		user.Config = FillConfigDefaults(&Configuration{})
		db.Create(&user)
	}
}
