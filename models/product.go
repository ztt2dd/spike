package models

import "github.com/jinzhu/gorm"

type Products struct {
	ID          int    `json:"id"`
	ProductName string `json:"productName"`
	ProductNum  int    `json:"productNum"`
	CreateAt    int    `json:"createAt"`
	ModifyAt    int    `json:"modifyAt"`
}

// 新增商品
func AddProduct(data map[string]interface{}) error {
	err := db.Create(&Products{
		ProductName: data["productName"].(string),
		ProductNum:  data["productNum"].(int),
	}).Error
	if err != nil {
		return err
	}

	return nil
}

// 修改商品
func UpdateProduct(id int, data interface{}) error {
	err := db.Model(&Products{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// 根据ID获取商品
func GetProductById(id int) (*Products, error) {
	var product Products
	err := db.Where("id = ? ", id).First(&product).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &product, nil
}

// 分页获取商品数据
func GetProductByPage(pageNum int, pageSize int, productName string) ([]*Products, error) {
	var products []*Products
	if productName != "" {
		productName = "'%" + productName + "%'"
		db = db.Where("product_name like ?", productName)
	}
	db = db.Offset(pageNum).Limit(pageSize)
	err := db.Find(&products).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return products, nil
}

// 获取商品总数
func GetProductTotal(maps interface{}) (count int) {
	db.Model(&Products{}).Where(maps).Count(&count)
	return
}

// 通过热更新更新产品的库存信息
func UpdateProductByVersion(productId int, productNum int, data map[string]interface{}) (int64, error) {
	db = db.Model(&Products{}).Where("id = ? and product_num = ?", productId, productNum).Updates(data)
	err := db.Error
	if err != nil {
		return 0, err
	}
	count := db.RowsAffected
	return count, nil
}
