package gin

import "gopkg.in/gin-gonic/gin.v1"

//Global router
var router *gin.Engine

// RunRouter : Used to run gin and all of it's endpoints
func RunRouter() {
	//Get enviromental data
	port := "10000" //os.Getenv("ROUTER_PORT")

	//gin router config
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.LoadHTMLGlob("templates/*")

	SetRoutes(router)
	router.Run(":" + port)
}
