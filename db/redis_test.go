package db

import (
	"encoding/json"
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
	var tt interface{}

	f := func() (a interface{}, err error) {
		a = &struct {
			Name string
			Age  int
		}{
			"cz", 11,
		}

		return
	}

	err = dao.GetCache("666", &tt, 1000, true, f)
	bs, err := json.Marshal(tt)

	t.Logf("over %s	%v", bs, err)
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

	dao.DelayConsume("testDelay", 2, func() int64 {
		return 10
	}, func(task interface{}, ctime int64) error {
		log.Printf("=================== %v", task)

		return fmt.Errorf("test")
	}, 3, func(err error) {
		fmt.Println("alert", err)
	})

}

func TestDelayAdd(t *testing.T) {

	dao, err := NewRedisDao("redis://@127.0.0.1:6379/0?idle=100&active=1000&wait=true&timeout=3s", true)
	if err != nil {
		panic(err)
	}

	t.Log(dao.DelayAdd("testDelay", 100, 0))
	t.Log(dao.DelayAdd("testDelay", 200, 1111111))
}

func TestStringSet(t *testing.T) {
	dao, err := NewRedisDao("redis://@127.0.0.1:6379/0?idle=100&active=1000&wait=true&timeout=3s", true)
	if err != nil {
		panic(err)
	}
	t.Log(dao.Set("k:1", "v1"))
	t.Log(dao.Set("k:2", "v2"))

	ks, err := dao.KEYS("k*")
	//ks, err := dao.KEYS("o*")

	rs, err := dao.MGet(ks)
	t.Log(err)
	t.Log(rs)
}

func BenchmarkMGet(b *testing.B) {
	dao, err := NewRedisDao("redis://@127.0.0.1:6379/0?idle=100&active=1000&wait=true&timeout=3s", true)
	if err != nil {
		panic(err)
	}
	b.ReportAllocs()
	dao.Set("k:1", "v1")
	dao.Set("k:2", "v2")

	ks, err := dao.KEYS("k*")

	for i := 0; i < b.N; i++ {
		b.Log(dao.MGet(ks))
	}
}
