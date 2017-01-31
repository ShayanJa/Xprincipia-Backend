package gin

import "github.com/gin-gonic/gin"

//Global router
var router *gin.Engine

// RunRouter : Used to run gin and all of it's endpoints
func RunRouter() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	SetRoutes(router)
	router.Run(":8080")

}
