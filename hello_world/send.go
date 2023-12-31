package main

import (
	"context"
	"github.com/atadzan/message-broker/helpers"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func main() {
	conn := helpers.NewBrokerConnection()
	defer conn.Close()

	// Opening a channel
	ch := helpers.NewChannel(conn)
	defer ch.Close()

	// Declaring a queue to publish messages
	q := helpers.NewQueue(ch)

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
	helpers.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
