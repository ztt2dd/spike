package routers

import (
	"github.com/gin-gonic/gin"
	"spikeKill/controller"
	"spikeKill/middleware/jwt"
	"spikeKill/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	apiAuth := r.Group("/auth")
	{
		// 新增用户
		apiAuth.POST("/addUser", controller.AddUser)
		// 用户登录授权（获取token）
		apiAuth.GET("/auth", controller.GetAuth)
	}

	apiUser := r.Group("/user")
	apiUser.Use(jwt.JWT())
	{
		// 获取用户信息
		apiUser.GET("/getUser", controller.GetUser)
	}

	apiProduct := r.Group("/product")
	apiProduct.Use(jwt.JWT())
	{
		// 新增商品
		apiProduct.POST("/addProduct", controller.AddProduct)
		// 删除指定商品
		// apiProduct.POST("/delProduct/:id", controller.DelProduct)
		// 修改指定商品
		apiProduct.POST("/updateProduct", controller.UpdateProduct)
		// 查寻指定商品
		apiProduct.GET("/selectProduct/:id", controller.SelectProduct)
		// 分页查寻全部商品
		apiProduct.GET("/selectProductAll", controller.SelectProductAll)
	}

	return r
}
