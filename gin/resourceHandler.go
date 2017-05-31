package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func postResource(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.ResourceForm{}
	c.Bind(&form)

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	gorm.CreateResource(form)
}

func getResourceByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("Getting Resource with ID : ", id)

	resource := gorm.Resource{}
	intID, err := strconv.Atoi(id)
	uintID := uint(intID)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	resource.GetResourceByID(uintID)
	c.JSON(http.StatusOK, resource)
}

func getAllResources(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllResources())
}

func getResourceByTypeIDHandler(c *gin.Context) {
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
	resources := gorm.GetAllResourcesByTypeID(intDataType, intID)

	c.JSON(http.StatusOK, resources)
}

// func deleteResourceByIDHandler(c *gin.Context) {
// 	id := c.Query("id")
// 	intID, err := strconv.Atoi(id)
// 	if err != nil {
// 		glog.Error("There was an error in converting string to integer")
// 	}
// 	gorm.DeleteResourceByID(intID)
// }
