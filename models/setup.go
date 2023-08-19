package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	db, err := gorm.Open(mysql.Open("root:mamank546@tcp(localhost:8111)/ticket_app"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}, &Product{})

	DB = db
}
