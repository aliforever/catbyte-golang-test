package main

import (
	"catbyte-golang-test/lib/processor"
	"catbyte-golang-test/shared/logger"
	"context"
	"github.com/go-redis/redis/v8"
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

	queue, err := ch.QueueDeclare("messages", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = ch.QueueBind("messages", "", "catbyte", false, nil)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:6379",
		Password: "",
		DB:       0,
	})

	// This is to make sure redis is alive
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	err = processor.NewProcessor(ch, queue.Name, rdb, logger.DefaultLogger{}).Start()
	if err != nil {
		panic(err)
	}
}
