package gin

import (
	"net/http"
	"work/xprincipia/backend/gorm"

	"gopkg.in/gin-gonic/gin.v1"
)

func postVote(c *gin.Context) {
	form := gorm.VoteForm{}
	c.Bind(&form)
	if gorm.CreateVote(form) {
		c.Status(http.StatusOK)
		return
	}
	c.Status(403)
	return
}
