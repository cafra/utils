package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
	"time"
)

var (
	KAddrs = "localhost:9092"
	KTopic = "test"
)

func TestConsumer(t *testing.T) {
	consumer, err := NewConsumer(KAddrs, KTopic, "f1")
	if err != nil {
		t.Fatal(err)
	}
	consumer2, err := NewConsumer(KAddrs, KTopic, "f2")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		consumer.Serve(func(m *sarama.ConsumerMessage) error {
			//t.Log(m.Topic, string(m.Value))
			fmt.Println("f1", m.Topic, string(m.Value))
			return nil
		})
	}()
	consumer2.Serve(func(m *sarama.ConsumerMessage) error {
		//t.Log(m.Topic, string(m.Value))
		fmt.Println("f2", m.Topic, string(m.Value))
		return nil
	})
}

func BenchmarkConsume(b *testing.B) {
	consumer, err := NewConsumer(KAddrs, KTopic, "f")
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
