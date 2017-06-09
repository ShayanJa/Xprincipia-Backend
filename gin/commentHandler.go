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

func updateCommentByIDHandler(c *gin.Context) {
	// Recieve problem Id
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	// Recieve update problem info
	form := gorm.CommentForm{}
	c.Bind(&form)

	// Check Token Validity
	err = gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	// Get problem in db
	comment := gorm.Comment{}
	comment.GetCommentByID(uint(intID))

	// Check if user is actually op
	if comment.Username != form.Username {
		c.JSON(401, err.Error())
		return
	}

	//update problem
	comment.UpdateComment(form)

}

func deleteCommentByIDHandler(c *gin.Context) {
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

	form := gorm.CommentDeleteForm{ID: intID, Username: username}
	gorm.DeleteCommentByID(form)
}
