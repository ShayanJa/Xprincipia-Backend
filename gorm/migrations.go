package gorm

import (
	"github.com/jinzhu/gorm"
)

//migrations
func runMigrations() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	db.AutoMigrate(&Problem{})

	db.AutoMigrate(&Solution{})

	db.AutoMigrate(&Story{})

	db.AutoMigrate(&Comment{})

	db.AutoMigrate(&Address{})

}
