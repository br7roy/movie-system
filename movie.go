package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/prometheus/common/log"
	"html/template"
	"moviex/app/kit"
	"moviex/app/middle"
	"moviex/app/router"
	"net/http"
)

//crud https://blog.csdn.net/wangshubo1989/article/details/70313667
//模板 https://www.popmars.com/category/html%e6%a8%a1%e6%9d%bf/
// https://juejin.im/user/582c029e2e958a0069b1feb2/posts
//图标 https://www.easyicon.net/1225507-camcorder_camera_movie_record_video_icon.html
//发布 https://halu.lu/post/goreleaser-tutorial/
//fileServer https://developer.qiniu.com/kodo/sdk/1238/go#3
//https://github.com/qiniu/api.v7/wiki
func main() {

	config := kit.AppConfig

	log.Info(config)

	engine := gin.Default()

	// 添加中间件
	engine.Use(middle.Logger())
	engine.Use(middle.CORS())

	// 引入静态资源
	//staticBox := packr.NewBox("./statics")
	engine.SetHTMLTemplate(initTemplate())
	//engine.StaticFS("/statics", staticBox)

	//engine.LoadHTMLGlob("templates/*")
	//engine.Static("/statics", "./statics")
	//engine.StaticFile("/statics", 	staticBox.String("favicon.ico"))
	// 添加网站icon
	engine.StaticFile("/favicon.ico", "./favicon.ico")

	// 添加上传静态资源目录
	engine.StaticFS("/avatar", http.Dir(kit.RootPath()+"avatar/"))
	// 初始化路由管理
	router := initRouter.SetupRouter(engine)

	gin.SetMode(kit.AppConfig.Server.Mode)
	_ = router.Run(":" + kit.AppConfig.Server.Port)

}

//https://segmentfault.com/q/1010000020928692?sort=created
func initTemplate() *template.Template {
	box := packr.NewBox("./templates")
	t := template.New("")

	tmpl := t.New("index.tmpl")
	data, _ := box.FindString("index.tmpl")
	tmpl.Parse(data)
	tmpl = t.New("401.tmpl")
	data, _ = box.FindString("401.tmpl")
	tmpl.Parse(data)

	tmpl = t.New("error.tmpl")
	data, _ = box.FindString("error.tmpl")
	tmpl.Parse(data)

	tmpl = t.New("footer.tmpl")
	data, _ = box.FindString("footer.tmpl")
	tmpl.Parse(data)

	tmpl = t.New("header.tmpl")
	data, _ = box.FindString("header.tmpl")
	tmpl.Parse(data)

	tmpl = t.New("login.tmpl")
	data, _ = box.FindString("login.tmpl")
	tmpl.Parse(data)

	tmpl = t.New("main.tmpl")
	data, _ = box.FindString("main.tmpl")
	tmpl.Parse(data)

	tmpl = t.New("movie.tmpl")
	data, _ = box.FindString("movie.tmpl")
	tmpl.Parse(data)

	tmpl = t.New("nav.tmpl")
	data, _ = box.FindString("nav.tmpl")
	tmpl.Parse(data)

	tmpl = t.New("user_profile.tmpl")
	data, _ = box.FindString("user_profile.tmpl")
	tmpl.Parse(data)

	tmpl = t.New("cusjs.tmpl")
	data, _ = box.FindString("cusjs.tmpl")
	tmpl.Parse(data)

	return t

}
