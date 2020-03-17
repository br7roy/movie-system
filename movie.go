package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/prometheus/common/log"
	"movie-system/app/kit"
	"movie-system/app/middle"
	"movie-system/app/router"
	"net/http"
)

//crud https://blog.csdn.net/wangshubo1989/article/details/70313667
//模板 https://www.popmars.com/category/html%e6%a8%a1%e6%9d%bf/
// https://juejin.im/user/582c029e2e958a0069b1feb2/posts
//图标 https://www.easyicon.net/1225507-camcorder_camera_movie_record_video_icon.html
func main() {

	config := kit.AppConfig

	log.Info(config)

	engine := gin.Default()

	// 添加中间件
	engine.Use(middle.Logger())

	// 引入静态资源
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/statics", "./statics")

	// 添加网站icon
	engine.StaticFile("/favicon.ico", "./favicon.ico")

	// 添加上传静态资源目录
	engine.StaticFS("/avatar", http.Dir(kit.RootPath()+"avatar/"))
	// 初始化路由管理
	router := initRouter.SetupRouter(engine)

	gin.SetMode(kit.AppConfig.Server.Mode)
	_ = router.Run(":" + kit.AppConfig.Server.Port)

}
