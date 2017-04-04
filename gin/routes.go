package gin

import "gopkg.in/gin-gonic/gin.v1"

// SetRoutes : sets routes for all gin
func SetRoutes(router *gin.Engine) {

	//Index Routes
	router.GET("/", indexHandler)

	//Solutions API
	router.GET("/solutions/ID", getSolutionByID)
	router.GET("/solutions/all", getAllSolutions)
	router.GET("/solutions/problemID", getSolutionByProblemIDHandler)
	router.POST("/solutions/create", postSolution)

	//Problems API
	router.GET("/problems/ID", getProblemByIDHandler)
	router.GET("/problems/all", getAllProblems)
	router.POST("/problems/create", postProblem)
	router.GET("/problems/search", searchProblemDB)

	//Questions API
	router.GET("/questions/ID", getQuestionByIDHandler)
	router.POST("/questions/create", postQuestion)

	//Suggestions API
	router.GET("/suggestions/ID", getSuggestionByIDHandler)
	router.POST("/suggestions/create", postSuggestion)

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

	router.POST("/register", registerHandler)
	//router.POST("/login", loginHandler)

}
