package service

import (
	"github.com/go-liam/cache"
	"github.com/go-liam/cache/redis"
)

var sv cache.InCache

func NewCache(url string ,redisPrefix string) bool {
	sv = new(redis.SvRedis)
	return sv.NewCache(url,redisPrefix)
}

// Delete :
func Delete(key string) {
	n := GetRedisPrefix() + key
	sv.Delete(n)
}

// Set :保存多少s: value:string, []byte
func Set(key string, value interface{}, second int) bool {
	n := GetRedisPrefix()+ key
	return sv.Set(n,value,second )
}

// Get :
func Get(key string) (string, error) {
	n := GetRedisPrefix() + key
	return  sv.Get(n)
}

// GetBytes :
func GetBytes(key string) ([]byte, error) {
	n := GetRedisPrefix() + key
	return sv.GetBytes(n)
}

// IsExist :
func IsExist(key string) bool {
	n := GetRedisPrefix() + key
	return sv.IsExist(n)
}

func GetRedisPrefix() string {
	return sv.GetPrefix()
}

func GetUrl() string{
	return sv.GetURL()
}

func NewConnect() bool  {
	return sv.Reconnect()
}