package main

import (
	"flag"
	"work/xprincipia/backend/gin"
	"work/xprincipia/backend/gorm"

	"github.com/golang/glog"
)

func main() {

	//Suppress Parsing Errors
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	flag.CommandLine.Parse([]string{})

	//initialize DB
	glog.Info("INITALIZING DATABASE...")
	db := gorm.InitializeDB()
	defer db.Close()

	//Start HTTP Network
	gin.RunRouter()
}
