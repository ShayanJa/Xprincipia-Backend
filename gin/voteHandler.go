package gin

import (
	"net/http"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
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
