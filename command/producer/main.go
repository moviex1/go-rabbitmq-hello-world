package main

import (
	"time"

	"example.com/logger"
	"example.com/rabbitmq"
)

func main() {
	conn, ch := rabbitmq.Initialize("amqp://guest:guest@localhost:5672/")
	defer conn.Close()
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"test",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		logger.Error("Failed to declare a queue!")
	}

	logger.Info("queue was created")

	err = rabbitmq.PublishMessageToQueue(ch, queue.Name, []byte("Hello world!"), 5*time.Second)

	if err != nil {
		logger.Error("failed to publish message")
	}

	logger.Info("message was sent to queue")
}
