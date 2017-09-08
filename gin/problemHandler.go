package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"
)

func getProblemByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("ID sent is: ", id)

	problem := gorm.Problem{}
	intID, err := strconv.Atoi(id)

	if err != nil {
		glog.Error("Unable to convert to int")
	}

	problem.GetProblemByID(uint(intID))
	c.JSON(http.StatusOK, problem)
}

func getAllProblems(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllProblemsWithLimit(10))
}

func getAllSubProblems(c *gin.Context) {

	id := c.Query("id")
	glog.Info("ID sent is: ", id)
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("Unable to convert to int")
	}

	problems := gorm.GetSubProblemsByID(intID)

	c.JSON(http.StatusOK, problems)
}

func postProblem(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.ProblemForm{}
	c.Bind(&form)

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	err = gorm.CreateProblem(form)
	if err != nil {
		glog.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.Status(http.StatusOK)

}

func postPrivateProblem(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.ProblemForm{}
	c.Bind(&form)

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	err = gorm.CreatePrivateProblem(form)
	if err != nil {
		glog.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.Status(http.StatusOK)

}

func searchProblemDB(c *gin.Context) {
	//WIP : Only shows search query based on name
	query := c.Query("q")

	//Query for last 10 problems
	response := gorm.QueryLast10Problems(query)

	glog.Info("Query value: " + query)
	glog.Info("length pf Query Response : " + string(len(response)))

	c.JSON(http.StatusOK, response)

}

func deleteProblemByIDHandler(c *gin.Context) {
	id := c.Query("id")
	username := c.Query("username")

	// Check Token Validity
	err := gorm.CheckToken(username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	form := gorm.ProblemDeleteForm{ID: intID, Username: username}
	gorm.DeleteProblemByID(form)
}

func updateProblemByIDHandler(c *gin.Context) {
	// Recieve problem Id
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	// Recieve update problem info
	form := gorm.ProblemForm{}
	c.Bind(&form)

	// Check Token Validity
	err = gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	// Get problem in db
	p := gorm.Problem{}
	p.GetProblemByID(uint(intID))

	// Check if user is actually op
	if p.OriginalPosterUsername != form.Username {
		c.JSON(401, err.Error())
		return
	}

	//update problem
	p.UpdateProblem(form)

}
