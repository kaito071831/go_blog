package blog_router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaito071831/go_blog/utility"
)

type Article struct {
	Id int
	Title string
	Body string
}

func init(){
	utility.Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Article{})
}

func Index(c *gin.Context) {
	articlelist := []Article{}
	utility.Db.Find(&articlelist)
	c.HTML(http.StatusOK, "article/index.html", gin.H{
		"list": articlelist,
	})
}
