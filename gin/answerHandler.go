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

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	// Create Answer
	err = gorm.CreateAnswer(form)
	if err != nil {
		// return error response if it exists
		glog.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func getAnswerByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("Getting Answer with ID : ", id)

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

func updateAnswerByIDHandler(c *gin.Context) {
	// Recieve problem Id
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	// Recieve update problem info
	form := gorm.AnswerForm{}
	c.Bind(&form)

	// Check Token Validity
	err = gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	// Get problem in db
	a := gorm.Answer{}
	a.GetAnswerByID(uint(intID))

	// Check if user is actually op
	if a.Username != form.Username {
		c.JSON(401, err.Error())
		return
	}

	//update Answer
	a.UpdateAnswer(form)

}

func deleteAnswerByIDHandler(c *gin.Context) {
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}
	gorm.DeleteAnswerByID(intID)
}
