package gin

import "gopkg.in/gin-gonic/gin.v1"

// SetRoutes : sets routes for all gin
func SetRoutes(router *gin.Engine) {

	//Index Routes
	router.GET("/", indexHandler)

	// //Solutions API
	// router.GET("/solutions/ID", getSolutionByID)
	// router.GET("/solutions/all", getAllSolutions)
	// router.GET("/solutions/problemID", getSolutionsByProblemIDHandler)
	// router.POST("/solutions/create", postSolution)

	// //Problems API
	// router.GET("/problems/ID", getProblemByIDHandler)
	// router.GET("/problems/all", getAllProblems)
	// router.GET("/problems/subproblems", getAllSubProblems)
	// router.POST("/problems/create", postProblem)
	// router.GET("/problems/search", searchProblemDB)

	// //Questions API
	// router.GET("/questions/ID", getQuestionByIDHandler)
	// router.GET("/questions/typeID", getQuestionByTypeIDHandler)
	// router.GET("/questions/all", getAllQuestions)
	// router.POST("/questions/create", postQuestion)

	// //Suggestions API
	// router.GET("/suggestions/ID", getSuggestionByIDHandler)
	// router.GET("/suggestions/typeID", getSuggestionByTypeIDHandler)
	// router.GET("/suggestions/all", getAllSuggestions)
	// router.POST("/suggestions/create", postSuggestion)

	//Authentication Middleware
	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		//Solutions API
		auth.GET("/solutions/ID", getSolutionByID)
		auth.GET("/solutions/all", getAllSolutions)
		auth.GET("/solutions/problemID", getSolutionsByProblemIDHandler)
		auth.POST("/solutions/create", postSolution)

		//Problems API
		auth.GET("/problems/ID", getProblemByIDHandler)
		auth.GET("/problems/all", getAllProblems)
		auth.GET("/problems/subproblems", getAllSubProblems)
		auth.POST("/problems/create", postProblem)
		auth.GET("/problems/search", searchProblemDB)

		//Questions API
		auth.GET("/questions/ID", getQuestionByIDHandler)
		auth.GET("/questions/typeID", getQuestionByTypeIDHandler)
		auth.GET("/questions/all", getAllQuestions)
		auth.POST("/questions/create", postQuestion)

		//Suggestions API
		auth.GET("/suggestions/ID", getSuggestionByIDHandler)
		auth.GET("/suggestions/typeID", getSuggestionByTypeIDHandler)
		auth.GET("/suggestions/all", getAllSuggestions)
		auth.POST("/suggestions/create", postSuggestion)
		//auth.POST("/problem", postProblem)
		//auth.POST("/solution", postSolution)
	}
	router.POST("/login", authMiddleware.LoginHandler)

	router.POST("/register", registerHandler)
	//router.POST("/login", loginHandler)

}
