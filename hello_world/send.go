package main

import (
	"context"
	"github.com/atadzan/message-broker/hello_world/helper"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func main() {
	conn := helper.NewBrokerConnection()
	defer conn.Close()

	// Opening a channel
	ch := helper.NewChannel(conn)
	defer ch.Close()

	// Declaring a queue to publish messages
	q := helper.NewQueue(ch)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World 2"
	err := ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	helper.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
