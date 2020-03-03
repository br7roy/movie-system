package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//https://blog.csdn.net/wangshubo1989/article/details/70313667
func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"hello": "https://takfu.tk/movie",
			"blog":  "https://takblog.netlify.com",
		})
	})
	engine.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Movie coming.",
			"text":  "Hello Gin",
		})
	})
	engine.Run(":8080")

}
