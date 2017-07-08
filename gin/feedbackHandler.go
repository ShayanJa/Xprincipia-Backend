package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func postFeedback(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.FeedbackForm{}
	c.Bind(&form)

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	gorm.CreateFeedback(form)
}

func getFeedbackByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("Getting Feedback with ID : ", id)

	f := gorm.Feedback{}
	intID, err := strconv.Atoi(id)
	uintID := uint(intID)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	f.GetFeedbackByID(uintID)
	c.JSON(http.StatusOK, f)
}

func getAllFeedback(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllFeedback())
}

func deleteFeedbackByIDHandler(c *gin.Context) {
	id := c.Query("id")
	username := c.Query("username")

	// Check Token Validity
	err := gorm.CheckToken(username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	form := gorm.FeedbackDeleteForm{ID: intID, Username: username}
	gorm.DeleteFeedbackByID(form)
}

func updateFeedbackByIDHandler(c *gin.Context) {
	// Recieve problem Id
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	// Recieve update problem info
	form := gorm.FeedbackForm{}
	c.Bind(&form)

	// Check Token Validity
	err = gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	// Get problem in db
	f := gorm.Feedback{}
	f.GetFeedbackByID(uint(intID))

	// Check if user is actually op
	if f.Username != form.Username {
		c.JSON(401, err.Error())
		return
	}

	//update problem
	f.UpdateFeedback(form)

}
