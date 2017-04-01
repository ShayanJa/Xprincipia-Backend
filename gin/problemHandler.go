package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

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

func getAllProblems(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllProblems())
}

func postProblem(c *gin.Context) {
	form := gorm.ProblemForm{}
	c.Bind(&form)
	if form.Description == "" || form.Title == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	gorm.CreateProblem(form)
	c.Status(http.StatusCreated)

}
