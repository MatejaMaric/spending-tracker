package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbEnv := os.Getenv("DB_CONN")

	var err error
	DB, err = gorm.Open(postgres.Open(dbEnv), &gorm.Config{})
	if err != nil {
		log.Println("Can't connect to database!")
	} else {
		DB.AutoMigrate(&Transaction{})
	}
}
