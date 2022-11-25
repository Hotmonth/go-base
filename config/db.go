package config

import (
	"gobase/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("base.db"))
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{})
	DB = db
}
