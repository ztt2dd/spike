package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"

	"spikeKill/pkg/setting"
)

var RedisConn *redis.Pool
var ctx = context.Background()
var mutex sync.Mutex

// Setup 通过redis连接池的方式
func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

// 保存数据（有效期为3小时）
func SetData(key string, data interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, 24*3600)
	if err != nil {
		return err
	}

	return nil
}

// 数据是否已存在
func isExist(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// 获取数据
func GetData(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// 删除数据
func DelData(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// 通过redis加分布式锁
func Lock(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	ts := time.Now() // 设置一个随机事件值
	v, err := conn.Do("SET", key, ts, "EX", 1, "NX")
	if err != nil {
		return err
	}
	if v != nil {
		return nil
	} else {
		err = fmt.Errorf("get lock failed")
	}

	return err
}

// 通过redis解锁
func UnLock(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := redis.Bool(conn.Do("DEL", key))
	if err != nil {
		return err
	}
	return nil
}
