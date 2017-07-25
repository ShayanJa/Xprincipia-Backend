package gin

import (
	"time"
	"work/xprincipia/backend/gorm"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"golang.org/x/crypto/bcrypt"
)

// the jwt middleware
var authMiddleware = &jwt.GinJWTMiddleware{
	Realm:      "test zone",
	Key:        []byte("secret key"),
	Timeout:    time.Hour * 168,
	MaxRefresh: time.Hour * 168,
	Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
		//  
		user := gorm.User{}
		passwordBytes := []byte(password)

		if user.GetUserByUsername(userId) {
			hashedPassword := user.HashedPassword
			err := bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)
			if err == nil {
				glog.Info("USER LOG IN SUCCESSFUL...")
				return userId, true
			}
			glog.Info(err)
		}
		return userId, false
	},
	Authorizator: func(userId string, c *gin.Context) bool {
		//  
		//check if this user is in the db based on the jwt

		return gorm.IsUserinDBbyUsername(userId)
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	},
	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"
	TokenLookup: "header:Authorization",
	// TokenLookup: "query:token",
	// TokenLookup: "cookie:token",
}
