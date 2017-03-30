package gin

import (
	"net/http"
	"strconv"
	"work/xprincipia/backend/gorm"

	"github.com/golang/glog"
	"gopkg.in/gin-gonic/gin.v1"
)

func postSuggestion(c *gin.Context) {
	form := gorm.SuggestionForm{}
	c.Bind(&form)
	glog.Info(form)
	gorm.CreateSuggestion(form)
	c.Status(http.StatusOK)
}

func getSuggestion(c *gin.Context) {
	id := c.Query("id")
	glog.Info("ID sent is: ", id)

	suggestion := gorm.Suggestion{}
	intID, err := strconv.Atoi(id)
	uintID := uint(intID)
	if err != nil {
		glog.Error("There was an error in converting string to integer")
	}

	suggestion.GetSuggestionByID(uintID)
	c.JSON(http.StatusOK, suggestion)
}
