package db

import (
	"errors"
	"log"
	"net/url"
	"time"
)
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/schema"
	"strconv"
	"strings"
)

const _NIL = "redigo: nil returned"

var Nil = errors.New(_NIL)

type RedisDao struct {
	debug     bool
	redisPool *redis.Pool
}

type _REDIS_CFG struct {
	MaxIdle     int           `schema:"idle"`
	MaxActive   int           `schema:"active"`
	Wait        bool          `schema:"wait"`
	IdleTimeout string        `schema:"timeout"`
	_time_out   time.Duration `schema:"-"`
}

func NewRedisDao(cfgStr string, debug bool) (dao *RedisDao, err error) {
	dao = &RedisDao{debug: debug}
	decoder := schema.NewDecoder()
	cfg := new(_REDIS_CFG)

	u, err := url.Parse(cfgStr)
	if err != nil {
		return nil, err
	}

	db, err := getDbId(u.Path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("db id=%v invalid", db))
	}

	values, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}

	if err = decoder.Decode(cfg, values); err != nil {
		return
	}

	if cfg._time_out, err = time.ParseDuration(cfg.IdleTimeout); err != nil {
		return nil, err
	}
	dao.Debug("cfg=%#v", *cfg)
	dao.redisPool = &redis.Pool{
		MaxIdle:     cfg.MaxIdle,
		MaxActive:   cfg.MaxActive,
		Wait:        cfg.Wait,
		IdleTimeout: cfg._time_out,
		Dial: func() (c redis.Conn, err error) {
			if c, err = redis.Dial("tcp", u.Host); err != nil {
				dao.Debug("Dial tcp host=%v err=%v", u.Host, err)
				return
			}

			pass, ok := u.User.Password()
			if ok {
				if _, err = c.Do("AUTH", pass); err != nil {
					dao.Debug("AUTH pass=%v err=%v", pass, err)
					c.Close()
					return
				}
			}

			if db > 0 {
				if _, err = c.Do("SELECT", db); err != nil {
					dao.Debug("SELECT db=%v err=%v", db, err)
					c.Close()
					return
				}
			}
			return
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return
}

func getDbId(a string) (int, error) {
	str := strings.TrimPrefix(a, "/")
	if len(str) == 0 {
		return 0, nil
	}
	return strconv.Atoi(str)
}

func (this *RedisDao) Debug(format string, args ...interface{}) {
	if !this.debug {
		return
	}
	log.Printf(format, args)
}

func (this *RedisDao) RedisPool() *redis.Pool {
	return this.redisPool
}

//基础redis操作
func (this *RedisDao) TTL(key string) (ttl int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	ttl, err = redis.Int(conn.Do("TTL", key))
	if err != nil {
		return
	}
	return
}

//集合操作
//SADD 可以添加多个 返回成功数量
func (this *RedisDao) SADD(key string, value interface{}) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("SADD", key, value))
	if err != nil {
		return
	}
	return
}

// Set 总是成功的
func (this *RedisDao) Set(key string, value interface{}) (err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return
	}
	return
}

//不存在则设置，存在则不设置
func (this *RedisDao) SetNX(key string, value interface{}) (err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	_, err = conn.Do("SET", key, value, "NX")
	if err != nil {
		return
	}
	return
}
func (this *RedisDao) SetNX2(key string, value interface{}) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("SETNX", key, value))
	if err != nil {
		return
	}
	return
}

func (this *RedisDao) SetNX3(key string, value interface{},sec int) (bt bool, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.String(conn.Do("SET", key, value,"ex",sec,"nx"))
	if err.Error() == _NIL {
		return bt,nil
	}
	return
}

// Del 可以删除多个key 返回删除key的num和错误
func (this *RedisDao) Del(key ...interface{}) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("DEL", key...))
	if err != nil {
		return
	}
	return
}

//Get
func (this *RedisDao) Get(key string) (s string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	s, err = redis.String(conn.Do("GET", key))
	if err != nil {
		return
	}
	return
}

//GetBytes
func (this *RedisDao) GetBytes(key string) (bs []byte, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	bs, err = redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return
	}
	return
}

