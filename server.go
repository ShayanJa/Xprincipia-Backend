package main

import (
	"flag"
	"log"
	"work/xprincipia/xGObackEnd/gin"
	"work/xprincipia/xGObackEnd/gorm"

	"github.com/golang/glog"
	"github.com/joho/godotenv"
)

func main() {

	// Load Enviromental Variables
	err := godotenv.Load("config/config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Suppress Parsing Errors
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	flag.CommandLine.Parse([]string{})

	//initialize DB
	glog.Info("INITALIZING DATABASE...")
	db := gorm.InitializeDB()
	defer db.Close()

	//testing
	glog.Info("Testing >>>>")
	y := gorm.RegistrationForm{}
	y.Username = "jackDanielss"
	y.Password = "bill"
	y.Email = "shay.talebi@gmail.com"

	gorm.CreateUser(y)
	//testing done

	//Start HTTP Network
	gin.RunRouter()

}
