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

	//Get Enviromental DB Variables
	dbHost := "localhost"  //os.Getenv("DB_HOST")
	dbPort := "3306"       //os.Getenv("DB_PORT")
	dbName := "xPrincipia" //os.Getenv("DB_NAME")
	dbUser := "root"       //os.Getenv("DB_USER")
	dbPass := "Popcan123"  //os.Getenv("DB_PASS")

	//initialize DB
	dbStr := dbUser + ":" + dbPass + "@" + "tcp(" + dbHost + ":" + dbPort + ")" + "/" + dbName + "?charset=utf8&parseTime=true"
	DB, err := gorm.Open("mysql", dbStr)
	if err != nil {
		glog.Error("There was a problem connecting to the database")
	}
	db = DB // make the database available to all gorm packages

	glog.Info("Running DB Migrations...")
	isDBEmpty := runMigrations(db)

	if isDBEmpty {
		glog.Info("Populating DB test data...")
		populateDBtestData(db)
	}
	return db
}
