package kafka

import (
	//"fmt"
	"github.com/Shopify/sarama"
	"testing"
	//"time"
	//"github.com/Shopify/sarama"
	//"fmt"
)

var (
	//KAddrs = "149.129.215.154:9092,149.129.215.154:9093,149.129.215.154:9094"
	KAddrs = "localhost:9092"
	KTopic = "cztest02"
)

func TestConsumer(t *testing.T) {
	NewConsumer2(KAddrs, KTopic, "group_cz", func(m *sarama.ConsumerMessage) (err error, reConsume bool) {
		t.Log(m.Value)
		return
	})

	//c, err := NewConsumer(KAddrs, KTopic, "group_test3")
	//t.Log(err)
	//c.Serve(func(message *sarama.ConsumerMessage) error {
	//	fmt.Printf("%s 	%s\n", message.Topic, message.Value)
	//	return nil
	//})
}
