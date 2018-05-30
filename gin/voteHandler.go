package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"strconv"
	"../gorm"
)

func postVote(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.VoteForm{}
	c.Bind(&form)

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	if gorm.CreateVote(form) {
		c.Status(http.StatusOK)
		return
	}
	c.Status(403)
	return
}

// func getVotesByTypeID(c *gin.Context) {
// 	c.Query()
// }

func isVotedOn(c *gin.Context) {

	Type, err := strconv.Atoi(c.Query("type"))
	typeID, err := strconv.Atoi(c.Query("typeID"))
	if err != nil {
		glog.Info("error")
	}
	username := c.Query("username")

	result := gorm.IsVotedOn(Type, typeID, username)
	c.JSON(http.StatusOK, result)
}

func deleteVote(c *gin.Context) {
	Type, err := strconv.Atoi(c.Query("type"))
	typeID, err := strconv.Atoi(c.Query("typeID"))
	if err != nil {
		glog.Info("error")
	}
	username := c.Query("username")

	// Check Token Validity
	err = gorm.CheckToken(username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}
	gorm.DeleteVote(Type, typeID, username)
	c.Status(http.StatusOK)
	return
}
