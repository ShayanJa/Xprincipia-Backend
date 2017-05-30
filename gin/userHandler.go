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
	err := gorm.CreateUser(form)
	if err != nil {
		glog.Error(err)
		// c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func getAllCreatedSolutions(c *gin.Context) {
	username := c.Query("username")
	glog.Info("ID sent is: ", username)

	u := gorm.User{}
	u.GetUserByUsername(username)
	c.JSON(http.StatusOK, u.GetAllCreatedSolutions())

}

func getAllFollowedSolutions(c *gin.Context) {
	username := c.Query("username")
	glog.Info("ID sent is: ", username)

	u := gorm.User{}
	u.GetUserByUsername(username)
	c.JSON(http.StatusOK, u.GetAllFollowedSolutions())
}

func getAllCreatedProblems(c *gin.Context) {
	username := c.Query("username")
	glog.Info("ID sent is: ", username)

	u := gorm.User{}
	u.GetUserByUsername(username)
	c.JSON(http.StatusOK, u.GetAllCreatedProblems())

}

func getAllFollowedProblems(c *gin.Context) {
	username := c.Query("username")
	glog.Info("ID sent is: ", username)

	u := gorm.User{}
	u.GetUserByUsername(username)
	c.JSON(http.StatusOK, u.GetAllFollowedProblems())
}

func saveToken(c *gin.Context) {
	form := gorm.LoginAttemptForm{}
	c.Bind(&form)

	gorm.CreateLoginAttempt(form.Username, form.Token)
	c.Status(http.StatusOK)
}
