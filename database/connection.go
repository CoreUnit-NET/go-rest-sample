package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func databaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("fuego.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}
	return db
}
