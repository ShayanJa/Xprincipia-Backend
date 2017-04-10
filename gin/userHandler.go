package gin

import (
	"net/http"
	"work/xprincipia/backend/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
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
	gorm.CreateUser(form)
	c.Status(http.StatusOK)
}
