package redis

import (
	"fmt"
	"github.com/gs-mblock/mbgo/lib/cache"
	_ "github.com/gs-mblock/mbgo/lib/cache/redis"
	"log"
	"testing"
)

func TestPkgRedis_Cache(t *testing.T) {
	//bm, err := cache.NewCache("redis", `{"conn": "Liam123@r-wz9a0dfcef584e14pd.redis.rds.aliyuncs.com:6379"}`)
	bm, err := cache.NewCache("redis", `{"conn": "Liam123@localhost:6379"}`)
	if err != nil {
		fmt.Println("err:", err)
		//t.Error("init err")
	}
	log.Printf("ok: %+v\n", bm)
}

func TestPkgRedis_AutoConnect(t *testing.T) {
	sv := new(SvRedis)
	sv.URL = "Liam123@localhost:6379"
	sv.AutoConnect()
}

func TestPkgRedis_ConnRedis(t *testing.T) {
	sv := new(SvRedis)
	sv.URL = "Liam123@localhost:6379"
	v := sv.Reconnect()
	println("result:", v)
}

func TestPKG_Redis_data(t *testing.T) {
	sv := new(SvRedis)
	sv.URL = "Liam123@localhost:6379"
	sv.Reconnect()
	n := "parent-1|heZi_u1"
	f := sv.IsExist(n)
	log.Println("IsExist=", f)
	sv.Set(n, "xxxx-xxx", 100)
	v, _ := sv.Get(n)
	log.Println("v=", v)
	f2 := sv.IsExist(n)
	log.Println("IsExist=", f2)
	sv.Delete(n)
	f3 := sv.IsExist(n)
	log.Println("IsExist=", f3)
}

func TestPkg_RedisByte(t *testing.T) {
	sv := new(SvRedis)
	sv.URL = "Liam123@localhost:6379"
	sv.Reconnect()
	n := "parent-1|heZi_u1"
	sv.Set(n, []byte("123456-xxx"), 100)
	v, err := sv.GetBytes(n)
	log.Printf("err=:%+v\n", err)
	log.Printf("v=:%+v\n", string(v))
}

/*
12259	     95130 ns/op = 10511 qps/s
13335	     85909 ns/op = 11640 qps/s
*/
//
//func BenchmarkLoopsSv_IsExist(b *testing.B) {
//	n := "p1:"
//	//Server.Set(n, "xx-xx-xxx", 10000)
//	//v, _ := Server.Get(n)
//	//log.Println("v=", v)
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			Server.IsExist(n)
//		}
//	})
//}

/*
11394	     92403 ns/op = 10822 qps/s
*/
//func BenchmarkLoopsSv_Get(b *testing.B) {
//	n := "p1:"
//	//Server.Set(n, "xx-xx-xxx", 10000)
//	//v, _ := Server.Get(n)
//	//log.Println("v=", v)
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			Server.Get(n)
//		}
//	})
//}