package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/itsjamie/gin-cors"
	"net/http"
	"os/exec"
	"time"
)

//Global router
var router *gin.Engine

// RunRouter : Used to run gin and all of it's endpoints
func RunRouter() {
	//Get enviromental data
	port := "10000" //os.Getenv("ROUTER_PORT")
	dockerEntryPoint := "src/work/xprincipia/backend"
	// Gin router config
	// Accept CORS Headers
	router := gin.New()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, POST, PUT, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          10000 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	SetRoutes(router)

	//Set templates
	// router.LoadHTMLGlob("templates/*")

	// Display LOGO
	cmd := "cat " + dockerEntryPoint + "/util/logo.txt"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		glog.Info(fmt.Sprintf("Failed to execute command: %s", cmd))
	}
	glog.Info(string(out))

	//HTTPS
	glog.Info("Listening and serving HTTPS on ", port)
	err = http.ListenAndServeTLS(":"+port, dockerEntryPoint+"/certificates/cert.pem", dockerEntryPoint+"/certificates/privkey.pem", router)
	glog.Info(err)

	//HTTP
	//Run Router on specified port
	// router.Run(":" + port)
}
