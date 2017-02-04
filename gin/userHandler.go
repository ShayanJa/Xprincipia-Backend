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

// GetTokenHandler : ~
// var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	/* Create the token */
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	/* Create a map to store our claims
// 	   claims := token.Claims.(jwt.MapClaims)

// 	   /* Set token claims */
// 	// claims["admin"] = true
// 	// claims["name"] = "Ado Kukic"
// 	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

// 	/* Sign the token with our secret */
// 	tokenString, _ := token.SignedString(mySigningKey)

// 	/* Finally, write the token to the browser window */
// 	w.Write([]byte(tokenString))
// })

func passwordResetHandler(c *gin.Context) {

	passwordResetForm := gorm.PasswordResetForm{}

	c.Bind(&passwordResetForm)

}

func logoutHandler(c *gin.Context) {

}
