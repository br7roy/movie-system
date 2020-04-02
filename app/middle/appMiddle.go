package middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"moviex/app/db"
	"moviex/app/kit"
	"moviex/app/model"
	"net/http"
	"strconv"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		fmt.Printf("%s::%s \t %s \t %s ", time.Now().Format("2006-01-02 15:04:05"), host, url, method)
		context.Next()
		fmt.Println(context.Writer.Status())
	}

}

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, e := context.Request.Cookie(kit.COOKIE_NAME)
		if e == nil {
			context.SetCookie(cookie.Name, cookie.Value, 1000, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

			atoi, _ := strconv.Atoi(cookie.Value)

			user := model.UserInfo[atoi]
			if user.ID < 1 {
				user.ID = atoi
				db.Db.Find(&user)
				model.UserInfo[atoi] = user
			}

			context.Next()
		} else {
			context.Abort()
			context.HTML(http.StatusUnauthorized, "401.tmpl", nil)
		}
	}
}

/*func identityCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, _ := context.Request.Cookie(kit.COOKIE_NAME)
		if cookie != nil {

			atoi, _ := strconv.Atoi(cookie.Value)
			user := kit.UserInfo[atoi]
			if user != nil {

			}

		}

	}
}
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			// TODO: 进行身份验证，比如校验cookie或token
			h(w, r)
		})
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		CurrentTime: time.Now().Unix(),
		Data:    data,
	})
}*/
