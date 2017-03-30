package main

import (
	"flag"
	"log"
	"work/xprincipia/backend/gin"
	"work/xprincipia/backend/gorm"

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

	s := gorm.SuggestionForm{Description: "fucl"}
	gorm.CreateSuggestion(s)

	//Start HTTP Network
	gin.RunRouter()

}
