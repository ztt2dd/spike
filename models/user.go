package models

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	CreateAt int    `json:"createAt"`
	ModifyAt int    `json:"modifyAt"`
}

func AddUser(name string, password string) error {
	err := db.Create(&User{
		Name:     name,
		Password: password,
	}).Error
	return err
}

func CheckUserAuth(name, password string) (int, error) {
	var user User
	err := db.Select("id").Where(User{Name: name, Password: password}).First(&user).Error
	if err != nil {
		return 0, err
	}
	if user.ID > 0 {
		return user.ID, nil
	}

	return 0, nil
}

func GetUserById(id int) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
