package main

import (
	"log"
	"net/http"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/kaito071831/go_blog/blog_router"
	"github.com/kaito071831/go_blog/utility"
)

func main() {
	// ルータを作成
	router := gin.Default()

	// クッキーストアを生成する
	store := gormsessions.NewStore(utility.Db, true, []byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// htmlのディレクトリを指定
	router.LoadHTMLGlob("templates/**/*")
	
	// 静的ファイルの場所を指定
	router.Static("/static", "static")

	// URIとハンドラを指定
	router.GET("/", blog_router.TopHandler)
	router.GET("/signup", blog_router.Signup)
	router.POST("/signup", blog_router.Signup)
	router.GET("/login", blog_router.Login)
	router.POST("/login", blog_router.Login)
	router.GET("/logout", blog_router.Logout)

	article_group := router.Group("/article")
	article_group.GET("/", blog_router.Index)
	article_group.GET("/new", blog_router.New)
	article_group.POST("/", blog_router.Create)
	article_group.GET("/:id", blog_router.Show)
	article_group.GET("/:id/edit", blog_router.Edit)
	article_group.POST("/:id", blog_router.Update)
	article_group.GET("/:id/delete", blog_router.Destroy)

	router.NoRoute(func(c *gin.Context) {
		title := "404 Notfound"
		c.HTML(http.StatusNotFound, "errors/404.html", gin.H{
			"title": title,
		})
	})

	// サーバーを起動
	if err := router.Run(); err != nil {
		log.Fatal("サーバーの起動に失敗しました", err)
	}
}
