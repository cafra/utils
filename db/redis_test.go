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
