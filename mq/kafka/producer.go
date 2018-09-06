package kafka

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
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
		log.Fatal("initProducer |NewSyncProducer err=", err)
		return
	}
	return
}
func (p *Producer) Write(topic string, data interface{}) (err error) {
	str := ""
	kind := reflect.ValueOf(data).Kind()
	switch kind {
	case reflect.String:
		str = data.(string)
	case reflect.Ptr:
		fallthrough
	case reflect.Struct:
		bs, _ := json.Marshal(data)
		str = string(bs)
	default:
		return errors.New(fmt.Sprintf("Write data type =%v err", kind))
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(str),
	}
	if _, _, err = p.cli.SendMessage(msg); err != nil {
		log.Println("Write |SendMessage err=%v", err)
	}
	return
}

func (p *Producer) Close() {
	p.cli.Close()
}
