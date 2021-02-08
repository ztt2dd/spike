package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Orders struct {
	ID        int   `json:"id"`
	ProductId int   `json:"productId"`
	UserId    int   `json:"userId"`
	OrderSn   int64 `json:"orderSn"`
	CreateAt  int   `json:"createAt"`
	ModifyAt  int   `json:"modifyAt"`
}

// 通过事务的方式去创建订单
func CreateLocalOrder(order *Orders) error {
	product, err := GetProductById(order.ProductId)
	if err != nil {
		return err
	}
	db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		// 新增订单信息
		if err := tx.Create(order).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		// 通过乐观锁的方式修改商品的库存信息
		tx = tx.Model(&Products{}).Where("id = ? and product_num = ?", order.ProductId, product.ProductNum).Updates(map[string]interface{}{"product_num": product.ProductNum - 1})
		err := tx.Error
		if err != nil {
			return err
		}
		count := tx.RowsAffected
		if count <= 0 { // 没有更新成功
			return fmt.Errorf("update product stock fail")
		}

		// 返回 nil 提交事务
		return nil
	})
	return err
}
