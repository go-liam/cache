package redis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	c "github.com/gs-mblock/mbgo/lib/cache"
	"log"
	"time"
	// _ "github.com/gs-mblock/mbgo/lib/cache/redis"
	_ "github.com/gs-mblock/mbgo/lib/cache/redis"
)

// 公共用

const connectTime = 600 //秒

// SvRedis :
type SvRedis struct {
	ModelRedis
}

type ModelRedis struct {
	URL         string // "Liam123@localhost:6379"
	adapter     c.Cache
	connectTime int64
	IsConnect   bool   // 1 ok
	RedisPrefix string // 前缀
}

func (sv *SvRedis) NewCache(url string, redisPrefix string) bool {
	sv.URL = url
	sv.RedisPrefix = redisPrefix
	return sv.Reconnect()
}

// AutoConnect :
func (sv *SvRedis) AutoConnect() {
	log.Println("Redis-AutoConnect-1:", sv.IsConnect)
	if sv.IsConnect {
		return
	}
	now := time.Now().Unix()
	// 10*60 秒
	if now-sv.connectTime < connectTime {
		return
	}
	log.Println("Redis-AutoConnect-2")
	sv.Reconnect()
}

// Reconnect :
func (sv *SvRedis) Reconnect() bool {
	sv.connectTime = time.Now().Unix()
	var err error
	connURL := fmt.Sprintf(`{"conn": "%s"}`, sv.URL)
	//adapter, err = c.NewCache("redis", `{"conn": "Liam123@localhost:6379"}`)
	sv.adapter, err = c.NewCache("redis", connURL)
	if err != nil {
		log.Println("[ERROR] redis conn err :", err)
		sv.IsConnect = false
		return false
	} else {
		log.Println("[INFO] redis conn ok")
	}
	sv.IsConnect = true
	return true
}

// 综上来说，存byte的方式要优于存string。

// Delete :
func (sv *SvRedis) Delete(key string) {
	if !sv.IsConnect {
		sv.AutoConnect()
		return
	}
	sv.adapter.Delete(key)
}

// Set :保存多少 秒 : value:string, []byte
func (sv *SvRedis) Set(key string, value interface{}, second int) bool {
	if !sv.IsConnect {
		sv.AutoConnect()
		return false
	}
	timeoutDuration := time.Duration(second) * time.Second
	//var err error
	if err := sv.adapter.Put(key, value, timeoutDuration); err != nil {
		log.Println("[ERROR] Redis Set err:", err)
		return false
	}
	return true
}

// Get :
func (sv *SvRedis) Get(key string) (string, error) {
	if !sv.IsConnect {
		sv.AutoConnect()
		return "", errors.New("redis conn error")
	}
	v, err := redis.String(sv.adapter.Get(key), nil)
	if err != nil {
		fmt.Println("Redis Get err:", err)
		return "", err
	}
	return v, nil
}

// GetBytes :
func (sv *SvRedis) GetBytes(key string) ([]byte, error) {
	if !sv.IsConnect {
		sv.AutoConnect()
		return nil, errors.New("redis conn error")
	}
	if !sv.IsExist(key) {
		return nil, errors.New("NoExist:" + key)
	}
	//fmt.Println("Redis Get n:", n)
	v, err := redis.Bytes(sv.adapter.Get(key), nil)
	if err != nil {
		fmt.Println("Redis Get err:", err)
		return nil, err
	}
	return v, nil
}

// IsExist :
func (sv *SvRedis) IsExist(key string) bool {
	if !sv.IsConnect {
		sv.AutoConnect()
		return false
	}
	if !sv.adapter.IsExist(key) {
		return false
	}
	return true
}

func (sv *SvRedis) GetPrefix() string {
	return sv.RedisPrefix
}

func (sv *SvRedis) GetURL() string{
	return sv.URL
}