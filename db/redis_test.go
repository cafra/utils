package db

import (
	"fmt"
	"log"

	//"github.com/garyburd/redigo/redis"
	"testing"
)

//URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
//func TestNewRedis(t *testing.T) {
//	//t.Log(NewRedis("redis://user:pass@127.0.0.1:6379/db?idle=100&active=100&wait=true&timeout=3s"))
//	dao, err := NewRedisDao("redis://@127.0.0.1:6379/1?idle=100&active=100&wait=true&timeout=3s", true)
//	if err != nil {
//		t.Log(err)
//		t.FailNow()
//	}
//	t.Log(dao, err)
//
//	t.Log(dao.Set("test", "testvalue"))
//	t.Log(dao.Get("test"))
//
//	_, err = dao.Get("test2")
//	t.Log(err, Nil, err == Nil)
//}

//func TestBRPOP(t *testing.T) {
//	dao, err := NewRedisDao("redis://@127.0.0.1:6379/0?idle=100&active=100&wait=true&timeout=3s", true)
//
//	fmt.Println(err)
//	for {
//
//		con := dao.RedisPool().Get()
//		//fmt.Println(con.Do("BRPOP", "test1", "test2", 0))
//		data, err := redis.Strings(con.Do("BRPOP", "test1", "test2", 0))
//		fmt.Printf("=======%v	%v\n", data, err)
//	}
//
//}

func TestGetCache(t *testing.T) {
	dao, err := NewRedisDao("redis://@127.0.0.1:6379/0?idle=100&active=100&wait=true&timeout=3s", true)
	if err != nil {
		panic(err)
	}
	tt := struct {
		Name string
		Age  int
	}{}

	f := func() (a interface{}, err error) {
		a = &struct {
			Name string
			Age  int
		}{
			"cz", 11,
		}

		return
	}

	err = dao.GetCache("666", tt, 1000, true, f)
	t.Logf("over %v	%v", tt, err)
}

var f = func() (a interface{}, err error) {
	a = &struct {
		Name string
		Age  int
	}{
		"cz", 11,
	}

	return
}

func BenchmarkGetCache(b *testing.B) {
	b.ReportAllocs()
	dao, err := NewRedisDao("redis://@127.0.0.1:6379/0?idle=100&active=1000&wait=true&timeout=3s", true)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		var tt = new(struct {
			Name string
			Age  int
		})
		key := fmt.Sprint(i)
		err = dao.GetCache(key, tt, 0, false, f)
		b.Log(tt)
	}
}
func BenchmarkGetXXX(b *testing.B) {
	b.Logf("111")
	b.ReportAllocs()
	dao, err := NewRedisDao("redis://@localhost:6379/0?idle=1000&active=1000&wait=true&timeout=10s", false)
	if err != nil {
		panic(err)
	}
	b.Logf("111")

	for i := 0; i < b.N; i++ {
		_, err = dao.GetBytes(fmt.Sprint(i))
	}

	//todo 并发10000 pprof 分析原因
}

func TestZRANGEDelayTask(t *testing.T) {
	dao, err := NewRedisDao("redis://@127.0.0.1:6379/0?idle=100&active=1000&wait=true&timeout=3s", true)
	if err != nil {
		panic(err)
	}
	t.Log(dao.ZRANGEDelayTask("testDelay", 1993886989))
}

func TestDelayConsume(t *testing.T) {
	dao, err := NewRedisDao("redis://@127.0.0.1:6379/0?idle=100&active=1000&wait=true&timeout=3s", true)
	if err != nil {
		panic(err)
	}

	t.Log(dao.DelayConsume("testDelay", 10, func() int64 {
		return 10
	}, func(task string) error {
		log.Printf("=================== %v", task)
		return nil
	}))
	t.Log("======")
}
