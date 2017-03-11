package gin

import (
	"app/gorm"
	"net/http"
	"strconv"

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

	problem.GetProblemByID(uint(intID))
	c.JSON(http.StatusOK, problem)
}

func postProblem(c *gin.Context) {
	form := gorm.ProblemForm{}
	c.Bind(&form)
	if form.Description == "" || form.Title == "" {
		c.Status(400)
		return
	}

	gorm.CreateProblem(form)
	c.Status(http.StatusOK)

}
