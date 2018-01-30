package mq

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println(NewNsqProducer(`{
      	  "host": "localhost:4150",
      	  "verbose": true,
      	  "default_requeue_delay": 0,
      	  "max_attempts": 1
      	}`))
}

func TestNsqCommonPush(t *testing.T) {
	NsqCommonPush("test","haha")
}

//func TestNewNsqConsumer(t *testing.T) {
//	cs, err := NewNsqConsumer([]string{"localhost:4150"},
//		map[string]nsq.HandlerFunc{
//			"test": func(message *nsq.Message) error {
//				t.Log("==========", message.Timestamp, message.Body)
//				//message.RequeueWithoutBackoff(10 * time.Second)
//				return nil
//			},
//		},
//		"api")
//	t.Log("===", cs, err)
//	cs.Run()
//}