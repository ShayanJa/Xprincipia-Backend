package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func postSuggestion(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	form := gorm.SuggestionForm{}
	c.Bind(&form)
	glog.Info(form)
	gorm.CreateSuggestion(form)

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

func deleteSuggestionByIDHandler(c *gin.Context) {
	id := c.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}
	gorm.DeleteSuggestionByID(intID)
}
