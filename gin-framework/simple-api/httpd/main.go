package main

import (
	"github.com/azuladotoujours/go-framework-tests/gin-framework/simple-api/httpd/handler"
	"github.com/azuladotoujours/go-framework-tests/gin-framework/simple-api/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	r.Run() // listen and serve on 0.0.0.0:8080
}
