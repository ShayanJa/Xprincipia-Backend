package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func indexHandler(c *gin.Context) {
	glog.Info("Hey")
	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Home Page",
		},
	)

}
