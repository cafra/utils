package mq
//
//import (
//	"fmt"
//	"os"
//	"os/signal"
//	"syscall"
//
//	"github.com/astaxie/beego/logs"
//	nsq "github.com/bitly/go-nsq"
//	"github.com/bitly/nsq/util"
//)
//
//type NsqConsumer struct {
//	channel         string
//	nsqdTCPAddrs    util.StringArray
//	consumerOpts    util.StringArray
//	topicsHandlers  map[string]nsq.HandlerFunc
//	topicsConsumers map[string]*nsq.Consumer
//
//	termChan chan os.Signal
//	hupChan  chan os.Signal
//}
//
//func NewNsqConsumer(hosts []string, topicsHandlers map[string]nsq.HandlerFunc, channelName string) (cs *NsqConsumer, err error) {
//	cs = new(NsqConsumer)
//	cs.nsqdTCPAddrs = make(util.StringArray, 0)
//	cs.consumerOpts = make(util.StringArray, 0)
//	for _, h := range hosts {
//		if err = cs.nsqdTCPAddrs.Set(h); err != nil {
//			logs.Error("NewNsqConsumer |nsqdTCPAddrs.Set host=%v err=%v", h, err)
//			return
//		}
//	}
//	cs.topicsConsumers = make(map[string]*nsq.Consumer)
//	cs.termChan = make(chan os.Signal)
//	cs.hupChan = make(chan os.Signal)
//	cs.channel = channelName
//	cs.topicsHandlers = topicsHandlers
//
//	return
//}
//
//func (nc *NsqConsumer) stop() {
//	for _, consumer := range nc.topicsConsumers {
//		consumer.Stop()
//		<-consumer.StopChan
//	}
//}
//
//func (t *NsqConsumer) watch() {
//	<-t.termChan
//	t.stop()
//	os.Exit(1)
//}
//
//func (nc *NsqConsumer) initConsumer() (err error) {
//	cfg := nsq.NewConfig()
//	cfg.UserAgent = fmt.Sprintf("nsq_to_file/%s go-nsq/%s", util.BINARY_VERSION, nsq.VERSION)
//	err = cfg.Set("max_attempts", 1)
//	//err = util.ParseOpts(cfg, consumerOpts)
//	if err != nil {
//		logs.Error("initConsumer |cfg.Set max_attempts err=%v", err)
//		return
//	}
//
//	for topic, consumer_handler := range nc.topicsHandlers {
//		consumer, err := nsq.NewConsumer(topic, nc.channel, cfg)
//		if err != nil {
//			logs.Error("initConsumer |nsq.initConsumer err (%v)", err)
//			return err
//		}
//
//		consumer.AddHandler(consumer_handler)
//
//		err = consumer.ConnectToNSQDs(nc.nsqdTCPAddrs)
//		if err != nil {
//			logs.Debug("initConsumer |ConnectToNSQDs err=%v", err)
//			return err
//		}
//		nc.topicsConsumers[topic] = consumer
//	}
//
//	return
//}
//
//func (n *NsqConsumer) Run() {
//	signal.Notify(n.hupChan, syscall.SIGHUP)
//	signal.Notify(n.termChan, syscall.SIGINT, syscall.SIGTERM)
//
//	go func() {
//		n.initConsumer()
//	}()
//
//	n.watch()
//}
//
