package cache

import (
	"github.com/go-liam/cache/redis"
)

var Sv InCache

func NewCache(url string ,redisPrefix string) bool {
	Sv = new(redis.SvRedis)
	return Sv.NewCache(url, redisPrefix)
}

// Delete :
func Delete(key string) {
	n := GetPrefix() + key
	Sv.Delete(n)
}

// Set :保存多少s: value:string, []byte
func Set(key string, value interface{}, second int) bool {
	n := GetPrefix() + key
	return Sv.Set(n, value, second)
}

// Get :
func Get(key string) (string, error) {
	n := GetPrefix() + key
	return Sv.Get(n)
}

// GetBytes :
func GetBytes(key string) ([]byte, error) {
	n := GetPrefix() + key
	return Sv.GetBytes(n)
}

// IsExist :
func IsExist(key string) bool {
	n := GetPrefix() + key
	return Sv.IsExist(n)
}

func GetPrefix() string {
	return Sv.GetPrefix()
}

func GetURL() string {
	return Sv.GetURL()
}

func Reconnect() bool {
	return Sv.Reconnect()
}
