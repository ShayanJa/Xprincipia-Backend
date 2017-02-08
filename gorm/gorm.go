package gorm

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

// db : Package db
var db *gorm.DB

// InitializeDB : Creates a DB Connection and runs migrations
func InitializeDB() *gorm.DB {

	//Get Enviromental DB Variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	dbStr := dbUser + ":" + dbPass + "@" + "tcp(" + dbHost + ":" + dbPort + ")" + "/" + dbName + "?charset=utf8&parseTime=true"
	//initialize DB
	DB, err := gorm.Open("mysql", dbStr)
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
