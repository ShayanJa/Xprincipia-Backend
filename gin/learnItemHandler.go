package gin

import (
	"net/http"
	"strconv"
	"../gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func postLearnItem(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.LearnItemForm{}
	c.Bind(&form)

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	gorm.CreateLearnItem(form)
}

func getLearnItemByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("Getting LearnItem with ID : ", id)

	l := gorm.LearnItem{}
	intID, err := strconv.Atoi(id)
	uintID := uint(intID)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	l.GetLearnItemByID(uintID)
	c.JSON(http.StatusOK, l)
}

func getAllLearnItems(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllLearnItems())
}

func getLearnItemByTypeIDHandler(c *gin.Context) {
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
	learnItems := gorm.GetAllLearnItemsByTypeID(intDataType, intID)

	c.JSON(http.StatusOK, learnItems)
}

func deleteLearnItemByIDHandler(c *gin.Context) {
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

	form := gorm.LearnItemDeleteForm{ID: intID, Username: username}
	gorm.DeleteLearnItemByID(form)
}

func updateLearnItemyIDHandler(c *gin.Context) {
	// Recieve problem Id
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	// Recieve update problem info
	form := gorm.LearnItemForm{}
	c.Bind(&form)

	// Check Token Validity
	err = gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	// Get problem in db
	l := gorm.LearnItem{}
	l.GetLearnItemByID(uint(intID))

	// Check if user is actually op
	if l.Username != form.Username {
		c.JSON(401, err.Error())
		return
	}

	//update problem
	l.UpdateLearnItem(form)

}
