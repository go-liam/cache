package cache

import (
	"github.com/go-liam/cache/redis"
)

var server InCache

func NewCache(url string ,redisPrefix string) bool {
	server = new(redis.SvRedis)
	return server.NewCache(url,redisPrefix)
}

// Delete :
func Delete(key string) {
	n := GetRedisPrefix() + key
	server.Delete(n)
}

// Set :保存多少s: value:string, []byte
func Set(key string, value interface{}, second int) bool {
	n := GetRedisPrefix()+ key
	return server.Set(n,value,second )
}

// Get :
func Get(key string) (string, error) {
	n := GetRedisPrefix() + key
	return  server.Get(n)
}

// GetBytes :
func GetBytes(key string) ([]byte, error) {
	n := GetRedisPrefix() + key
	return server.GetBytes(n)
}

// IsExist :
func IsExist(key string) bool {
	n := GetRedisPrefix() + key
	return server.IsExist(n)
}

func GetRedisPrefix() string {
	return server.GetPrefix()
}

func GetUrl() string{
	return server.GetUrl()
}

func NewConnect() bool  {
	return server.Reconnect()
}