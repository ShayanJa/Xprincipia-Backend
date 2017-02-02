package gorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

// db : Package db
var db *gorm.DB

// InitializeDB : Creates a DB Connection and runs migrations
func InitializeDB() *gorm.DB {

	//initialize DB
	DB, err := gorm.Open("mysql", "root:Popcan123@/xPrincipia")
	if err != nil {
		glog.Error("There was a problem connecting to the database")
	}
	db = DB // make the database available to all gorm packages

	glog.Info("Running DB Migrations...")
	runMigrations(db)

	glog.Info("Populating DB test data...")
	populateDBtestData(db)

	return db
}
