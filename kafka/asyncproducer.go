package producer

import (
	"errors"
	"fmt"

	"github.com/Shopify/sarama"
)

type KafkaAsyncProducer struct {
	sarama.AsyncProducer
}

func (producer *KafkaAsyncProducer) SendString(topic, value string) {
	if len(value) == 0 {
		return
	}

	if topic == "" {
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(value),
	}
	producer.Input() <- msg
}

func (producer *KafkaAsyncProducer) SendBytes(topic string, value []byte) {
	if len(value) == 0 {
		return
	}

	if topic == "" {
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}
	producer.Input() <- msg
}

func (producer *KafkaAsyncProducer) Close() {
	producer.AsyncClose()
}

func InitKafkaAsyncProducer(brokers []string) (*KafkaAsyncProducer, error) {
	if len(brokers) == 0 {
		return nil, errors.New("kafka brokers is empty")
	}

	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
	config.Version = sarama.V0_10_0_1

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {

		return nil, err

	}

	go func(p sarama.AsyncProducer) {
		for {
			select {
			case <-p.Successes():
				//fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
			case fail := <-p.Errors():
				fmt.Println("kafka send err: ", fail.Err)
			}
		}
	}(producer)

	p := KafkaAsyncProducer{
		producer,
	}
	return &p, nil
}
