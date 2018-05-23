package utils

import (
	//"fmt"
	"github.com/astaxie/beego/logs"
	//"os"
	////"path"
	//"path/filepath"
	//"fmt"
	"testing"
	//"time"
)

func TestTest(t *testing.T) {
	for i := 0; i < 1; i++ {
		l.Debug(`{"uid":10000,"age":11}`)
		l2.Debug(`{"uid":10001,"age":11}`)
		//fmt.Sprintf(MsgLayoutJson, "xx", "xx", "xxx", "1111111111", "cdcdsccdcd", "cdcdcdc")
	}
}

var (
	l  = logs.NewLogger(0)
	l2 = logs.NewLogger(0)
)

func init() {
	l.SetLogger(AdapterMyFile, `{"filename":"/Users/cz/Downloads/api/api-action1.log","maxdays":365,"module":"main","action":"test","ip":"127.0.0.1","logger_func_call_depth":4,"layout":"json"}`)
	l2.SetLogger(AdapterMyFile, `{"filename":"/Users/cz/Downloads/api/api-action1.log","maxdays":365,"module":"main","action":"test2","ip":"127.0.0.1","logger_func_call_depth":4,"layout":"json"}`)
}
func BenchmarkMylog(b *testing.B) {
	//logs.Async(10e5)

	//logs.SetLogger(logs.AdapterFile, `{"filename":"/Users/cz/Downloads/api/api.log","maxdays":60}`)
	//logs.SetLogFuncCall(true)
	//logs.SetLogFuncCallDepth(3)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l.Debug(`{"uid":10000,"age":11}`)
		//l2.Debug(`{"uid":10001,"age":11}`)
		//fmt.Sprintf(MsgLayoutJson, "xx", "xx", "xxx", "1111111111", "cdcdsccdcd", "cdcdcdc")
	}

	//time.Sleep(time.Second)
}
