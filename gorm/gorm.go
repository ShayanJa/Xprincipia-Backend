package gorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

// db : Package db
var db *gorm.DB

// DB : Global Database Variable
var DB = db

// InitializeDB : Creates a DB Connection and runs migrations
func InitializeDB() {
	//initialize DB
	db, err := gorm.Open("mysql", "root:Popcan123@/xPrincipia")
	if err != nil {
		glog.Error("There was a problem connecting to the database")
	}

	glog.Info("Running DB Migrations...")
	runMigrations(db)

	glog.Info("Populating DB test data...")
	populateDBtestData(db)

}