//GetInt
func (this *RedisDao) GetInt(key string) (n int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	n, err = redis.Int(conn.Do("GET", key))
	if err != nil {
		return
	}
	return
}
func (this *RedisDao) GetInt64(key string) (n int64, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	n, err = redis.Int64(conn.Do("GET", key))
	if err != nil {
		return
	}
	return
}

//EXIST
func (this *RedisDao) EXISTS(key string) (ok bool, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	ok, err = redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return
	}
	return
}

//KEYS cz
func (this *RedisDao) KEYS(pattern string) (keys []string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	keys, err = redis.Strings(conn.Do("KEYS", pattern))
	if err != nil {
		return
	}
	return
}

//SCARD
func (this *RedisDao) SCARD(key string) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("SCARD", key))
	if err != nil {
		return
	}
	return
}

//SPOP 弹出被移除的元素, 当key不存在的时候返回 nil
func (this *RedisDao) SPOP(key string) (out string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	out, err = redis.String(conn.Do("SPOP", key))
	if err != nil {
		return
	}
	return
}

//SREM
func (this *RedisDao) SREM(key string, value interface{}) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("SREM", key, value))
	if err != nil {
		return
	}
	return
}

//SISMEMBER
func (this *RedisDao) SISMEMBER(key string, value interface{}) (ok bool, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	ok, err = redis.Bool(conn.Do("SISMEMBER", key, value))
	if err != nil {
		return
	}
	return
}

//SMEMBERS
func (this *RedisDao) SMEMBERS(key string) (reply []string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	reply, err = redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		return
	}
	return
}

//List操作
//LLEN
//LPOP
func (this *RedisDao) LPOP(key string) (out string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	out, err = redis.String(conn.Do("LPOP", key))
	if err != nil {
		return
	}
	return
}
func (this *RedisDao) LLEN(key string) (len int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	len, err = redis.Int(conn.Do("LLEN", key))
	if err != nil {
		return
	}
	return
}
func (this *RedisDao) RPOP(key string) (out string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	out, err = redis.String(conn.Do("RPOP", key))
	if err != nil {
		return
	}
	return
}
func (this *RedisDao) BRPOP(key string) (out []string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	out, err = redis.Strings(conn.Do("BRPOP", key, 0))
	if err != nil {
		return
	}
	return
}

//LPUSH 整型回复: 在 push 操作后的 list 长度。
func (this *RedisDao) LPUSH(key string, value ...interface{}) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("LPUSH", key, value))
	if err != nil {
		return
	}
	return
}

//LINDEX 当 key 位置的值不是一个列表的时候，会返回一个error
func (this *RedisDao) LINDEX(key string, index int) (out string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	out, err = redis.String(conn.Do("LINDEX", key, index))
	if err != nil {
		return
	}
	return
}

//LTRIM
//LRANGE

//哈希操作
//HDEL
//HEXISTS
func (this *RedisDao) HEXISTS(key, field string) (ok bool, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	ok, err = redis.Bool(conn.Do("HEXISTS", key, field))
	if err != nil {
		return
	}
	return
}

//HGET 该字段所关联的值。当字段不存在或者 key 不存在时返回nil。
func (this *RedisDao) HGET(key, field string) (out string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	out, err = redis.String(conn.Do("HGET", key, field))
	if err != nil {
		return
	}
	return
}

//HINCRBY 增值操作执行后的该字段的值。
func (this *RedisDao) HINCRBY(key, field string, in int) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("HINCRBY", key, field, in))
	if err != nil {
		return
	}
	return
}

//HMGETSTRUCT
func (this *RedisDao) HMGETSTRUCT(key, value interface{}) (err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	v, err := redis.Values(conn.Do("HGETALL", key))
	if err == nil {
		if err = redis.ScanStruct(v, value); err != nil {
			return
		}
	}
	return
}

