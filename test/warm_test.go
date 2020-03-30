package test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"moviex/app/router"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	if mode := gin.Mode(); mode == gin.DebugMode {
		engine.LoadHTMLGlob("./templates/*")
		engine.Static("../statics", "../statics")
	} else {
		engine.LoadHTMLGlob("templates/*")
		engine.Static("../statics", "../statics")
	}
	router = initRouter.SetupRouter(engine)

}

/*func TestUserSave(t *testing.T) {
	username := "lisi"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已保存", w.Body.String())

}*/

func TestUserNameAgeSave(t *testing.T) {
	name := "zhangsan"
	age := 30
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+name+"&age="+strconv.Itoa(age), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "user"+name+(strconv.Itoa(age))+"save", w.Body.String())

}
func TestUserPostForm(t *testing.T) {
	value := url.Values{}
	value.Add("email", "123@gmail.com")
	value.Add("password", "123")
	value.Add("password-again", "123")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/movie/user/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusMovedPermanently, w.Code)
}

func TestUserLogin(t *testing.T) {
	email := "123@gmail.com"
	value := url.Values{}
	value.Add("email", email)
	value.Add("password", "123")
	value.Add("password-again", "123")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/movie/user/login", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, strings.Contains(w.Body.String(), email), true)
}
