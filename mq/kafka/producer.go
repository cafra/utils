package kafka

import (
	"log"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

type Producer struct {
	cli sarama.SyncProducer
}

func NewProducer(brokers string) (pd *Producer, err error) {
	pd = new(Producer)
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true //必须有这个选项
	config.Producer.Timeout = 5 * time.Second
	//p, err := sarama.NewAsyncProducer(strings.Split(brokers, ","), config)
	pd.cli, err = sarama.NewSyncProducer(strings.Split(brokers, ","), config)
	if err != nil {
		log.Fatal("initProducer |NewSyncProducer err=%v", err)
		return
	}
	return
}
func (p *Producer) Write(topic, m string) (err error) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(m),
	}
	_, _, err = p.cli.SendMessage(msg)
	return
}

func (p *Producer) Close() {
	p.cli.Close()
}
