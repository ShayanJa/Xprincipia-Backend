package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func postComment(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.CommentForm{}
	c.Bind(&form)

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	gorm.CreateComment(form)
	c.Status(http.StatusOK)
}

func getCommentByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("Getting Suggestion with ID : ", id)

	comment := gorm.Comment{}
	intID, err := strconv.Atoi(id)
	uintID := uint(intID)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	comment.GetCommentByID(uintID)
	c.JSON(http.StatusOK, comment)
}

func getAllComments(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllComments())
}

func getCommentsBySuggestionIDHandler(c *gin.Context) {
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}
	comments := gorm.GetAllCommentsBySuggestionID(intID)
	c.JSON(http.StatusOK, comments)
}

func deleteCommentByIDHandler(c *gin.Context) {
	//TODO: change to delete form
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}
	gorm.DeleteCommentByID(intID)
}
