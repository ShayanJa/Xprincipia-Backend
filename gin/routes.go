package gin

import "github.com/gin-gonic/gin"

// SetRoutes : sets routes for all gin
func SetRoutes(router *gin.Engine) {
	router.GET("/", indexHandler)
	router.GET("/solutions/ID", getSolutionByIDHandler)
	router.GET("/solutions/problemID", getSolutionByProblemIDHandler)

	//router.GET("/problems/ID", getProblemByID)
}
