package service

import (
	"fmt"
	"github.com/go-liam/cache"
	"log"
)

// 业务用的

func init() {
	//Server = new(redis.SvRedis)
	url := fmt.Sprintf("%s@%s:%s", RedisConfig.RedisPwd, RedisConfig.RedisHost, RedisConfig.RedisPort)
	//"Liam123@localhost:6379"
	redisPrefix := RedisConfig.RedisPrefix
	v := cache.NewCache(url,redisPrefix)
	if v {
		log.Printf("[INFO] Redis connect scuess")
	} else {
		log.Printf("[ERROR]Redis connect failse")
	}
	log.Println("[INFO] redis url=", url)
	log.Println("[INFO] redis RedisPrefix=", redisPrefix)
}
