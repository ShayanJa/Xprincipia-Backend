package gin

import (
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

//TODO
func loginHandler(c *gin.Context) {

	loginForm := gorm.LoginForm{}

	c.Bind(&loginForm)
	glog.Info(loginForm)
	c.JSON(200, loginForm)
}

func passwordResetHandler(c *gin.Context) {

}

func logoutHandler(c *gin.Context) {

}
