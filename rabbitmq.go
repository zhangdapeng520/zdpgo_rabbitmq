package zdpgo_rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// RabbitMQ结构体
type RabbitMQ struct {
	Username string // 用户名
	Password string // 密码
	Ip string // IP
	Port int // 端口号
	Conn    *amqp.Connection // 连接
}

// 获取默认的结构体
func NewDefaultRabbitMQ() *RabbitMQ{
	mq := RabbitMQ{
		Username: "guest",
		Password: "guest",
		Ip: "127.0.0.1",
		Port: 5672,
	}
	return &mq
}

//  建立连接
func (mq *RabbitMQ)Connect(){
	addr := fmt.Sprintf("amqp://%s:%s@%s:%d/",
	mq.Username,
	mq.Password,
	mq.Ip, 
	mq.Port )
	conn, err := amqp.Dial(addr)
	if err != nil{
		panic(err)
	}
	mq.Conn = conn
}

// 创建通道
func (mq *RabbitMQ) CreateChannel()*amqp.Channel{
	//创建一个Channel，所有的连接都是通过Channel管理的
	channel, err := mq.Conn.Channel()
	if ; err != nil{
		panic(err)
	}
	return channel
}

// 创建队列
func (mq *RabbitMQ) CreateQueue(channel *amqp.Channel, queueName string)amqp.Queue{
	queue, err := channel.QueueDeclare(queueName, false, false, false, false, nil)
	if  err != nil{
		panic(err)
	}
	return queue
}

// 发布消息
func (mq *RabbitMQ) Pub(channel *amqp.Channel, queue amqp.Queue, msg string) error{
	err := channel.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	}); 
	return err
}

// 订阅消息
func (mq *RabbitMQ) Sub(channel *amqp.Channel, queue amqp.Queue) <-chan amqp.Delivery{
	msgs, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil{
		panic(err)
	}
	return msgs
}

// 发送数据
func (mq *RabbitMQ) Send(queueName, msg string){
	//连接MQServer
	mq.Connect()
	defer mq.Conn.Close()

	//创建一个Channel，所有的连接都是通过Channel管理的
	channel := mq.CreateChannel()

	//创建队列
	queue := mq.CreateQueue(channel,queueName)

	//发送数据
	mq.Pub(channel, queue,msg)
}

// 接收消息
func (mq *RabbitMQ)Receive(queueName string){
	//连接MQServer
	mq.Connect()
	defer mq.Conn.Close()

	//创建一个Channel，所有的连接都是通过Channel管理的
	channel := mq.CreateChannel()

	//创建队列
	queue := mq.CreateQueue(channel, queueName)

	//读取数据
	msgs := mq.Sub(channel, queue)

	go func(){
		for msg := range msgs{
			fmt.Println("Receive Msg =", string(msg.Body))
		}
	}()
}