package tests

import (
	database "github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"../gorm"
)

//sets up the db for testing
func SetupTestingDB() *database.DB {
	err := godotenv.Load("../config/config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := gorm.InitializeDB()
	return db
}
