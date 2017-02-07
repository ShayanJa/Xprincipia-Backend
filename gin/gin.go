package gin

import (
	"os"
	
	"gopkg.in/gin-gonic/gin.v1"
)

//Global router
var router *gin.Engine

// RunRouter : Used to run gin and all of it's endpoints
func RunRouter() {
	//Get enviromental data
	port := os.Getenv("PORT")

	//gin router config
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//port config
	if port == "" {
		port = "8000"
	}

	router.LoadHTMLGlob("templates/*")

	SetRoutes(router)
	router.Run(":" + port)
}
