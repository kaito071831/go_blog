package blog_router

import(
	"net/http"

	"github.com/gin-gonic/gin"
)

func TopHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "static/index.html", gin.H{
		"title": "Hello World",
	})
}
