package utils

import (
	"encoding/json"
	nsq "github.com/bitly/go-nsq"
)

var _Producer *nsq.Producer
var _cfg *_NSQProducerCfg

type _NSQProducerCfg struct {
	Host                string `json:"host"`
	Verbose             bool   `json:"verbose"`
	DefaultRequeueDelay int    `json:"default_requeue_delay"`
	MaxAttempts         int    `json:"max_attempts"`
}

func NewNsqProducer(cfg string) (err error) {
	if _cfg == nil {
		_cfg = new(_NSQProducerCfg)
		if err = json.Unmarshal([]byte(cfg), _cfg); err != nil {
			return
		}
		newp()
	}
	return nil
}
func newp() {
	nsqConfig := nsq.NewConfig()
	nsqConfig.Set("verbose", _cfg.Verbose)
	nsqConfig.Set("default_requeue_delay", _cfg.DefaultRequeueDelay)
	nsqConfig.Set("max_attempts", _cfg.MaxAttempts)
	_Producer, _ = nsq.NewProducer(_cfg.Host, nsqConfig)
}
func producer() *nsq.Producer {
	if _Producer != nil {
		return _Producer
	}
	newp()
	return _Producer
}

func NsqCommonPush(topic string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return producer().Publish(topic, data)
}
