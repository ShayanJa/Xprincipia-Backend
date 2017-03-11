package gin

import (
	"app/gorm"
	"net/http"
	"strconv"

	"github.com/golang/glog"
	"gopkg.in/gin-gonic/gin.v1"
)

func getSolutionByIDHandler(c *gin.Context) {
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

func getSolutionByProblemIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("ID sent is: ", id)

	solution := gorm.Solution{}
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("Unable to convert to int")
	}

	solution.GetSolutionByProblemID(intID)
	c.JSON(http.StatusOK, solution)
}

func postSolution(c *gin.Context) {
	form := gorm.SolutionForm{}
	c.Bind(&form)
	glog.Info(form)
	gorm.CreateSolution(form)
	c.Status(http.StatusOK)
}
