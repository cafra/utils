package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
	"time"
)

var (
	KAddrs = "10.60.81.181:9092"
	KTopic = "test"
)

func TestConsumer(t *testing.T) {
	consumer, err := NewConsumer(KAddrs, KTopic)
	if err != nil {
		t.Fatal(err)
	}

	consumer.Serve(func(m *sarama.ConsumerMessage) error {
		t.Log(m.Topic, string(m.Value))
		return nil
	})
}

func BenchmarkConsume(b *testing.B) {
	consumer, err := NewConsumer(KAddrs, KTopic)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < 5; i++ {
		go b.Log(consumer.Serve(func(m *sarama.ConsumerMessage) error {
			b.Log(i, "|", m.Topic, string(m.Value))
			fmt.Println(i, "|", m.Topic, string(m.Value))
			return nil
		}))
	}

	time.Sleep(200 * time.Second)
}
