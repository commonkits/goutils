package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func DefaultConfig() *sarama.Config {
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
	return config
}

func DefaultSuccessFunc() func(msg *sarama.ProducerMessage) {
	return func(msg *sarama.ProducerMessage) {}
}

func DefaultErrorFunc() func(fail *sarama.ProducerError) {
	return func(fail *sarama.ProducerError) { fmt.Println("ERROR kafka send err: ", fail.Err) }
}
