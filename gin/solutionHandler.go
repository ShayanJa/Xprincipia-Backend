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
	if err != nil {
		glog.Fatal("hey")
	}

	solution.GetSolutionByID(intID)
	c.JSON(http.StatusOK, solution)
}

func getSolutionByProblemID(c *gin.Context) {
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
