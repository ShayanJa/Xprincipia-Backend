package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func getSolutionByID(c *gin.Context) {

	id := c.Query("id")
	glog.Info("ID sent is: ", id)

	solution := gorm.Solution{}
	intID, err := strconv.Atoi(id)
	uintID := uint(intID)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	solution.GetSolutionByID(uintID)
	c.JSON(http.StatusOK, solution)
}

func getSolutionsByProblemIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("ID sent is: ", id)

	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("Unable to convert to int")
	}

	solutions := gorm.GetSolutionsByProblemID(intID)
	c.JSON(http.StatusOK, solutions)
}

func getAllSolutions(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllSolutions())
}

func postSolution(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.SolutionForm{}
	c.Bind(&form)

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	gorm.CreateSolution(form)
	c.Status(http.StatusOK)
}

func deleteSolutionByIDHandler(c *gin.Context) {
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}
	gorm.DeleteSolutionByID(intID)
}

func updateSolutionByIDHandler(c *gin.Context) {
	// Recieve problem Id
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	// Recieve update problem info
	form := gorm.SolutionForm{}
	c.Bind(&form)

	// Check Token Validity
	err = gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	// Get problem in db
	s := gorm.Solution{}
	s.GetSolutionByID(uint(intID))

	// Check if user is actually op
	if s.OriginalPosterUsername != form.Username {
		c.JSON(401, err.Error())
		return
	}

	//update problem
	s.UpdateSolution(form)

}
