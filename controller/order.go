package controller

import (
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
	productIdStr := c.Query("productId")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}
	userIdStr := c.Query("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}
	result, err := services.AddOrder(productId, userId)
	if err != nil || result != 1 {
		log.Println("生成订单出错err：", err)
		log.Println("生成订单出错result：", result)
		code = e.ERROR
	}
	elapsed := time.Since(startTime)
	log.Println("生成订单接口执行时间：", elapsed)
	appG.Response(http.StatusOK, code, nil)
}
