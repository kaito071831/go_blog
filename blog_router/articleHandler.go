package blog_router

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaito071831/go_blog/utility"
)

// 記事の型
type Article struct {
	ID int `gorm:"type:int;autoIncrement;primarykey;<-:false"`
	Title string `gorm:"type:string;not null;<-"`
	Body string `gorm:"type:string;<-"`
	CreatedAt time.Time `gorm:"type:time;<-:create;not null"`
	UpdatedAt time.Time `gorm:"type:time;<-:update"`
}

// データベースを自動的にマイグレーションする
func init(){
	utility.Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Article{})
}

// 記事の一覧を表示
func Index(c *gin.Context) {
	articlelist := []Article{}
	utility.Db.Find(&articlelist)
	c.HTML(http.StatusOK, "article/index.html", articlelist)
}

// 記事の詳細表示
func Show(c *gin.Context) {
	article := Article{}
	id := c.Param("id")
	utility.Db.First(&article, id)
	c.HTML(http.StatusOK, "article/show.html", article)
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

// 記事編集フォームを表示
func Edit(c *gin.Context) {
	article := Article{}
	id := c.Param("id")
	utility.Db.First(&article, id)
	c.HTML(http.StatusOK, "article/edit.html", article)
}

// 記事を更新する
func Update(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		log.Fatalf("フォームの送信に失敗しました: %v", err)
	}

}