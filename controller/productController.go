package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spikeKill/models"
	"spikeKill/pkg/e"
	"spikeKill/services"
	"strconv"
)

// 新增商品
func AddProduct(c *gin.Context) {
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
	productService := new(services.ProductService)
	_, err = productService.AddProduct(data)
	if err != nil {

	}
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 更新商品
func UpdateProduct(c *gin.Context) {
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
	models.UpdateProduct(id, data)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 查询指定商品
func SelectProduct(c *gin.Context) {

}

// 分页查询全部商品
func SelectProductAll(c *gin.Context) {
	productName := c.Query("productName")
	// 构造查询参数maps
	maps := make(map[string]interface{})
	// 构造数据返回参数data
	data := make(map[string]interface{})

	if productName != "" {
		maps["productName"] = productName
	}

	// 获取列表数据
	//data["list"] = models.GetProductByPage(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetProductTotal(maps)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
