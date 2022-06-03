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
	router.GET("/article", blog_router.Index)

	// サーバーを起動
	if err := router.Run(); err != nil {
		log.Fatal("サーバーの起動に失敗しました", err)
	}
}
