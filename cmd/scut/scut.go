package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	amqpURI := flag.String("amqp-uri", "amqp://guest:guest@127.0.0.1:5672", "amqp uri to connect to")
	exchangeName := flag.String("exchange-name", "", "name of the exchange to consume from")
	bindingKey := flag.String("binding-key", "", "binding key to the exchange")
	flag.Parse()

	conn, err := amqp.Dial(*amqpURI)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	queue, err := channel.QueueDeclare("", false, true, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(queue.Name, *bindingKey, *exchangeName, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	deliveries, err := channel.Consume(queue.Name, "scut", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	for d := range deliveries {
		fmt.Println(string(d.Body))
	}
}
