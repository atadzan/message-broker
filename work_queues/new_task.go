package main

import (
	"context"
	"github.com/atadzan/message-broker/helpers"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"time"
)

func main() {
	conn := helpers.NewBrokerConnection()
	defer conn.Close()

	ch := helpers.NewChannel(conn)
	defer ch.Close()

	q := helpers.NewQueue(ch)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := helpers.BodyFrom(os.Args)

	err := ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType:  "text/plain",
			DeliveryMode: amqp091.Persistent,
			Body:         []byte(body),
		},
	)
	helpers.FailOnError(err, "Failed to publish")
	log.Printf(" [x] Sent %s", body)
}
