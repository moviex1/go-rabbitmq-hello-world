package rabbitmq

import (
	"context"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func PublishMessageToQueue(ch *amqp091.Channel, queueName string, msg []byte, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return ch.PublishWithContext(
		ctx,
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
}
