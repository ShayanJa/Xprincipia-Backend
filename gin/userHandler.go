package gin

import (
	"work/xprincipia/backend/gorm"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

//TODO
//ADD AUTHENTICATION
func loginHandler(c *gin.Context) {

	token := jwt.New(jwt.SigningMethodHS256)
	glog.Info(token)

	loginForm := gorm.LoginForm{}
	c.Bind(&loginForm)

	loginForm.LoginAttempt()
	glog.Info(loginForm)

	c.JSON(200, loginForm)
}

func passwordResetHandler(c *gin.Context) {

	passwordResetForm := gorm.PasswordResetForm{}

	c.Bind(&passwordResetForm)

}

func logoutHandler(c *gin.Context) {

}
