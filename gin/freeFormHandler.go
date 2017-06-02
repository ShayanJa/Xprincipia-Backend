package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func postFreeForm(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.FreeFormForm{}
	c.Bind(&form)

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	gorm.CreateFreeForm(form)
	c.Status(http.StatusOK)
}

func getFreeFormByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("Getting FreeForm with ID : ", id)

	freeForm := gorm.FreeForm{}
	intID, err := strconv.Atoi(id)
	uintID := uint(intID)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	freeForm.GetFreeFormByID(uintID)
	c.JSON(http.StatusOK, freeForm)
}

func getAllFreeForms(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllFreeForms())
}

func getFreeFormByTypeIDHandler(c *gin.Context) {
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
	freeForms := gorm.GetAllFreeFormsByTypeID(intDataType, intID)

	c.JSON(http.StatusOK, freeForms)
}

func updateFreeFormByIDHandler(c *gin.Context) {
	// Recieve problem Id
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	// Recieve update problem info
	form := gorm.FreeFormForm{}
	c.Bind(&form)

	// Check Token Validity
	err = gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	// Get problem in db
	f := gorm.FreeForm{}
	f.GetFreeFormByID(uint(intID))

	// Check if user is actually op
	if f.Username != form.Username {
		c.JSON(401, err.Error())
		return
	}

	//update problem
	f.UpdateFreeForm(form)

}

func deleteFreeFormByIDHandler(c *gin.Context) {
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}
	gorm.DeleteFreeFormByID(intID)
}
