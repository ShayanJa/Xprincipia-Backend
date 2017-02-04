package gin

import (
	"os"

	"gopkg.in/gin-gonic/gin.v1"
)

//Global router
var router *gin.Engine

// RunRouter : Used to run gin and all of it's endpoints
func RunRouter() {
	port := os.Getenv("PORT")
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	if port == "" {
		port = "8080"
	}

	router.LoadHTMLGlob("templates/*")

	SetRoutes(router)
	router.Run(":" + port)
}
