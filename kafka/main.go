package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {

	//1.生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // ACK
	config.Producer.Partitioner = sarama.NewRandomPartitioner //分区
	config.Producer.Return.Successes = true                   //回复确认

	//2.连接kafka
	client, err := sarama.NewSyncProducer([]string{"0.0.0.0:9092"},
		config)
	if err != nil {
		fmt.Println("producer closed,err:", err)
	}
	defer client.Close()

	//3.封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("2023-07-15 this is a test log22")

	//4.发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed,err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
