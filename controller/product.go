package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spikeKill/models"
	"spikeKill/pkg/app"
	"spikeKill/pkg/e"
	"spikeKill/pkg/setting"
	"spikeKill/pkg/util"
	"spikeKill/services"
	"strconv"
)

// 新增商品
func AddProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	productName := c.Query("productName")
	productNumStr := c.Query("productNum")
	productNum, err := strconv.Atoi(productNumStr)
	if err != nil {
		return
	}
	data := make(map[string]interface{})
	data["productName"] = productName
	data["productNum"] = productNum
	// 处理参数后通过service层处理逻辑
	_, err = services.AddProduct(data)
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
	}
	appG.Response(http.StatusOK, code, nil)
}

// 新增活动
func AddActive(c *gin.Context) {
	appG := app.Gin{C: c}
	code := e.SUCCESS
	productIdStr := c.Query("productId")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}
	err = services.AddProductSpike(productId)
	if err != nil {
		code = e.ERROR
	}
	appG.Response(http.StatusOK, code, nil)
}

// 更新商品
func UpdateProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	idStr := c.Query("id")
	productName := c.Query("productName")
	productNumStr := c.Query("productNum")
	id, err := strconv.Atoi(idStr)
	code := e.SUCCESS
	if err != nil {
		code = e.INVALID_PARAMS
	}
	productNum, err := strconv.Atoi(productNumStr)
	if err != nil {
		code = e.INVALID_PARAMS
	}
	data := make(map[string]interface{})
	data["product_name"] = productName
	data["product_num"] = productNum
	_, err = services.UpdateProduct(id, data)
	if err != nil {
		code = e.ERROR
	}

	appG.Response(http.StatusOK, code, nil)
}

// 分页查询全部商品
func SelectProductAll(c *gin.Context) {
	appG := app.Gin{C: c}
	productName := c.Query("productName")
	// 构造查询参数maps
	maps := make(map[string]interface{})
	// 构造数据返回参数data
	data := make(map[string]interface{})

	if productName != "" {
		maps["productName"] = productName
	}

	// 获取列表数据
	code := e.SUCCESS
	list, err := models.GetProductByPage(util.GetPage(c), setting.AppSetting.PageSize, productName)
	if err != nil {
		code = e.ERROR
	}
	data["list"] = list
	data["total"] = models.GetProductTotal(maps)
	appG.Response(http.StatusOK, code, data)
}
