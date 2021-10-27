package database

import "gorm.io/gorm"

var db *gorm.DB

func Initialize(database *gorm.DB) (err error) {
	db = database
	return
}

func GetDBConnection() *gorm.DB {
	return db
}
