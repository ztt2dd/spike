package models

type Users struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	CreateAt int    `json:"createAt"`
	ModifyAt int    `json:"modifyAt"`
}

// 新增用户
func AddUser(name string, password string) error {
	err := db.Create(&Users{
		Name:     name,
		Password: password,
	}).Error
	return err
}

// 用户登录验证
func CheckUserAuth(name, password string) (int, error) {
	var user Users
	err := db.Select("id").Where(Users{Name: name, Password: password}).First(&user).Error
	if err != nil {
		return 0, err
	}
	if user.ID > 0 {
		return user.ID, nil
	}

	return 0, nil
}

// 根据id获取用户
func GetUserById(id int) (*Users, error) {
	var user Users
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
