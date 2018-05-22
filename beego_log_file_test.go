package utils

import (
	//"fmt"
	"github.com/astaxie/beego/logs"
	//"os"
	////"path"
	//"path/filepath"
	"testing"
)

func TestTest(t *testing.T) {
	//dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//cfg := fmt.Sprintf(`{"filename":"/Users/cz/Downloads/api/api.log","maxdays":365,"module":"main","ip":"127.0.0.1","logger_func_call_depth":5}`)
	//fmt.Println(logs.SetLogger(AdapterMyFile, cfg))
	////fmt.Printf(dir, cfg)
	//
	////logs.SetLogger(logs.AdapterConsole)
	//
	////l := logs.GetLogger("action=input")
	////logs.EnableFuncCallDepth(true)
	////log := logs.GetLogger("service1")
	////logs.Debug(`{"uid":10000,"age":11}`)
	//logs.Debug(`{"uid":10000,"age":11}`)

	//fmt.Printf(path.Base("/Users/cz/go/src/github.com/cafra/utils/beego_log_file_test.go"))
	//xxxlog("hello", "user")
}

func BenchmarkMylog(b *testing.B) {
	cfg := `{"filename":"/Users/cz/Downloads/api/api.log","maxdays":365,"module":"main","ip":"127.0.0.1","logger_func_call_depth":5}`
	logs.SetLogger(AdapterMyFile, cfg)
	//logs.Async(10e5)

	//logs.SetLogger(logs.AdapterFile, `{"filename":"/Users/cz/Downloads/api/api.log","maxdays":60}`)
	//logs.SetLogFuncCall(true)
	//logs.SetLogFuncCallDepth(3)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		logs.Debug(`{"uid":10000,"age":11}`)
		//fmt.Sprintf(MsgLayoutJson, "xx", "xx", "xxx", "1111111111", "cdcdsccdcd", "cdcdcdc")
	}
}
