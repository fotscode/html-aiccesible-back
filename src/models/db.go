package models

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
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
		gormDB.AutoMigrate(&User{})
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
	var user User
	db.Where("username = ?", os.Getenv("ADMIN_USERNAME")).First(&user)
	if user.ID == 0 {
		user.Username = os.Getenv("ADMIN_USERNAME")
		pwd := os.Getenv("ADMIN_PASSWORD")
		cost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
		if err != nil {
			panic("BCRYPT_COST must be an integer")
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
		if err != nil {
			panic(err)
		}
		user.Password = string(hash)
		db.Create(&user)
	}
}
