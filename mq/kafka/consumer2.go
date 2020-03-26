package kafka

import (
	"context"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
)

type CommonHandler func(*sarama.ConsumerMessage) (err error, reConsume bool)

type exampleConsumerGroupHandler struct {
	handler CommonHandler
}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) (err error) {
	for msg := range claim.Messages() {
		err, reConsume := h.handler(msg)
		fmt.Printf("Message topic:%q 	partition:%d 	offset:%d	value:%s	err=%v	reConsume=%v\n ",
			msg.Topic, msg.Partition, msg.Offset, msg.Value, err, reConsume)
		if !reConsume {
			// 不需要重复消费，则自动提交
			sess.MarkMessage(msg, "")
		}
	}
	return nil
}

// brokers: 逗号分隔的服务器列表
// topics:	逗号分隔的topic列表
// group_id:消费组名字
// handler:	消费接口
func NewConsumer2(brokers, topics, group_id string, handler CommonHandler) {
	// Init config, specify appropriate version
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_0
	config.Consumer.Return.Errors = true

	// 重要！！！！！！！！！！！
	// OffsetNewest:pub 每次启动从队列的最新数据开始消费
	// OffsetOldest: pub 每次启动从队列上次消费的地方开始消费
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Start with a client
	brokersList := strings.Split(brokers, ",")
	client, err := sarama.NewClient(brokersList, config)
	if err != nil {
		panic(err)
	}
	// gorputine 退出时，关闭客户端
	defer func() { _ = client.Close() }()

	// Start a new consumer group
	group, err := sarama.NewConsumerGroupFromClient(group_id, client)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("KAFKA |group Errors=", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	topicsList := strings.Split(topics, ",")

	for {
		err := group.Consume(ctx, topicsList,
			exampleConsumerGroupHandler{
				handler: handler,
			})
		if err != nil {
			fmt.Println("KAFKA |Consume err=", err)
		}
	}

	return
}
