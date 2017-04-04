package gin

import (
	"fmt"
	"os/exec"

	"github.com/golang/glog"
	"gopkg.in/gin-gonic/gin.v1"
)

//Global router
var router *gin.Engine

// RunRouter : Used to run gin and all of it's endpoints
func RunRouter() {
	//Get enviromental data
	port := "10000" //os.Getenv("ROUTER_PORT")

	//gin router config
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	SetRoutes(router)

	//Set templates
	// router.LoadHTMLGlob("templates/*")

	// Display LOGO
	cmd := "cat image.txt"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		glog.Info(fmt.Sprintf("Failed to execute command: %s", cmd))
	}
	glog.Info(string(out))

	//Run Router on specified port
	router.Run(":" + port)
}
