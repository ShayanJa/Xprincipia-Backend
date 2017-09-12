package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func postSuggestion(c *gin.Context) {

	form := gorm.SuggestionForm{}
	c.Bind(&form)

	// Check Token Validity
	err := gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	// Create Suggestion
	err = gorm.CreateSuggestion(form)
	if err != nil {
		// return error response if it exists
		glog.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func getSuggestionByIDHandler(c *gin.Context) {
	id := c.Query("id")
	glog.Info("Getting Suggestion with ID : ", id)

	suggestion := gorm.Suggestion{}
	intID, err := strconv.Atoi(id)
	uintID := uint(intID)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	suggestion.GetSuggestionByID(uintID)
	c.JSON(http.StatusOK, suggestion)
}

func getAllSuggestions(c *gin.Context) {
	c.JSON(http.StatusOK, gorm.GetAllSuggestions())
}

func getSuggestionByTypeIDHandler(c *gin.Context) {
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
	suggestions := gorm.GetAllSuggestionsByTypeID(intDataType, intID)

	c.JSON(http.StatusOK, suggestions)
}

func updateSuggestionByIDHandler(c *gin.Context) {
	// Recieve problem Id
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	// Recieve update problem info
	form := gorm.SuggestionForm{}
	c.Bind(&form)

	// Check Token Validity
	err = gorm.CheckToken(form.Username, c.Request.Header["Authorization"][0])
	if err != nil {
		//if Token not in table
		c.JSON(401, err.Error())
		return
	}

	// Get problem in db
	s := gorm.Suggestion{}
	s.GetSuggestionByID(uint(intID))

	// Check if user is actually op
	if s.Username != form.Username {
		c.JSON(401, err.Error())
		return
	}

	//update suggestion
	s.UpdateSuggestion(form)

}

func deleteSuggestionByIDHandler(c *gin.Context) {
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

	form := gorm.SuggestionDeleteForm{ID: intID, Username: username}
	gorm.DeleteSuggestionByID(form)
}
