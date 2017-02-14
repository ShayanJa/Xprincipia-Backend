package gin

import (
	"net/http"

	"github.com/golang/glog"
	"gopkg.in/gin-gonic/gin.v1"
)

func indexHandler(c *gin.Context) {
	glog.Info("Accesing index Page ...")
	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "XPrincipia",
		},
	)

}
func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"text": "Hello World.",
	})
}
