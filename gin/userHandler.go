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
	c.Status(http.StatusOK)
}

func logoutHandler(c *gin.Context) {

}

func registerHandler(c *gin.Context) {
	form := gorm.RegistrationForm{}
	c.Bind(&form)

	glog.Info("REGISTERING USER:  " + form.Username)

	c.Status(http.StatusOK)
}
