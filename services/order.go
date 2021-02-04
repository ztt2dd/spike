package services

import (
	"encoding/json"
	"log"
	"spikeKill/models"
	"spikeKill/pkg/redis"
	"spikeKill/services/cacheService"
)

type IOrderService interface {
	AddOrder(productId int, userId int) (int, error)
}

// 创建订单
func AddOrder(productId int, userId int) (int, error) {
	// (1. 从redis中获取数据判断用户是否能够购买)
	if !RedisStock(productId) { // 2. 从redis里获取库存是否足够，不够则进行返回
		return -1, nil
	}
	if !LocalStock(productId) { // 3. redis库存足够的话再查询数据库
		return -2, nil
	}
	if !DeductionRedisStock(productId) { // 4. 库存足够，先在redis里预减库存
		return -3, nil
	}
	// 5. 将创建订单的接口放入队列中
	err := DeductionLocalStock(productId, userId)
	if err != nil {
		log.Println("生成订单失败：", err)
		return -4, err
	}
	return 1, nil
}

// 查询redis的库存数据
func RedisStock(productId int) bool {
	key := cacheService.GetProductKey(productId)
	productData, err := redis.GetData(key)
	if err != nil {
		log.Println("获取产品缓存信息失败：", err)
		return false
	}
	var product *models.Product
	json.Unmarshal(productData, &product)
	if product.ProductNum >= 1 {
		return true
	}
	return false
}

// 查询本地的库存数据
func LocalStock(productId int) bool {
	product, err := models.GetProductById(productId)
	if err != nil {
		log.Println("获取产品库存信息失败：", err)
		return false
	}
	if product.ProductNum >= 1 {
		return true
	}
	return false
}

// 扣减redis的库存
func DeductionRedisStock(productId int) bool {
	lKey := cacheService.GetProductActiveKey(productId)
	redis.Lock(lKey)
	defer redis.UnLock(lKey)
	pKey := cacheService.GetProductKey(productId)
	productData, err := redis.GetData(pKey)
	if err != nil {
		log.Println("获取产品缓存信息失败：", err)
		return false
	}
	var product *models.Product
	json.Unmarshal(productData, &product)
	if product.ProductNum >= 1 {
		product.ProductNum -= 1
		err = redis.SetData(pKey, product)
		if err != nil {
			log.Println("写入产品缓存信息失败：", err)
			return false
		}
		return true
	}
	return false
}

// 生成订单并进行本地库存的扣减
func DeductionLocalStock(productId int, userId int) error {
	return models.CreateLocalOrder(productId, userId)
}
