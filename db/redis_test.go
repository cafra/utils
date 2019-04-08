package db

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
)

//URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
func TestNewRedis(t *testing.T) {
	//t.Log(NewRedis("redis://user:pass@127.0.0.1:6379/db?idle=100&active=100&wait=true&timeout=3s"))
	dao, err := NewRedisDao("redis://@127.0.0.1:6379/1?idle=100&active=100&wait=true&timeout=3s", true)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(dao, err)

	t.Log(dao.Set("test", "testvalue"))
	t.Log(dao.Get("test"))

	_, err = dao.Get("test2")
	t.Log(err, Nil, err == Nil)
}

func TestBRPOP(t *testing.T) {
	dao, err := NewRedisDao("redis://@127.0.0.1:6379/0?idle=100&active=100&wait=true&timeout=3s", true)

	fmt.Println(err)
	for {

		con := dao.RedisPool().Get()
		//fmt.Println(con.Do("BRPOP", "test1", "test2", 0))
		data, err := redis.Strings(con.Do("BRPOP", "test1", "test2", 0))
		fmt.Printf("=======%v	%v\n", data, err)
	}

}

func TestGetCache(t *testing.T) {
	dao, err := NewRedisDao("redis://@127.0.0.1:6379/0?idle=100&active=100&wait=true&timeout=3s", true)
	if err != nil {
		panic(err)
	}
	var tt = new(struct {
		Name string
		Age  int
	})

	f := func() (a interface{}, err error) {
		a = &struct {
			Name string
			Age  int
		}{
			"cz", 11,
		}

		return
	}

	err = dao.GetCache("5555", 1, tt, f)
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
	dao, err := NewRedisDao("redis://@127.0.0.1:6379/0?idle=100&active=100&wait=true&timeout=3s", true)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		var tt = new(struct {
			Name string
			Age  int
		})
		err = dao.GetCache("i", 10000, tt, f)
		b.Log(tt)
	}
}
