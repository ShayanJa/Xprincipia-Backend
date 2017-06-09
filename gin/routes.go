package gin

import "github.com/gin-gonic/gin"

// SetRoutes : sets routes for all gin
func SetRoutes(router *gin.Engine) {

	// //Index Routes
	router.GET("/", indexHandler)

	//User API
	router.GET("/users/createdSolutions", getAllCreatedSolutions)
	router.GET("/users/followedSolutions", getAllFollowedSolutions)
	router.GET("/users/createdProblems", getAllCreatedProblems)
	router.GET("/users/followedProblems", getAllFollowedProblems)

	//Solutions API
	router.GET("/solutions/ID", getSolutionByID)
	router.GET("/solutions/all", getAllSolutions)
	router.GET("/solutions/problemID", getSolutionsByProblemIDHandler)
	router.POST("/solutions/create", postSolution)

	//Problems API
	router.GET("/problems/ID", getProblemByIDHandler)
	router.GET("/problems/all", getAllProblems)
	router.GET("/problems/subproblems", getAllSubProblems)
	router.POST("/problems/create", postProblem)
	router.GET("/problems/search", searchProblemDB)

	//Questions API
	router.GET("/questions/ID", getQuestionByIDHandler)
	router.GET("/questions/typeID", getQuestionByTypeIDHandler)
	router.GET("/questions/all", getAllQuestions)
	router.POST("/questions/create", postQuestion)

	//Suggestions API
	router.GET("/suggestions/ID", getSuggestionByIDHandler)
	router.GET("/suggestions/typeID", getSuggestionByTypeIDHandler)
	router.GET("/suggestions/all", getAllSuggestions)
	router.POST("/suggestions/create", postSuggestion)

	//Answers API
	router.GET("/answers/ID", getAnswerByIDHandler)
	router.GET("/answers/questionID", getAnswersByQuestionIDHandler)
	router.GET("/answers/all", getAllAnswers)
	router.POST("/answers/create", postAnswer)
	router.DELETE("/answers/ID", deleteAnswerByIDHandler)

	//Comments API
	router.GET("/comments/ID", getCommentByIDHandler)
	router.GET("/comments/questionID", getCommentsBySuggestionIDHandler)
	router.GET("/comments/all", getAllComments)
	router.POST("/comments/create", postComment)

	router.POST("/vote/create", postVote)

	// //Authentication Middleware
	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)

		//User API
		auth.GET("/users/createdSolutions", getAllCreatedSolutions)
		auth.GET("/users/followedSolutions", getAllFollowedSolutions)
		auth.GET("/users/createdProblems", getAllCreatedProblems)
		auth.GET("/users/followedProblems", getAllFollowedProblems)

		//Solutions API
		auth.GET("/solutions/ID", getSolutionByID)
		auth.GET("/solutions/all", getAllSolutions)
		auth.GET("/solutions/problemID", getSolutionsByProblemIDHandler)
		auth.POST("/solutions/create", postSolution)
		auth.PUT("/solutions/update", updateSolutionByIDHandler)

		//Problems API
		auth.GET("/problems/ID", getProblemByIDHandler)
		auth.GET("/problems/all", getAllProblems)
		auth.GET("/problems/subproblems", getAllSubProblems)
		auth.POST("/problems/create", postProblem)
		auth.GET("/problems/search", searchProblemDB)
		auth.PUT("problems/update", updateProblemByIDHandler)

		//Questions API
		auth.GET("/questions/ID", getQuestionByIDHandler)
		auth.GET("/questions/typeID", getQuestionByTypeIDHandler)
		auth.GET("/questions/all", getAllQuestions)
		auth.POST("/questions/create", postQuestion)
		auth.PUT("/questions/update", updateQuestionyIDHandler)
		auth.DELETE("/questions/delete", deleteQuestionByIDHandler)

		//Suggestions API
		auth.GET("/suggestions/ID", getSuggestionByIDHandler)
		auth.GET("/suggestions/typeID", getSuggestionByTypeIDHandler)
		auth.GET("/suggestions/all", getAllSuggestions)
		auth.POST("/suggestions/create", postSuggestion)
		auth.PUT("/suggestions/update", updateSuggestionByIDHandler)

		//Answers API
		auth.GET("/answers/ID", getAnswerByIDHandler)
		auth.GET("/answers/questionID", getAnswersByQuestionIDHandler)
		auth.GET("/answers/all", getAllAnswers)
		auth.POST("/answers/create", postAnswer)
		auth.DELETE("/answers/ID", deleteAnswerByIDHandler)
		auth.PUT("answers/update", updateAnswerByIDHandler)

		//Comments API
		auth.GET("/comments/ID", getCommentByIDHandler)
		auth.GET("/comments/suggestionID", getCommentsBySuggestionIDHandler)
		auth.GET("/comments/all", getAllComments)
		auth.POST("/comments/create", postComment)
		auth.PUT("/comments/update", updateCommentByIDHandler)
		auth.DELETE("/comment/delete", deleteCommentByIDHandler)

		//FreeForm API
		auth.GET("/freeForms/ID", getFreeFormByIDHandler)
		auth.GET("/freeForms/typeID", getFreeFormByTypeIDHandler)
		auth.GET("/freeForms/all", getAllFreeForms)
		auth.POST("/freeForms/create", postFreeForm)
		auth.PUT("/freeForms/update", updateFreeFormByIDHandler)

		//Pro API
		auth.GET("/pros/ID", getProByIDHandler)
		auth.GET("/pros/typeID", getProByTypeIDHandler)
		auth.GET("/pros/all", getAllPros)
		auth.POST("/pros/create", postPro)
		auth.PUT("/pros/update", updateProByIDHandler)

		//Pro API
		auth.GET("/cons/ID", getConByIDHandler)
		auth.GET("/cons/typeID", getConByTypeIDHandler)
		auth.GET("/cons/all", getAllCons)
		auth.POST("/cons/create", postCon)
		auth.PUT("/cons/update", updateConByIDHandler)

		//Learn Item API
		auth.GET("/learnItems/ID", getLearnItemByIDHandler)
		auth.GET("/learnItems/typeID", getLearnItemByTypeIDHandler)
		auth.GET("/learnItems/all", getAllLearnItems)
		auth.POST("/learnItems/create", postLearnItem)
		auth.PUT("learnItems/update", updateLearnItemyIDHandler)

		//Resource API
		auth.GET("/resouces/ID", getResourceByIDHandler)
		auth.GET("/resources/typeID", getResourceByTypeIDHandler)
		auth.GET("/resources/all", getAllResources)
		auth.POST("/resources/create", postResource)
		auth.PUT("/resources/update")

		//Vote API
		auth.POST("/vote/create", postVote)

		//save Token API
		auth.POST("/saveToken", saveToken)

	}
	router.POST("/login", authMiddleware.LoginHandler)

	router.POST("/register", registerHandler)
	//router.POST("/login", loginHandler)

}
