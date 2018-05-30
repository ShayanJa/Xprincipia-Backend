package gin

import (
	"net/http"
	"strconv"
	"../gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func postCon(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.ConForm{}
	c.Bind(&form)

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	gorm.CreateCon(form)
	c.Status(http.StatusOK)
}

func getConByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("Getting Con with ID : ", id)

	con := gorm.Con{}
	intID, err := strconv.Atoi(id)
	uintID := uint(intID)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	con.GetConByID(uintID)
	c.JSON(http.StatusOK, con)
}

func getAllCons(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllCons())
}

func getConByTypeIDHandler(c *gin.Context) {
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
	cons := gorm.GetAllConsByTypeID(intDataType, intID)

	c.JSON(http.StatusOK, cons)
}

// func deleteConByIDHandler(c *gin.Context) {
// 	id := c.Query("id")
// 	intID, err := strconv.Atoi(id)
// 	if err != nil {
// 		glog.Error("There was an error in converting string to integer")
// 	}
// 	gorm.DeleteConByID(intID)
// }

func updateConByIDHandler(c *gin.Context) {
	// Recieve problem Id
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	// Recieve update problem info
	form := gorm.ConForm{}
	c.Bind(&form)

	// Check Token Validity
	err = gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	// Get problem in db
	con := gorm.Con{}
	con.GetConByID(uint(intID))

	// Check if user is actually op
	if con.Username != form.Username {
		c.JSON(401, err.Error())
		return
	}

	//update problem
	con.UpdateCon(form)

}

func deleteConByIDHandler(c *gin.Context) {
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

	form := gorm.ConDeleteForm{ID: intID, Username: username}
	gorm.DeleteConByID(form)
}
