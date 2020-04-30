package service

import (
	"log"
	"time"
)

func Health() string {
	log.Println("[INFO] redis url=", sv.GetURL())
	log.Println("[INFO] redis RedisPrefix=", sv.GetPrefix())
	n := sv.GetPrefix()+"|health|0"
	sv.Set(n, time.Now().Unix(), 2)
	f,err := sv.Get(n)
	if err != nil{
		log.Printf("[ERROR] redis health:%+v\n",err)
		sv.Reconnect() //重新连接
	}
	return f
}
