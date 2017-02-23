package gin

import (
	"net/http"
	"work/xprincipia/backend/gorm"

	"github.com/golang/glog"
	"gopkg.in/gin-gonic/gin.v1"
)

//TODO
//ADD AUTHENTICATION
func loginHandler(c *gin.Context) {

	//Create jwt Token to authenticate user
	// token := jwt.New(jwt.SigningMethodHS256)
	// glog.Info(token)

	loginForm := gorm.LoginForm{}
	c.Bind(&loginForm)

	//Log login attempt
	loginForm.LoginAttempt()
	glog.Info(loginForm)

	c.JSON(http.StatusOK, loginForm)
}

func passwordResetHandler(c *gin.Context) {

	passwordResetForm := gorm.PasswordResetForm{}

	c.Bind(&passwordResetForm)

}

func logoutHandler(c *gin.Context) {

}
