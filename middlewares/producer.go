package middlewares

import (
	"DailyEnglish/utils"

	"log"

	"github.com/streadway/amqp"
)

func RunProducer() {
	conn, err := amqp.Dial("amqp://guest:guest@47.113.117.103:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个队列
	q, err := ch.QueueDeclare(
		"hello", // name: 队列名称
		false,   // durable: 是否持久化
		false,   // delete when unused: 是否自动删除
		false,   // exclusive: 是否独占
		false,   // no-wait: 是否非阻塞
		nil,     // arguments: 额外参数
	)
	utils.FailOnError(err, "Failed to declare a queue")

	// 准备要发送的消息
	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange: 交换器名称，这里使用默认的交换器
		q.Name, // routing key: 路由键，这里使用队列名称
		false,  // mandatory: 是否强制
		false,  // immediate: 是否立即
		amqp.Publishing{
			ContentType: "text/plain", // 内容类型
			Body:        []byte(body), // 消息体
		})
	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
