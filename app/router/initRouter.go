package initRouter

import (
	"github.com/gin-gonic/gin"
	"moviex/app/handler"
	"moviex/app/middle"
)

func SetupRouter(engine *gin.Engine) *gin.Engine {

	/*	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"hello": "https://movie.takfu.tk",
			"blog":  "https://blog.takfu.tk",
		})
	})*/

	/*	uGroup := engine.Group("/user")
		{
			uGroup.Any("", handler.UserNameAgeSave)
			//uGroup.Any("/:name", handler.UserSave)
			uGroup.GET("/rj", handler.UserRegisterJump)
			uGroup.POST("/register", handler.UserRegister)
			uGroup.Any("/login", handler.UserLogin)
			uGroup.GET("/profile/", middle.Auth(), handler.UserProfile)
			uGroup.POST("/update", middle.Auth(), handler.UpdateUserProfile)

		}*/

	mGroup := engine.Group("")
	{
		mGroup.GET("", handler.MovieIndex)

		mGroup.GET("/show", handler.ShowAllMovies)

		uGroup := mGroup.Group("/user")
		uGroup.Any("", handler.UserNameAgeSave)
		//uGroup.Any("/:name", handler.UserSave)
		uGroup.GET("/rj", handler.UserRegisterJump)
		uGroup.POST("/register", handler.UserRegister)
		uGroup.Any("/login", handler.UserLogin)
		uGroup.GET("/profile", middle.Auth(), handler.UserProfile)
		uGroup.POST("/update", middle.Auth(), handler.UpdateUserProfile)
	}
	//mGroup := engine.Group("/movie")
	//{
	//	mGroup.GET("/show", handler.ShowAllMovies)
	//}

	return engine
}
