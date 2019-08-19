package producer

import (
	"errors"
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

var (
	asyncProducer *KafkaAsyncProducer
	isInit        bool
)

func InitKafkaAsyncProducer(config *sarama.Config,
	brokers []string,
	sucFuc func(msg *sarama.ProducerMessage),
	errFunc func(fail *sarama.ProducerError)) (*KafkaAsyncProducer, error) {

	if isInit {
		return asyncProducer, nil
	}
	if len(brokers) == 0 {
		return nil, errors.New("kafka brokers is empty")
	}

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	go func(p sarama.AsyncProducer) {
		for {
			select {
			case msg := <-p.Successes():
				sucFuc(msg)
			case fail := <-p.Errors():
				errFunc(fail)
			}
		}
	}(producer)

	p := KafkaAsyncProducer{
		producer,
	}

	asyncProducer = &p
	isInit = true

	return &p, nil
}

func GetAsyncProducer() *KafkaAsyncProducer {
	return asyncProducer
}

func Reset() {
	isInit = false
}
