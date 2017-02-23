package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/xGObackend/gorm"

	"github.com/golang/glog"
	"gopkg.in/gin-gonic/gin.v1"
)

func getProblemByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("ID sent is: ", id)

	problem := gorm.Problem{}
	intID, err := strconv.Atoi(id)

	if err != nil {
		glog.Error("Unable to convert to int")
	}

	problem.GetProblemByID(intID)
	c.JSON(http.StatusOK, problem)
}
