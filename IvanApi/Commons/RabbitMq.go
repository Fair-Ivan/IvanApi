package Commons

import (
	"fmt"
	"github.com/streadway/amqp"
)

var r *amqp.Connection

// 新建实例
func NewConnection() *amqp.Connection {
	var connectString string = GetConfigJson().RabbitMqConnString
	conn, err := amqp.Dial(connectString)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		panic("连接Mq失败")
	}
	return conn
}

// 初始化连接
func RabbitMqInit() {
	r = NewConnection()
}

// 获取连接
func GetRabbitMqConn() *amqp.Connection {
	if r.IsClosed() {
		r = NewConnection()
	}
	return r
}

// 推送消息
func PublishMessage(message string) {
	r = GetRabbitMqConn()
	channel, err := r.Channel()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		panic("获取MqChannel失败")
		return
	}
	// 绑定交换机
	err = channel.ExchangeDeclare(
		"broadtest",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		panic("绑定exchange失败")
		return
	}
	// 推送消息
	err = channel.Publish(
		"broadtest",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// 消费消息
func ConsumeMessage() {
	r = GetRabbitMqConn()
	channel, err := r.Channel()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		panic("获取MqChannel失败")
		return
	}
	// 绑定交换机
	err = channel.ExchangeDeclare(
		"broadtest",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		panic("绑定exchange失败")
		return
	}
	// 创建队列
	q, err := channel.QueueDeclare(
		"broadtest", //随机生产队列名称
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		panic("绑定queue失败")
		return
	}
	// 绑定队列
	err = channel.QueueBind(
		q.Name,
		"", //在pub/sub模式下key要为空
		"broadtest",
		false,
		nil,
	)

	messages, err := channel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range messages {
			fmt.Printf("收到的消息：%s\n", d.Body)
			d.Ack(true)
		}
	}()

	<-forever
}
