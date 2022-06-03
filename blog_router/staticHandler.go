package blog_router

import(
	"net/http"

	"github.com/gin-gonic/gin"
)

// トップページを表示
func TopHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "static/index.html", gin.H{
		"title": "Hello World",
	})
}
