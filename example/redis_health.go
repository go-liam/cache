package example

import (
	"log"
	"time"
)

func Health() string {
	log.Println("[INFO] redis url=", Server.GetUrl())
	log.Println("[INFO] redis RedisPrefix=", Server.GetPrefix())
	n := Server.GetPrefix()+"|health|0"
	Server.Set(n, time.Now().Unix(), 2)
	f,err := Server.Get(n)
	if err != nil{
		log.Printf("[ERROR] redis health:%+v\n",err)
		Server.Reconnect() //重新连接
	}
	return f
}
