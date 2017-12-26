package utils

import (
	"bytes"
	"os"
	"testing"
	"utils/db"

	"github.com/dchest/captcha"
	"io/ioutil"
)

var (
	store *RedisStore4Captcha
)

func init() {
	dao, err := db.NewRedisDao("redis://@127.0.0.1:6379/1?idle=100&active=100&wait=true&timeout=3s", true)
	if err != nil {
		panic(err)
	}
	store = NewRedisStore4Captcha(dao)
	captcha.SetCustomStore(store)
}

func TestGet(t *testing.T) {
	store.Set("11111111", S2B("222"))
	t.Log(B2S(store.Get("11111111", false)))
}

func TestNewcaptcha(t *testing.T) {
	id := captcha.New()

	buf := bytes.NewBufferString("")
	captcha.WriteImage(buf, id, 240, 80)
	if err := ioutil.WriteFile("/Users/cz/Downloads/1.png", buf.Bytes(), os.ModePerm); err != nil {
		t.Fatal(err)
	}
	t.Logf("id=%v", id)
}
func TestCheck(t *testing.T) {
	t.Log(captcha.Verify("MzGiwBlgyKzrd5teySPa", ConvertS2B("901965")))
}

func TestConvertS2B(t *testing.T) {
	t.Log(ConvertS2B("123"))
}
