package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func postLearnContent(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.LearnContentForm{}
	c.Bind(&form)

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	gorm.CreateLearnContent(form)
}

func getLearnContentByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("Getting LearnContent with ID : ", id)

	l := gorm.LearnContent{}
	intID, err := strconv.Atoi(id)
	uintID := uint(intID)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	l.GetLearnContentByID(uintID)
	c.JSON(http.StatusOK, l)
}

func getAllLearnContents(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllLearnContents())
}

func getLearnContentByTypeIDHandler(c *gin.Context) {
	id := c.Query("id")
	dataType := c.Query("dataType")
	glog.Info("ID: ", id)
	glog.Info("dataType: ", dataType)

	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer for id")
	}
	intDataType, err := strconv.Atoi(dataType)
	if err != nil {
		glog.Error("There was an error in converting string to integer for datatype")
	}
	pros := gorm.GetAllLearnContentsByTypeID(intDataType, intID)

	c.JSON(http.StatusOK, pros)
}

// func deleteLearnContentByIDHandler(c *gin.Context) {
// 	id := c.Query("id")
// 	intID, err := strconv.Atoi(id)
// 	if err != nil {
// 		glog.Error("There was an error in converting string to integer")
// 	}
// 	gorm.DeleteLearnContentByID(intID)
// }
