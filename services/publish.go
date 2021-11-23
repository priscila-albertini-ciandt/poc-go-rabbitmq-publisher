package services

import (
	"poc-rabbitmq/error"

	"github.com/streadway/amqp"
)

func Publish(ch *amqp.Channel, q amqp.Queue, message string) {
	body := message
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	error.FailOnError(err, "Failed to publish a message")
}
