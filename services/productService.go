package services

import (
	"spikeKill/models"
)

// 商品相关定义接口
type IProductService interface {
	// 新增商品
	AddProduct(data map[string]interface{}) (int, error)
	// 修改商品信息
	UpdateProduct(id int, data map[string]interface{}) (int, error)
	// 根据id查询商品
	SelectProductById(int int) *models.Product
	// 分页查询商品信息
	SelectProductByPage(pageNum int, pageSize int, name string) []*models.Product
}

type ProductService struct{}

func (productService *ProductService) AddProduct(data map[string]interface{}) (int, error) {
	// 通过model层实现数据库操作
	err := models.AddProduct(data)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (productService *ProductService) UpdateProduct(id int, data map[string]interface{}) (int, error) {
	err := models.UpdateProduct(id, data)
	if err != nil {
		return 0, nil
	}
	return 1, nil
}

func (productService *ProductService) SelectProductById(id int) (*models.Product, error) {
	product, err := models.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (productService *ProductService) SelectProductByPage(pageNum int, pageSize int, name string) ([]*models.Product, error) {
	products, err := models.GetProductByPage(pageNum, pageSize, name)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (productService *ProductService) UpdateProductNum(productId int, productNum int) (int64, error) {
	data := make(map[string]interface{})
	data["product_num"] = productNum - 1
	count, err := models.UpdateProductByVersion(productId, productNum, data)
	return count, err
}
