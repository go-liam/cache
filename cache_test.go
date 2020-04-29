package cache

import (
	"log"
	"testing"

	"github.com/gs-mblock/mbgo/lib/cache"
	_ "github.com/gs-mblock/mbgo/lib/cache/redis"
)

func TestRedis_data(t *testing.T) {
	NewCache("Liam123@localhost:6379","dev|")
	n := "heZi_u1"
	f := IsExist(n)
	log.Println("IsExist=", f)
	Set(n, "xxxx-xxx", 1)
	v, _ := Get(n)
	log.Println("v=", v)
	f2 := IsExist(n)
	log.Println("IsExist=", f2)
	Delete(n)
	f3 := IsExist(n)
	log.Println("IsExist=", f3)
}

func TestRedis_Cache(t *testing.T) {
	//bm, err := cache.NewCache("redis", `{"conn": "Liam123@r-wz9a0dfcef584e14pd.redis.rds.aliyuncs.com:6379"}`)
	bm, err := cache.NewCache("redis", `{"conn": "Liam123@localhost:6379"}`)
	if err != nil {
		log.Println(err)
		//t.Error("init err")
	}
	println((bm))
}

func TestRedisByte(t *testing.T) {
	NewCache("Liam123@localhost:6379","dev|")
	GetBytes("aa")
}

func TestC_GetRedisPrefix(t*testing.T)  {
	log.Println(GetRedisPrefix())
	log.Println(GetRedisPrefix())
}

func TestGetUrl(t *testing.T) {
	println(GetUrl())
}

func TestNewConnect(t *testing.T) {
	NewConnect()
}