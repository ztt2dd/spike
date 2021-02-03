package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spikeKill/pkg/app"
	"spikeKill/pkg/e"
	"spikeKill/services"
	"strconv"
)

var userService = new(services.UserService)

// 新增用户
func AddUser(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	password := c.Query("password")
	_, err := userService.AddUser(name, password)
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
	}
	appG.Response(http.StatusOK, code, nil)
}

// 授权登录
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	password := c.Query("password")
	user, err := userService.GetAuth(name, password)
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR_LOGIN
	}
	appG.Response(http.StatusOK, code, user)
}

// 获取用户信息
func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}
	user, err := userService.GetUserInfo(id)
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR_AUTH
	}
	appG.Response(http.StatusOK, code, user)
}
