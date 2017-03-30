package gin

import (
	"net/http"
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
