package main

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
)

var producer *nsq.Producer

func main() {
	IP1 := "127.0.0.1:4150"
	IP2 := "127.0.0.1:4152"
	InitProducer(IP1)

	running := true
	// 从控制台输入
	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "stop" {
			running = false
		}

		for err := Publish("test", command); err != nil; err = Publish("test", command) {
			IP1, IP2 = IP2, IP1
			InitProducer(IP1)
		}

	}

	producer.Stop()

}

func InitProducer(ip string) {
	var err error
	fmt.Println("address:" + ip)
	producer, err = nsq.NewProducer(ip, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
}

func Publish(topic string, message string) error {
	var err error
	if producer != nil {
		// 不能发空字符串
		if message == "" {
			return nil
		}
		return producer.Publish(topic, []byte(message))
	}
	return fmt.Errorf("producer is nil", err)
}
