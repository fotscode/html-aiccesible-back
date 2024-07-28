package models

import (
	"fmt"
	ct "html-aiccesible/constants"

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
		ct.MYSQL_USER,
		ct.MYSQL_PASSWORD,
		ct.MYSQL_HOST,
		ct.MYSQL_PORT,
		ct.MYSQL_DATABASE,
	)
	return dsn
}

func CreateDefaultUser() {
	db := GetDB()
	username := ct.ADMIN_USERNAME
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		user.Username = username
		pwd := ct.ADMIN_PASSWORD
		hash, err := HashPassword(pwd)
		if err != nil {
			panic(err)
		}
		user.Password = hash
		user.Config = FillConfigDefaults(&Configuration{})
		db.Create(&user)
	}
}
