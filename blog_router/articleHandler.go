package blog_router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaito071831/go_blog/utility"
)

// 記事の型
type Article struct {
	Id int `gorm:"type:int AUTO_INCREMENT"`
	Title string
	Body string
}

// データベースを自動的にマイグレーションする
func init(){
	utility.Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Article{})
}

// 記事の一覧を表示
func Index(c *gin.Context) {
	articlelist := []Article{}
	utility.Db.Find(&articlelist)
	c.HTML(http.StatusOK, "article/index.html", gin.H{
		"list": articlelist,
	})
}

// 記事作成フォームを表示
func New(c *gin.Context) {
	c.HTML(http.StatusOK, "article/new.html", nil)
}

// 記事を作成する
func Create(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		log.Fatalf("フォームの送信に失敗しました: %v", err)
	}
	article := Article{Title: c.PostForm("title"), Body: c.PostForm("body")}
	utility.Db.Create(&article)
	c.Redirect(http.StatusSeeOther, "/article")
}
