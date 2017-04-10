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
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

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
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	form := gorm.RegistrationForm{}
	c.Bind(&form)

	glog.Info("REGISTERING USER:  " + form.Username)
	gorm.CreateUser(form)
	c.Status(http.StatusOK)
}
