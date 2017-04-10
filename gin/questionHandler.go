package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
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

func getQuestionByTypeIDHandler(c *gin.Context) {
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}
	questions := gorm.GetAllQuestionsByTypeID(1, intID)

	c.JSON(http.StatusOK, questions)
}

func getAllQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllQuestions())
}
