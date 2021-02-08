package services

import (
	"spikeKill/models"
	"spikeKill/pkg/redis"
	"spikeKill/services/cacheService"
)

func AddProduct(data map[string]interface{}) (int, error) {
	// 通过model层实现数据库操作
	err := models.AddProduct(data)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func UpdateProduct(id int, data map[string]interface{}) (int, error) {
	err := models.UpdateProduct(id, data)
	if err != nil {
		return 0, nil
	}
	return 1, nil
}

func SelectProductById(id int) (*models.Products, error) {
	product, err := models.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func SelectProductByPage(pageNum int, pageSize int, name string) ([]*models.Products, error) {
	products, err := models.GetProductByPage(pageNum, pageSize, name)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func AddProductSpike(productId int) error {
	// 查询产品信息并放入redis
	product, err := models.GetProductById(productId)
	if err != nil {
		return err
	}
	key := cacheService.GetProductKey(productId)
	err = redis.SetData(key, product)
	if err != nil {
		return err
	}
	return nil
}
