package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

	"github.com/golang/glog"
	"gopkg.in/gin-gonic/gin.v1"
)

func postQuestion(c *gin.Context) {
	form := gorm.QuestionForm{}
	c.Bind(&form)
	glog.Info(form)
	gorm.CreateQuestion(form)
	c.Status(http.StatusOK)
}

func getQuestionByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("ID sent is: ", id)

	question := gorm.Question{}
	intID, err := strconv.Atoi(id)
	uintID := uint(intID)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	question.GetQuestionByID(uintID)
	c.JSON(http.StatusOK, question)
}
