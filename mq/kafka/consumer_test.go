package kafka

import (
	//"fmt"
	"github.com/Shopify/sarama"
	"testing"
	//"time"
	//"github.com/Shopify/sarama"
	"fmt"
)

var (
	//KAddrs = "149.129.215.154:9092,149.129.215.154:9093,149.129.215.154:9094"
	KAddrs = "localhost:9092"
	KTopic = "cztest01"
)

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d	value:%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func TestConsumer(t *testing.T) {
	NewConsumer2(KAddrs, KTopic, "group_cz", exampleConsumerGroupHandler{})

	//c, err := NewConsumer(KAddrs, KTopic, "group_test3")
	//t.Log(err)
	//c.Serve(func(message *sarama.ConsumerMessage) error {
	//	fmt.Printf("%s 	%s\n", message.Topic, message.Value)
	//	return nil
	//})
}
