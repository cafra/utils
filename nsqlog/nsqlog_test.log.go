package nsqlog

import (
	//"gintest/beegolog/nsqlog"
	"github.com/astaxie/beego/logs"
	"testing"
)

func TestLog(t *testing.T) {
	logs.SetLogger(AdapterNSQ, `{
	  "nsqaddr": "172.16.13.21:4150",
	  "topic": "readingmate_wx",
	  "tplId":"0HaC5K_8AM5nqwH4DNcPCFNwtV_IoR9cAXx-oCvBHac",
	  "wxids":["obXYRxEDQWjqaWmkXgaUrrzsk_gA"]
	}`)

	logs.Error("111")
}
