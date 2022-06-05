package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kaito071831/go_blog/blog_router"
)



func main() {
	// ルータを作成
	router := gin.Default()

	// htmlのディレクトリを指定
	router.LoadHTMLGlob("templates/**/*")
	
	// 静的ファイルの場所を指定
	router.Static("/static", "static")

	// URIとハンドラを指定
	router.GET("/", blog_router.TopHandler)

	
	article_group := router.Group("/article")
	article_group.GET("/", blog_router.Index)
	article_group.GET("/new", blog_router.New)
	article_group.POST("/", blog_router.Create)
	article_group.GET("/:id", blog_router.Show)
	article_group.GET("/:id/edit", blog_router.Edit)
	article_group.POST("/:id", blog_router.Update)

	// サーバーを起動
	if err := router.Run(); err != nil {
		log.Fatal("サーバーの起動に失敗しました", err)
	}
}