//HMGETMAP
func (this *RedisDao) HMGETMAP(key string) (m map[string]string, err error) {
	m = make(map[string]string)
	conn := this.redisPool.Get()
	defer conn.Close()
	m, err = redis.StringMap(conn.Do("HGETALL", key))
	return
}
func (this *RedisDao) HMGETINTMAP(key string) (m map[string]int, err error) {
	m = make(map[string]int)
	conn := this.redisPool.Get()
	defer conn.Close()
	m, err = redis.IntMap(conn.Do("HGETALL", key))
	return
}
func (this *RedisDao) HMGETINT64MAP(key string) (m map[string]int64, err error) {
	m = make(map[string]int64)
	conn := this.redisPool.Get()
	defer conn.Close()
	m, err = redis.Int64Map(conn.Do("HGETALL", key))
	return
}

func (r *RedisDao) HGETALLMAP(key string) (interface{}, error) {
	conn := r.redisPool.Get()
	defer conn.Close()
	return conn.Do("HGETALL", key)
}

//HMSET
func (this *RedisDao) HMSET(key, value interface{}) (ok string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	ok, err = redis.String(conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(value)...))
	if err != nil {
		return
	}
	return
}

//HMGET
func (this *RedisDao) HMGET(key, feild string) (data []string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	data, err = redis.Strings(conn.Do("HMGET", key, feild))
	if err != nil {
		return
	}
	return
}

//HKEYS
func (this *RedisDao) HKEYS(key string) (data []string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	data, err = redis.Strings(conn.Do("HKEYS", key))
	if err != nil {
		return
	}
	return
}

//HMGET
func (this *RedisDao) HMGET2(key string, feild ...string) (data []string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	data, err = redis.Strings(conn.Do("HMGET", redis.Args{}.Add(key).AddFlat(feild)...))
	if err != nil {
		return
	}
	return
}

//HSCAN
//HSET 1如果field是一个新的字段  0如果field原来在map里面已经存在
func (this *RedisDao) HSET(key, field string, value interface{}) (ok bool, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	ok, err = redis.Bool(conn.Do("HSET", key, field, value))
	if err != nil {
		return
	}
	return
}

//HLEN 哈希集中字段的数量，当 key 指定的哈希集不存在时返回 0
func (this *RedisDao) HLEN(key string) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("HLEN", key))
	if err != nil {
		return
	}
	return
}

//ZREMRANGEBYRANK myzset 0 1  0 -200(保留200名)
func (this *RedisDao) ZREMRANGEBYRANK(key string, stop int) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("ZREMRANGEBYRANK", key, 0, stop))
	if err != nil {
		return
	}
	return
}

//ZADD
func (this *RedisDao) ZADD(key string, sorce int, member string) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("ZADD", key, sorce, member))
	if err != nil {
		return
	}
	return
}

//ZADD float64
func (this *RedisDao) ZFADD(key string, sorce float64, member string) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("ZADD", key, sorce, member))
	if err != nil {
		return
	}
	return
}

//ZCARD cz
func (this *RedisDao) ZCARD(key string) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("ZCARD", key))
	return
}

//ZRANGE cz
func (this *RedisDao) ZRANGE(key string, start, stop int) (list []string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	list, err = redis.Strings(conn.Do("ZRANGE", key, start, stop))
	if err != nil {
		return
	}
	return
}

//ZREVRANGE cz
func (this *RedisDao) ZREVRANGE(key string, start, stop int) (list []string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	list, err = redis.Strings(conn.Do("ZREVRANGE", key, start, stop))
	if err != nil {
		return
	}
	return
}

// ZSCORE cz
func (this *RedisDao) ZSCORE(key string, member string) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("ZSCORE", key, member))
	return
}

// ZSCORE cz
func (this *RedisDao) ZFSCORE(key string, member string) (num float64, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Float64(conn.Do("ZSCORE", key, member))
	return
}

// ZSCORE cz
func (this *RedisDao) ZREM(key string, member string) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("ZREM", key, member))
	return
}

