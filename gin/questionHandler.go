package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func postQuestion(c *gin.Context) {
	// c.Header("Access-Control-Allow-Origin", "*")
	// c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.QuestionForm{}
	c.Bind(&form)
	glog.Info(form)
	gorm.CreateQuestion(form)
	c.JSON(http.StatusOK, "")
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
	dataType := c.Query("dataType")
	glog.Info("ID: ", id)
	glog.Info("dataType: ", dataType)

	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}
	intDataType, err := strconv.Atoi(dataType)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}
	questions := gorm.GetAllQuestionsByTypeID(intDataType, intID)

	c.JSON(http.StatusOK, questions)
}

func getAllQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllQuestions())
}

func deleteQuestionByIDHandler(c *gin.Context) {
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}
	gorm.DeleteQuestionByID(intID)
}
