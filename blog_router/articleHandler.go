package blog_router

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kaito071831/go_blog/utility"
	"gorm.io/gorm"
)

// 記事の型
type Article struct {
	gorm.Model
	Title string `gorm:"type:string;not null;<-"`
	Body string `gorm:"type:string;<-"`
	UserID uint
}

// データベースの自動設定
func init() {
	utility.Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Article{})
}

// 記事の一覧を表示
func Index(c *gin.Context) {
	username := sessions.Default(c).Get(userKey)
	articlelist := []Article{}
	utility.Db.Where("user_id = ?", getUser(username.(string)).ID).Find(&articlelist)
	c.HTML(http.StatusOK, "article/index.html", gin.H{
		"articlelist": articlelist,
		"username": username,
		"title": "記事一覧",
	})
}

// 記事の詳細表示
func Show(c *gin.Context) {
	article := Article{}
	id := c.Param("id")
	utility.Db.First(&article, id)
	if article.UserID != getUser(sessions.Default(c).Get(userKey).(string)).ID {
		c.HTML(http.StatusNotFound, "errors/404.html", nil)
	} else {
		article.CreatedAt = utility.ConvertJST(article.CreatedAt)
		article.UpdatedAt = utility.ConvertJST(article.UpdatedAt)
		c.HTML(http.StatusOK, "article/show.html", gin.H{
			"article": article,
			"title": fmt.Sprintf("%s | 記事詳細", article.Title), 
		})
	}
}

// 記事作成フォームを表示
func New(c *gin.Context) {
	username := sessions.Default(c).Get(userKey)
	user := getUser(username.(string))
	c.HTML(http.StatusOK, "article/new.html", gin.H{
		"user": user,
		"title": "記事作成",
	})
}

// 記事を作成する
func Create(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		log.Fatalf("フォームの送信に失敗しました: %v", err)
	}
	userid, _ := strconv.Atoi(c.PostForm("userid"))
	article := Article{Title: c.PostForm("title"), Body: c.PostForm("body"), UserID: uint(userid)}
	utility.Db.Create(&article)
	c.Redirect(http.StatusSeeOther, "/article")
}

// 記事編集フォームを表示
func Edit(c *gin.Context) {
	article := Article{}
	id := c.Param("id")
	utility.Db.First(&article, id)
	c.HTML(http.StatusOK, "article/edit.html", gin.H{
		"article": article,
		"title": "記事編集",
	})
}

// 記事を更新する
func Update(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		log.Fatalf("フォームの送信に失敗しました: %v", err)
	}
	article := Article{}
	id := c.Param("id")
	utility.Db.First(&article, id)
	utility.Db.Model(&article).Updates(Article{Title: c.PostForm("title"), Body: c.PostForm("body")})
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/article/%s", id))
}

// 記事を削除する
func Destroy(c *gin.Context) {
	atricle := Article{}
	id := c.Param("id")
	utility.Db.First(&atricle, id)
	utility.Db.Delete(&atricle)
	c.Redirect(http.StatusSeeOther, "/article")
}