//ZREVRANGEBYSCORE 逆序份数  获取的 前N个数据
func (this *RedisDao) ZREVRANGEBYSCORE(key string, limit int) (list map[string]string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	list, err = redis.StringMap(conn.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf", "WITHSCORES", "limit", 0, limit))
	if err != nil {
		return
	}
	return
}

//ZREVRANGEBYSCORE 逆序份数  获取start len的数据
func (this *RedisDao) ZREVRANGEBYSCORE2(key string, start, len int) (list map[string]int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	list, err = redis.IntMap(conn.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf", "WITHSCORES", "limit", start, len))
	if err != nil {
		return
	}
	return
}

//ZREVRANGEBYSCORE 逆序份数  获取start len的数据
func (this *RedisDao) ZREVRANGEBYSCORE3(key string, start, len int) (list map[string]float64, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	list, err = FloatMap(conn.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf", "WITHSCORES", "limit", start, len))
	if err != nil {
		return
	}
	return
}

func FloatMap(result interface{}, err error) (map[string]float64, error) {
	values, err := redis.Values(result, err)
	if err != nil {
		return nil, err
	}
	if len(values)%2 != 0 {
		return nil, errors.New("redigo: IntMap expects even number of values result")
	}
	m := make(map[string]float64, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].([]byte)
		if !ok {
			return nil, errors.New("redigo: IntMap key not a bulk string value")
		}
		value, err := redis.Float64(values[i+1], nil)
		if err != nil {
			return nil, err
		}
		m[string(key)] = value
	}
	return m, nil
}

//ZREVRANGE

//ZREVRANGEBYSCORE 逆序份数  获取的 前N个数据 不要scores
func (this *RedisDao) GetSearchKeys(key string, limit int) (list []string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	list, err = redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf", "limit", 0, limit))
	if err != nil {
		return
	}
	return
}

//ZREVRANGEBYSCORE 逆序份数  获取的 start,len 不要scores
func (this *RedisDao) GetSearchKeys2(key string, start, len int) (list []string, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	list, err = redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf", "limit", start, len))
	if err != nil {
		return
	}
	return
}

//ZINCRBY +increment  如果没有key 插入
func (this *RedisDao) ZINCRBY(key string, increment int, member string) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("ZINCRBY", key, increment, member))
	if err != nil {
		return
	}
	return
}

//ZRANK 判断一个member 在key中的索引 如果不在 返回nil ,在 返回索引
func (this *RedisDao) ZRANK(key string, member string) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("ZRANK", key, member))
	if err != nil {
		return
	}
	return
}

//ZREVRANK
func (this *RedisDao) ZREVRANK(key string, member string) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("ZREVRANK", key, member))

	return num, err
}

//EXPIRE 设置一个key 的过期时间 返回值int 1 如果设置了过期时间 0 如果没有设置过期时间，或者不能设置过期时间
func (this *RedisDao) EXPIRE(key string, expireTime int) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("EXPIRE", key, expireTime))
	if err != nil {
		return
	}
	return
}

//EXPIREAT 设置一个key 的在指定时间过期 返回值：如果生存时间设置成功，返回 1 ;当 key 不存在或没办法设置生存时间，返回 0 。

func (this *RedisDao) EXPIREAT(key string, expireAtTime int64) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("EXPIREAT", key, expireAtTime))
	if err != nil {
		return
	}
	return
}

//SETEX key seconds value
func (this *RedisDao) SETEX(key string, seconds int, value interface{}) (err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	_, err = conn.Do("SETEX", key, seconds, value)
	if err != nil {
		return
	}
	return
}

//SETEX key seconds value
func (this *RedisDao) INCR(key string) (err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	_, err = conn.Do("INCR", key)
	if err != nil {
		return
	}
	return
}
func (this *RedisDao) INCRRET(key string) (num int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	num, err = redis.Int(conn.Do("INCR", key))
	if err != nil {
		return
	}
	return
}

func (this *RedisDao) INCRBY(key string,num int64) (err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	_, err = conn.Do("INCRBY", key,num)
	if err != nil {
		return
	}
	return
}

func (this *RedisDao) SETBIT(key string, bit, value int) (ret int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	ret, err = redis.Int(conn.Do("SETBIT", key, bit, value))
	if err != nil {
		return
	}
	return
}
func (this *RedisDao) GETBIT(key string, bit int) (ret int, err error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	ret, err = redis.Int(conn.Do("GETBIT", key, bit))
	if err != nil {
		return
	}
	return
}
