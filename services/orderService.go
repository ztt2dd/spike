package services

import (
	"encoding/json"
	"spikeKill/models"
	"spikeKill/pkg/redis"
	"spikeKill/services/cacheService"
)

func AddOrder(productId int, userId int) (int, error) {
	// 1. 从redis中获取数据判断用户是否能够购买

	// 2. 从redis里获取库存是否足够，不够则进行返回
	count, err := RemoteDeductionStock(productId)
	if err != nil {
		return 0, err
	}
	if count < 1 { // 库存不足
		return -1, nil
	}

	// 3. redis库存足够的话再查询数据库

	// 4. 库存足够，先在redis里预减库存，然后将创建订单的接口放入队列中
	return 1, nil
}

func LocalDeductionStock(productId int) (int, error) {
	key := cacheService.GetProductKey(productId)
	productData, err := redis.GetData(key)
	if err != nil {
		return 0, err
	}
	var product *models.Product
	json.Unmarshal(productData, &product)
	return product.ProductNum, nil
}

func RemoteDeductionStock(productId int) (int, error) {
	key := cacheService.GetProductKey(productId)
	productData, err := redis.GetData(key)
	if err != nil {
		return 0, err
	}
	var product *models.Product
	json.Unmarshal(productData, &product)
	return product.ProductNum, nil
}
