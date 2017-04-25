package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func postAnswer(c *gin.Context) {
	form := gorm.AnswerForm{}
	c.Bind(&form)
	glog.Info(form)
	gorm.CreateAnswer(form)
	c.Status(http.StatusOK)
}

func getAnswerByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("Getting Suggestion with ID : ", id)

	answer := gorm.Answer{}
	intID, err := strconv.Atoi(id)
	uintID := uint(intID)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	answer.GetAnswerByID(uintID)
	c.JSON(http.StatusOK, answer)
}

func getAllAnswers(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllAnswers())
}

func getAnswersByQuestionIDHandler(c *gin.Context) {
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}
	answers := gorm.GetAllAnswersByQuestionID(intID)
	c.JSON(http.StatusOK, answers)
}

func deleteAnswerByIDHandler(c *gin.Context) {
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}
	gorm.DeleteAnswerByID(intID)
}
