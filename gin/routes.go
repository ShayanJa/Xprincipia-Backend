package gin

import "github.com/gin-gonic/gin"

// SetRoutes : sets routes for all gin
func SetRoutes(router *gin.Engine) {

	//Index Routes
	router.GET("/", indexHandler)

	//Login Routes
	router.POST("/login", loginHandler)

	//Solutions API
	router.GET("/solutions/ID", getSolutionByIDHandler)
	router.GET("/solutions/problemID", getSolutionByProblemIDHandler)

	//Problems API
	router.GET("/problems/ID", getProblemByIDHandler)

}
