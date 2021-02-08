package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"spikeKill/pkg/app"
	"spikeKill/pkg/e"
	"spikeKill/services"
	"strconv"
	"time"
)

// 生成订单
func AddOrder(c *gin.Context) {
	startTime := time.Now()
	appG := app.Gin{C: c}
	code := e.SUCCESS
	productIdStr := c.PostForm("productId")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}
	userIdStr := c.PostForm("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}
	result, err := services.AddOrder(productId, userId)
	fmt.Println("生成订单返回结果result：", result)
	if err != nil || result != 1 {
		log.Println("生成订单出错err：", err)
		code = e.ERROR
	}
	elapsed := time.Since(startTime)
	log.Println("生成订单接口执行时间：", elapsed)
	appG.Response(http.StatusOK, code, nil)
}
