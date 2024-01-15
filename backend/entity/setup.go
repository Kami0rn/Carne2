package entity

import (
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("Carne.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema
	database.AutoMigrate(
		&User{},
		&Breed{},
		&Result{},
		&Wagyu{},
		Wallet{},
	)

	db = database

}
