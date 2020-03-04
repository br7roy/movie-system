package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"movie-system/app/model"
	"net/http"
)

func UserSave(context *gin.Context) {
	userName := context.Param("name")
	context.String(http.StatusOK, "用户"+userName+"已保存")
}
func UserNameAgeSave(context *gin.Context) {
	name := context.Query("name")
	age := context.Query("age")
	context.String(http.StatusOK, "user"+name+age+"save")
}

func UserRegisterJump(ctx *gin.Context) {
	ctx.Request.URL.Path = "login.tmpl"
	ctx.HTML(http.StatusOK, "login.tmpl", nil)

}

func UserRegister(ctx *gin.Context) {
	var user model.UserModel

	if er := ctx.ShouldBind(&user); er != nil {
		log.Print("err", er.Error())
		ctx.String(http.StatusBadRequest, "数据不合法")
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/movie")
	}

	//email := ctx.PostForm("email")
	//password := ctx.DefaultPostForm("password", "123")
	//passwordAgain := ctx.DefaultPostForm("password-again", "123")

	println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)

}
