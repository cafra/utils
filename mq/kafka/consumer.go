package kafka

import (
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster" //support automatic consumer-group rebalancing and offset tracking
)

type Consumer struct {
	signals chan os.Signal
	cli     *cluster.Consumer
}
type Handler func(*sarama.ConsumerMessage) error

func NewConsumer(brokers, topics string) (consumer *Consumer, err error) {
	consumer = new(Consumer)
	groupID := "group-1"
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	consumer.cli, err = cluster.NewConsumer(strings.Split(brokers, ","), groupID, strings.Split(topics, ","), config)
	if err != nil {
		log.Printf("Failed open consumer: %v", err)
		return
	}

	consumer.signals = make(chan os.Signal, 1)
	signal.Notify(consumer.signals, os.Interrupt)

	go func(c *cluster.Consumer) {
		errors := c.Errors()
		noti := c.Notifications()
		for {
			select {
			case err := <-errors:
				log.Printf("Errors errrs %v", err)
			case <-noti:
				//log.Printf("Notifications errrs %v", tmp)
			}
		}
	}(consumer.cli)
	return
}

func (c *Consumer) Serve(h Handler) (err error) {
	for {
		select {
		case msg, ok := <-c.cli.Messages():
			if ok {
				if h(msg) != nil {
					log.Printf("Consumer|Serve handler err=%v", err)
					continue
				}
				c.cli.MarkOffset(msg, "") // mark message as processed
			}
		case <-c.signals:
			return
		}
	}
	return
}
