package main

import (
	"github.com/atadzan/message-broker/helpers"
	"log"
)

func main() {
	conn := helpers.NewBrokerConnection()
	defer conn.Close()

	ch := helpers.NewChannel(conn)
	defer ch.Close()

	q := helpers.NewQueue(ch)

	// Consuming messages
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s\n", d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C\n")
	<-forever
}
