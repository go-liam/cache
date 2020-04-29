package cache

type InCache interface {
	Delete(key string)
	Set(key string, value interface{}, second int) bool
	Get(key string) (string, error)
	GetBytes(key string) ([]byte, error)
	IsExist(key string) bool
	GetPrefix() string
	GetUrl() string
	NewCache(url string ,redisPrefix string) bool
	Reconnect() bool
}
