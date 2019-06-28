package rabbitmq

import (
	"log"
	"strategy_middle/setting"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
)
var Allws []*websocket.Conn

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Send(sendchanneL_name string, msg string) {
	//账号和密码
	//
	upip := fmt.Sprintf("amqp://%s:%s@%s/",setting.Rabbitmq_USER,setting.Rabbitmq_PASSWORD,setting.Rabbitmq_HOST)


	conn, err := amqp.Dial(upip) //dev:dev@192.168.18.24:5672 //guest:guest@localhost:5672
	failOnError(err, "rabbitmq:Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "rabbitmq:Failed to open a channel")
	defer ch.Close()
	//定义队列
	q, err := ch.QueueDeclare(
		sendchanneL_name, // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "rabbitmq:Failed to declare a queue")

	body := msg
	//订阅,发送消息
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "rabbitmq:Failed to publish a message")
}

func Recv(sendchanneL_name string) {

	upip := fmt.Sprintf("amqp://%s:%s@%s/",setting.Rabbitmq_USER,setting.Rabbitmq_PASSWORD,setting.Rabbitmq_HOST)
	conn, err := amqp.Dial(upip)

	failOnError(err, "rabbitmq:Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "rabbitmq:Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		sendchanneL_name, // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "rabbitmq:Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "rabbitmq:Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("rabbitmq:Received a message: %s", d.Body)
			for e := range Allws{

				fmt.Println(e)
				if len(Allws)>0{
					err = e.WriteMessage('', []byte(d.Body))
					if err!=nil{
						fmt.Println(err)
					}
				}

			}
		}
	}()
	log.Printf(" rabbitmq:[*] Waiting for messages. ")
	<-forever
}
