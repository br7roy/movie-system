package main

import (
	"github.com/gin-gonic/gin"
	"movie-system/app/service"
	"net/http"
)

//crud https://blog.csdn.net/wangshubo1989/article/details/70313667
//模板 https://www.popmars.com/category/html%e6%a8%a1%e6%9d%bf/
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
	engine.GET("/show", func(context *gin.Context) {
		movies := service.GetAllMovies()
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"movies": movies,
		})

	})

	_ = engine.Run(":8080")

}
