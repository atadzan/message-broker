package main

import (
	"bytes"
	"github.com/atadzan/message-broker/helpers"
	"log"
	"time"
)

func main() {
	conn := helpers.NewBrokerConnection()
	defer conn.Close()

	ch := helpers.NewChannel(conn)
	defer ch.Close()

	q := helpers.NewQueue(ch)

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
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Println("Done")
		}
	}()
	log.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
