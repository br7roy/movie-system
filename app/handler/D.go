package handler

import (
	"github.com/gin-gonic/gin"
	"moviex/app/kit"
	"moviex/app/model"
	"strconv"
)

func GetUserIfAvaliable(ctx *gin.Context) model.User {
	cookie, er := ctx.Request.Cookie(kit.COOKIE_NAME)

	var user = model.User{}

	if er == nil {
		atoi, _ := strconv.Atoi(cookie.Value)

		user = model.UserInfo[atoi]
	}
	return user
}
