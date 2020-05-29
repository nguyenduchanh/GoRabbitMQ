package main

import (
	"GoRabbitMQ/api"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	conn, err := amqp.Dial("amqp://admin:abc123@10.1.3.104:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"Sender-Exchange", // name
		"topic",           // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		true,  // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,            // queue name
		"red-code",        // routing key
		"Sender-Exchange", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			//CallCarTypeInsert(d.Body)
			OperatorApi(d.Body)
			//log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func OperatorApi(msg []byte) {
	var obj *api.ResponseModel
	if err := json.Unmarshal(msg, &obj); err != nil {
		panic(err)
	}
	model := obj.Model
	method := obj.Method

	if model == "carType" {
		var carTypeModel *api.CarType
		json.Unmarshal([]byte(obj.Entity), &carTypeModel)
		if method == "post" {
			api.InsertCarType(carTypeModel)
		} else if method == "put" {
			api.UpdateCaeType(carTypeModel.Id, carTypeModel)
		} else if method == "get" {
			api.GetAllCarType()
		}
	} else if model == "operator" {
		var operatorModel *api.Operator
		json.Unmarshal([]byte(obj.Entity), &operatorModel)
		if method == "post" {
			api.OperatorInsertFirstCall([]byte(obj.Entity))
		}
	}
}
