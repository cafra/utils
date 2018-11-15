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

func NewConsumer(brokers, topics, group_id string) (consumer *Consumer, err error) {
	consumer = new(Consumer)
	groupID := group_id
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	// 重要！！！！！！！！！！！
	// OffsetNewest:pub 每次启动从队列的最新数据开始消费
	// OffsetOldest: pub 每次启动从队列上次消费的地方开始消费
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

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

// Handler 错误则不commit.下次启动可在此消费
func (c *Consumer) Serve(h Handler) (err error) {
	for {
		select {
		case msg, ok := <-c.cli.Messages():
			if ok {
				if h(msg) != nil {
					log.Printf("Consumer|Serve handler err=%v", err)
					continue
				}
				//注意！！！！ 如果panic ，系统重启等，下次都会从上次最后一个数据消费。保证数据不丢失
				c.cli.MarkOffset(msg, "") // mark message as processed
			}
		case <-c.signals:
			return
		}
	}
	return
}
