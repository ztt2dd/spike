package services

import (
	"encoding/json"
	"fmt"
	"spikeKill/models"
	"spikeKill/pkg/e"
	"spikeKill/pkg/kafka"
	"spikeKill/pkg/redis"
	"spikeKill/services/cacheService"
)

type IOrderService interface {
	AddOrder(productId int, userId int) (int, error)
}

// 创建订单
func AddOrder(productId int, userId int) (int, error) {
	// (1. 从redis中获取数据判断用户是否能够购买)
	if !RedisStock(productId) { // 2. 查询redis库存是否足够
		return -1, nil
	}
	if !LocalStock(productId) { // 3.查询数据库预存是否足够
		return -2, nil
	}
	if !DeductionRedisStock(productId) { // 4.库存充足，redis预减库存
		return -3, nil
	}

	// 5. 将创建订单的接口放入队列中
	err := DeductionKafkaStock(productId, userId)
	if err != nil {
		return -4, err
	}
	return 1, nil
}

// 查询redis的库存数据
func RedisStock(productId int) bool {
	key := cacheService.GetProductKey(productId)
	productData, err := redis.GetData(key)
	if err != nil {
		fmt.Println("获取产品缓存信息失败：", err)
		return false
	}
	var product *models.Products
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
		fmt.Println("获取产品库存信息失败：", err)
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
	err := redis.Lock(lKey)
	if err != nil { // 获取锁失败则直接返回
		return false
	}
	defer redis.UnLock(lKey)
	pKey := cacheService.GetProductKey(productId)
	productData, err := redis.GetData(pKey)
	if err != nil {
		fmt.Println("获取产品缓存信息失败：", err)
		return false
	}
	var product *models.Products
	json.Unmarshal(productData, &product)
	if product.ProductNum >= 1 {
		fmt.Println("redis缓存锁这里到底执行了多少次？：ProductNum：", product.ProductNum)
		product.ProductNum -= 1
		err = redis.SetData(pKey, product)
		if err != nil {
			fmt.Println("写入产品缓存信息失败：", err)
			return false
		} else {
			return true
		}
	}
	return false
}

// 将数据发送到消息队列
func DeductionKafkaStock(productId int, userId int) error {
	order := &models.Orders{
		ProductId: productId,
		UserId:    userId,
	}
	orderJson, err := json.Marshal(order)
	if err != nil {
		return err
	}
	err = kafka.Producer(e.ORDER_TOPIC, string(orderJson))
	if err != nil {
		return err
	}
	return nil
}
