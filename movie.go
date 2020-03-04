package main

import (
	"github.com/gin-gonic/gin"
	"movie-system/app/router"
)

//crud https://blog.csdn.net/wangshubo1989/article/details/70313667
//模板 https://www.popmars.com/category/html%e6%a8%a1%e6%9d%bf/
// https://juejin.im/user/582c029e2e958a0069b1feb2/posts
//图标 https://www.easyicon.net/1225507-camcorder_camera_movie_record_video_icon.html
func main() {

	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")

	engine.Static("/statics", "./statics")

	// 添加网站icon
	engine.StaticFile("/favicon.ico", "./favicon.ico")

	router := initRouter.SetupRouter(engine)

	_ = router.Run(":8080")

}
