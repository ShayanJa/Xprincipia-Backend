package main

import (
	"flag"
	"work/xprincipia/backend/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {

	//Suppress Parsing Errors
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	flag.CommandLine.Parse([]string{})

	//goflag.CommandLine.Parse([]string{})

	// //I like how this is written
	// router := gin.Default()

	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        router,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()

	gin.RunRouter()

}
