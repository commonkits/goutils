package producer

import (
	"testing"
	"time"
)

func Test_Producer(t *testing.T) {
	producer, err := InitKafkaAsyncProducer(DefaultConfig(),
		[]string{"172.20.61.203:6667", "172.20.61.204:6667", "172.20.61.205:6667"},
		DefaultSuccessFunc(),
		DefaultErrorFunc())

	if err != nil {
		panic(err)
	}

	producer.SendString("monitor.log.test", "sdfadfasdfasf")

	time.Sleep(10 * time.Second)
}
