package main

import (
	"flag"
	"work/xprincipia/backend/gin"
)

func main() {

	//Suppress Parsing Errors
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	flag.CommandLine.Parse([]string{})

	//Start HTTP Network
	gin.RunRouter()

}
