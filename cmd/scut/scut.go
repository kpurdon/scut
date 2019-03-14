package main

import (
	"flag"
	"fmt"

	"github.com/streadway/amqp"
)

// version is populated by go build
var version = "development"

func main() {
	amqpURI := flag.String("amqp-uri", "amqp://guest:guest@127.0.0.1:5672", "amqp uri to connect to")
	exchangeName := flag.String("exchange-name", "", "name of the exchange to consume from")
	bindingKey := flag.String("binding-key", "", "binding key to the exchange")
	printVersion := flag.Bool("version", false, "prints the version")
	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		return
	}

	conn, err := amqp.Dial(*amqpURI)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	queue, err := channel.QueueDeclare("", false, true, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = channel.QueueBind(queue.Name, *bindingKey, *exchangeName, false, nil)
	if err != nil {
		panic(err)
	}

	deliveries, err := channel.Consume(queue.Name, "scut", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for d := range deliveries {
		fmt.Println(string(d.Body))
	}
}
