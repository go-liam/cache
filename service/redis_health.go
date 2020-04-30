package service

import (
	"github.com/go-liam/cache"
	"log"
	"time"
)

func Health() string {
	log.Println("[INFO] redis url=", cache.GetURL())
	log.Println("[INFO] redis RedisPrefix=", cache.GetPrefix())
	n := cache.GetPrefix()+"|health|0"
	cache.Set(n, time.Now().Unix(), 2)
	f,err := cache.Get(n)
	if err != nil{
		log.Printf("[ERROR] redis health:%+v\n",err)
		cache.Reconnect() //重新连接
	}
	return f
}
