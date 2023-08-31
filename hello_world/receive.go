package main

import (
	"github.com/atadzan/message-broker/hello_world/helper"
	"log"
)

func main() {
	conn := helper.NewBrokerConnection()
	defer conn.Close()

	ch := helper.NewChannel(conn)
	defer ch.Close()

	q := helper.NewQueue(ch)

	// Consuming messages
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	helper.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s\n", d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C\n")
	<-forever
}
