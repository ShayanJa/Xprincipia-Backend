package gin

import "gopkg.in/gin-gonic/gin.v1"

// SetRoutes : sets routes for all gin
func SetRoutes(router *gin.Engine) {

	//Index Routes
	router.GET("/", indexHandler)

	//Solutions API
	router.GET("/solutions/ID", getSolutionByIDHandler)
	router.GET("/solutions/problemID", getSolutionByProblemIDHandler)
	router.POST("/solutions/create", postSolution)

	//Problems API
	router.GET("/problems/ID", getProblemByIDHandler)
	router.GET("/problems/all", getAllProblems)
	router.POST("/problems/create", postProblem)

	//Questions API
	router.GET("/questions/ID", getProblemByIDHandler)
	router.POST("/questions/create", postProblem)

	//Suggestions API
	router.POST("/suggestions/create", postSuggestion)
	router.GET("/suggestions/ID", getSuggestion)

	//Authentication Middleware
	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		//auth.POST("/problem", postProblem)
		//auth.POST("/solution", postSolution)
	}
	router.POST("/login", authMiddleware.LoginHandler)

	//Login Routes
	//router.POST("/login", loginHandler)

}
