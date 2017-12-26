package utils

import (
	"fmt"
	"strconv"
	"strings"

	"utils/db"
)

const _RedisTTl = 15 * 60 //防止恶意产生大量redis

type RedisStore4Captcha struct {
	redis *db.RedisDao
}

const captcha_prefix = "captcha:%v"

func NewRedisStore4Captcha(r *db.RedisDao) (store *RedisStore4Captcha) {
	return &RedisStore4Captcha{redis: r}
}

/*
s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, "")
*/
func ConvertS2B(d string) (s []byte) {
	s = make([]byte, len(d))

	for i, c := range strings.Split(d, "") {
		t, err := strconv.Atoi(c)
		if err != nil {
			return nil
		}
		s[i] = byte(t)
	}
	return
}

func (store *RedisStore4Captcha) Set(id string, digits []byte) {
	key := fmt.Sprintf(captcha_prefix, id)

	if err := store.redis.SETEX(key, _RedisTTl, digits); err != nil {
		fmt.Errorf("RedisStore4Captcha |Set err=%v", err)
	}
}

func (store *RedisStore4Captcha) Get(id string, clear bool) (digits []byte) {
	var err error
	key := fmt.Sprintf(captcha_prefix, id)
	digits, err = store.redis.GetBytes(key)
	if err != nil {
		fmt.Errorf("RedisStore4Captcha |Get err=%v", err)
	}
	if clear {
		if _, err = store.redis.Del(key); err != nil {
			fmt.Errorf("RedisStore4Captcha |Del err=%v", err)
		}
	}
	return
}
