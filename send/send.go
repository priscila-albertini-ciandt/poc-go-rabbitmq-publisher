package send

import (
	"poc-rabbitmq/config"
	"poc-rabbitmq/services"
)

func Send() {
	conn := config.Connect()
	defer conn.Close()

	ch := services.OpenChannel(conn)
	defer ch.Close()

	q := services.CreateQueue(ch, "queue.hello")
	services.Publish(ch, q, "Hello World!")
}
