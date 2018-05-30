package gin

import (
	"net/http"
	"strconv"
	"../gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func postPro(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.ProForm{}
	c.Bind(&form)

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	gorm.CreatePro(form)
}

func getProByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("Getting Pro with ID : ", id)

	pro := gorm.Pro{}
	intID, err := strconv.Atoi(id)
	uintID := uint(intID)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	pro.GetProByID(uintID)
	c.JSON(http.StatusOK, pro)
}

func getAllPros(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllPros())
}

func getProByTypeIDHandler(c *gin.Context) {
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
	pros := gorm.GetAllProsByTypeID(intDataType, intID)

	c.JSON(http.StatusOK, pros)
}

func deleteProByIDHandler(c *gin.Context) {
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

	form := gorm.ProDeleteForm{ID: intID, Username: username}
	gorm.DeleteProByID(form)
}

func updateProByIDHandler(c *gin.Context) {
	// Recieve problem Id
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	// Recieve update problem info
	form := gorm.ProForm{}
	c.Bind(&form)

	// Check Token Validity
	err = gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	// Get problem in db
	p := gorm.Pro{}
	p.GetProByID(uint(intID))

	// Check if user is actually op
	if p.Username != form.Username {
		c.JSON(401, err.Error())
		return
	}

	//update problem
	p.UpdatePro(form)

}
