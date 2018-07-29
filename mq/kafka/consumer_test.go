package kafka

import (
	"github.com/Shopify/sarama"
	"testing"
)

var (
	KAddrs = "10.60.81.181:9092"
	KTopic = "test"
)

func TestConsumer(t *testing.T) {
	//NewConsumer(KAddrs, KTopic, func(msg *sarama.ConsumerMessage) error {
	//	t.Log(msg.Topic, msg.Value)
	//	return nil
	//})
}
