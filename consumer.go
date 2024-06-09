package main

import (
	"DailyEnglish/utils"

	"log"

	"github.com/streadway/amqp"
)

func RunConsumer() {
	conn, err := amqp.Dial("amqp://guest:guest@47.113.117.103:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	// 注册消费者
	msgs, err := ch.Consume(
		q.Name, // queue: 队列名称
		"",     // consumer: 消费者标签（空字符串表示自动生成）
		true,   // auto-ack: 自动确认
		false,  // exclusive: 是否独占
		false,  // no-local: 不接收自身发布的消息
		false,  // no-wait: 是否非阻塞
		nil,    // args: 额外参数
	)
	utils.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool) //创建一个通道来保持程序运行

	// 启动一个 goroutine 来处理消息
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
