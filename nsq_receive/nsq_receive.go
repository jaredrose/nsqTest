package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

type ConsumerT struct {
}

func main() {
	InitConsumer("test", "test-channel", "127.0.0.1:4161")
	for {
		time.Sleep(time.Second * 10)
	}
}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, string(msg.Body))
	return nil
}

func InitConsumer(topic string, channel string, address string) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		panic(err)
	}

	c.SetLogger(nil, 0)        //屏蔽系统日志
	c.AddHandler(&ConsumerT{}) //添加消费者接口
	// 建立nsqlookupd连接
	// if err := c.ConnectToNSQLookupd(address); err != nil {
	// 	panic(err)
	// }

	// 建立多个nsqd 连接
	// if err := c.ConnectToNSQDs([]string{"127.0.0.1:4150", "127.0.0.1:4152"}); err != nil {
	// 	panic(err)
	// }

	// 建立一个nsqd 连接

	if err := c.ConnectToNSQD("127.0.0.1:4150"); err != nil {
		panic(err)
	}
}
