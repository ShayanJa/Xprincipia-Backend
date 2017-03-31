package gorm

import (
	"os"
	"time"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

// db : Package db
var db *gorm.DB

// InitializeDB : Creates a DB Connection and runs migrations
func InitializeDB() *gorm.DB {

	//Get Enviromental DB Variables
	//dbHost := "172.18.0.2" //Bens Mysql xPrincipia id
	//dbHost := "172.19.0.2" //Shayan's Mysql xPrincipia id

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	//initialize DB
	dbStr := dbUser + ":" + dbPass + "@" + "tcp(" + dbHost + ":" + dbPort + ")" + "/" + dbName + "?charset=utf8&parseTime=true"
	DB, err := gorm.Open("mysql", dbStr)

	// glog.Info(dbStr)
	//Try connecting to the database 10 more times
	//We must wait for the Mysql Service to finish building
	for i := 1; i < 10; i++ {
		if err == nil {
			break
		}
		time.Sleep(3000 * time.Millisecond) //sleep for a bit
		s := strconv.Itoa(i)
		glog.Error("Trying to Connect to DB...    Attempt " + s)

		DB, err = gorm.Open("mysql", dbStr)
		// DB, err = gorm.Open("mysql", "root:Popcan123@/xPrincipia?charset=utf8&parseTime=True&loc=Local")
	}

	//If unable to connect, exit
	if err != nil {
		glog.Fatal("There was a problem connecting to the database")
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
