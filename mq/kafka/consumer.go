package kafka

import (
	"log"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster" //support automatic consumer-group rebalancing and offset tracking
)

type Consumer struct {
	cli *cluster.Consumer
}
type Handler func(*sarama.ConsumerMessage) error

func NewConsumer(brokers, topics string) (consumer *Consumer, err error) {
	consumer = new(Consumer)
	groupID := "group-1"
	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer.cli, err = cluster.NewConsumer(strings.Split(brokers, ","), groupID, strings.Split(topics, ","), config)
	if err != nil {
		log.Fatal("Failed open consumer: %v", err)
		return
	}

	go func(c *cluster.Consumer) {
		errors := c.Errors()
		noti := c.Notifications()
		for {
			select {
			case err := <-errors:
				log.Printf("Notifications errrs %v", err)
			case <-noti:
			}
		}
	}(consumer.cli)
	return
}

func (c *Consumer) Serve(h Handler) (err error) {
	var cnt int = 0
	for msg := range c.cli.Messages() {
		c.cli.MarkOffset(msg, "")
		cnt++

		if err = h(msg); err != nil {
			return
		}
	}
	return
}
