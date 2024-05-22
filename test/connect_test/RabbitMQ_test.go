package test

import (
	"testing"

	"github.com/streadway/amqp"
)

func TestRabbitMQConnection(t *testing.T) {
	// 假设RabbitMQ服务运行在本地，且默认端口5672上
	conn, err := amqp.Dial("amqp://guest:guest@47.113.117.103:5672/")
	if err != nil {
		t.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// 测试是否能够成功打开一个通道
	ch, err := conn.Channel()
	if err != nil {
		t.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// 如果我们能够到达这里，说明连接和通道都建立成功了
	// 可以添加更多的测试来验证队列声明、消息发布和消费等功能
	t.Logf("Connected to RabbitMQ and opened a channel successfully")
}
