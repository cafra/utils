package nsqlog

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/astaxie/beego/logs"
	nsq "github.com/bitly/go-nsq"
	"github.com/cafra/utils"
)

type WxTemplateData struct {
	Touser     string                     `json:"touser"`
	TemplateId string                     `json:"template_id"` //模版id
	Url        string                     `json:"url"`
	TopColor   string                     `json:"top_color"`
	Data       map[string]*WxTemplateInfo `json:"data"` //string就是模版的key
}
type WxTemplateInfo struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

//  return a LoggerInterface
func NewNSQ() logs.Logger {
	nsq := &nsqLogger{
		Level: logs.LevelError,
	}

	nsq.HostName, _ = os.Hostname()
	nsq.Ip = GetOutboundIP()
	//nsq.Debug = true

	return nsq
}
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

type nsqLogger struct {
	nsqProducer *nsq.Producer
	NsqAddr     string `json:"nsqaddr"`
	Topic       string `json:"topic"`
	//开启测试输出模式，默认关闭
	Debug bool `json:"debug"`
	//host ip
	HostName string `json:"host"`
	Ip       string `json:"ip"`
	Level    int    `json:"level"`
	//这部分针对错误时,通过nsq推送错误信息,当然这部分只是推送，还需要有消费着，完成微信推送
	TmpId string   `json:"tplId"` //微信消息模版ID
	Wxids []string `json:"wxids"` //关注微信公众号，接受错误推送者的openid集合
}

func (el *nsqLogger) debug(format string, v ...interface{}) {
	if el.Debug {
		logs.Debug(format, v...)
	}
}

/*
{
  "nsqadd": "172.16.13.21:4150",
  "topic": "api_log",
  "host": "172.16.13.73"
}
*/
func (el *nsqLogger) Init(jsonconfig string) error {
	el.debug("Init |args=%v", jsonconfig)
	err := json.Unmarshal([]byte(jsonconfig), el)
	if err != nil {
		el.debug("Init err=%v", err)
		return err
	}
	if len(el.NsqAddr) == 0 {
		return errors.New("empty nsqaddr")
	}
	if len(el.Topic) == 0 {
		return errors.New("empty topic")
	}
	if len(el.TmpId) > 0 && len(el.Wxids) == 0 {
		return errors.New("wx push args err")
	}

	nsqConfig := nsq.NewConfig()
	nsqConfig.Set("verbose", el.Debug)
	nsqConfig.Set("default_requeue_delay", 0)
	nsqConfig.Set("max_attempts", 1)
	el.nsqProducer, err = nsq.NewProducer(el.NsqAddr, nsqConfig)
	if err != nil {
		el.debug("nsq NewProducer err=%v", err)
		return err
	}

	return nil
}

// WriteMsg will write the msg and level into es
func (el *nsqLogger) WriteMsg(when time.Time, msg string, level int) error {
	if level > el.Level {
		el.debug("WriteMsg |level=%v|set-level=%v|msg=%v return", level, el.Level, msg)
		return nil
	}
	el.debug("WriteMsg |level=%v|set-level=%v|msg=%v continue", level, el.Level, msg)
	if len(el.TmpId) == 0 {
		el.debug("TmpId=%v log write", el.TmpId)
		return el.nsqProducer.Publish(el.Topic, utils.S2B(fmt.Sprintf("%v|%v|%v|%v", el.HostName, el.Ip, time.Now().Format(_TimeFormat), msg)))
	}
	el.debug("WriteMsg|wx push msg=%v", msg)
	return el.eachPush(el.packWxMsg(msg))
}

// Destroy is a empty method
func (el *nsqLogger) Destroy() {

}

// Flush is a empty method
func (el *nsqLogger) Flush() {

}
func (el *nsqLogger) packWxMsg(msg string) (wxTemplate *WxTemplateData) {
	// 推送参数
	wxTemplate = new(WxTemplateData)
	wxTemplate.TemplateId = el.TmpId
	wxTemplate.Url = ""
	wxTemplate.TopColor = "#FF0000"
	wxTemplate.Data = map[string]*WxTemplateInfo{
		"first": &WxTemplateInfo{
			Value: fmt.Sprintf("服务器组[图书馆GO_API(%v)]消息", el.Ip),
			Color: "#173177",
		},
		"keyword1": &WxTemplateInfo{
			Value: "业务错误",
			Color: "#173177",
		},
		"keyword2": &WxTemplateInfo{
			Value: msg,
			Color: "#173177",
		},
		"keyword3": &WxTemplateInfo{
			Value: time.Now().Format("2006-01-02 15:04:05"),
			Color: "#173177",
		},
		"remark": &WxTemplateInfo{
			Value: "服务器运行状态监控消息，请持续关注",
			Color: "#173177",
		},
	}
	return
}

func (el *nsqLogger) eachPush(wxTemplate *WxTemplateData) (err error) {
	el.debug("eachPush |Wxids=%v|wxTemplate=%v", el.Wxids, *wxTemplate)
	for _, id := range el.Wxids {
		wxTemplate.Touser = id
		buf, err := json.Marshal(wxTemplate)
		if err != nil {
			el.debug("eachPush |wxTemplate=%v |err=%v", wxTemplate, err)
			return err
		}
		if err = el.nsqProducer.Publish(el.Topic, buf); err != nil {
			el.debug("Publish |wxTemplate=%v|err=%v", wxTemplate, err)
			return err
		}
	}
	return nil
}

const AdapterNSQ = "nsq"
const _TimeFormat = "15:04:05.99999"

func init() {
	logs.Register(AdapterNSQ, NewNSQ)
}
