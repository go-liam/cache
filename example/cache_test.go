package example

import (
	"github.com/go-liam/cache"
	"github.com/go-liam/cache/redis"
	"log"
	"testing"

	_ "github.com/gs-mblock/mbgo/lib/cache/redis"
)


/*
新建对象
=== RUN   TestRedis_NewServer
2020/04/30 09:25:21 [INFO] redis conn ok
2020/04/30 09:25:21 false-IsExist= false
2020/04/30 09:25:21 v= xxxx-xxx
2020/04/30 09:25:21 v2= [120 120 120 120 45 120 120 120]
2020/04/30 09:25:21 true IsExist= true
2020/04/30 09:25:21 false-IsExist= false
2020/04/30 09:25:21 [INFO] redis conn ok
--- PASS: TestRedis_NewServer (0.01s)
*/
func TestRedis_NewServer(t *testing.T) {
	server2 =  new(redis.SvRedis)
	server2.NewCache("Liam123@localhost:6379","dev|")
	n := "key_u1"
	f := server2.IsExist(n)
	log.Println("false-IsExist=", f)
	server2.Set(n, "xxxx-xxx", 10)
	v, _ := server2.Get(n)
	log.Println("v=", v)
	v2, _ := server2.GetBytes(n)
	log.Println("v2=", v2)
	f2 := server2.IsExist(n)
	log.Println("true IsExist=", f2)
	server2.Delete(n)
	f3 := server2.IsExist(n)
	log.Println("false-IsExist=", f3)
	server2.Reconnect()
}

// 默认对象
func TestRedis_Server(t *testing.T) {
	n := "key_u2"
	f := cache.IsExist(n)
	log.Println("false-IsExist=", f)
	cache.Set(n, "xxxx-xxx", 10)
	v, _ := cache.Get(n)
	log.Println("v=", v)
	f2 := cache.IsExist(n)
	log.Println("true-IsExist=", f2)
	cache.Delete(n)
	f3 := cache.IsExist(n)
	log.Println("false-IsExist=", f3)
}
