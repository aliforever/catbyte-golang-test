package main

import (
	"catbyte-golang-test/lib/publishing-api"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://user:password@host.docker.internal:7001/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare("catbyte", "fanout", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	_, err = ch.QueueDeclare("messages", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = ch.QueueBind("messages", "", "catbyte", false, nil)
	if err != nil {
		panic(err)
	}

	err = publishing_api.NewAPIService(ch).Start(":8080")
	if err != nil {
		panic(err)
	}
}
