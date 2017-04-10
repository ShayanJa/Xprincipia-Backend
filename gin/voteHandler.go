package gin

import (
	"net/http"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
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