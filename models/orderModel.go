package models

type Order struct {
	ID        int `json:"id"`
	ProductId int `json:"productId"`
	UserId    int `json:"userId"`
}

func AddOrder(data map[string]interface{}) error {
	err := db.Create(&Order{
		ProductId: data["productId"].(int),
		UserId:    data["userId"].(int),
	}).Error
	if err != nil {
		return err
	}
	return nil
}
