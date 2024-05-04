package rabbitmq

import (
	"example.com/logger"
	"github.com/rabbitmq/amqp091-go"
)

func Initialize(url string) (*amqp091.Connection, *amqp091.Channel) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		logger.Error("cannot connect to rabbitmq!")
	} else {
		logger.Info("successfully connected to rabbitmq!")
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Error("cannot open a channel!")
	} else {
		logger.Info("successfully opened a channel!")
	}

	return conn, ch
}
