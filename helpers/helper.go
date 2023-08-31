package helpers

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"strings"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func NewBrokerConnection() *amqp091.Connection {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func NewChannel(conn *amqp091.Connection) *amqp091.Channel {
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	return ch
}

func NewQueue(ch *amqp091.Channel) amqp091.Queue {
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil)
	FailOnError(err, "Failed to declare a queue")
	return q
}

func BodyFrom(args []string) (s string) {
	if len(args) > 2 || os.Args[1] == "" {
		s = "Hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return
}
