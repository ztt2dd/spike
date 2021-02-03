package services

import (
	"encoding/json"
	"log"
	"spikeKill/models"
	"spikeKill/pkg/redis"
	"spikeKill/pkg/util"
	"spikeKill/services/cacheService"
)

type User struct {
	User  *models.User
	Token string `json:"token"`
}

type IUserService interface {
	// 新增用户
	AddUser(name string, password string) (int, error)
	// 用户授权登录
	GetAuth(name string, password string) (*models.User, error)
	// 获取用户信息
	GetUserInfo(id int) (*models.User, error)
}

type UserService struct {
}

func (u *UserService) AddUser(name string, password string) (int, error) {
	err := models.AddUser(name, password)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (u *UserService) GetAuth(name string, password string) (*User, error) {
	id, err := models.CheckUserAuth(name, password)
	if err != nil {
		return nil, err
	}
	if id > 0 {
		token, err := util.GenerateToken(name, password)
		if err != nil {
			log.Printf("Token生成失败：%s", name)
			return nil, err
		}
		userData, err := models.GetUserById(id)
		if err != nil {
			log.Printf("获取用户数据失败：%s", name)
		}
		user := &User{
			User:  userData,
			Token: token,
		}
		// 将数据写入redis
		key := cacheService.GetUserKey(id)
		err = redis.SetData(key, user)
		if err != nil {
			log.Printf("用户数据写入redis失败：%s", name)
			log.Println(err)
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

func (u *UserService) GetUserInfo(id int) (*models.User, error) {
	key := cacheService.GetUserKey(id)
	var user *models.User
	userRData, err := redis.GetData(key)
	if err != nil {
		log.Printf("用户数据从redis获取失败：%d", id)
		return nil, err
	}
	if userRData != nil {
		json.Unmarshal(userRData, &user)
	} else {
		user, err = models.GetUserById(id)
		if err != nil {
			return nil, err
		}
		redis.SetData(key, user)
	}

	return user, nil
}
