package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var error error
	dsn := "root:@tcp(127.0.0.1:3306)/gormapi?charset=utf8mb4&parseTime=True&loc=Local"
	DB, error = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	}
}
