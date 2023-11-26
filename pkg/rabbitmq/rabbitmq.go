package rabbitmq

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ashkan-jafarzadeh/delay/config"
)

type Service struct {
	config        config.Rabbitmq
	RabbitCon     *amqp.Connection
	RabbitChannel *amqp.Channel
}

func New(rabbitCon *amqp.Connection, config config.Rabbitmq) *Service {
	return &Service{
		RabbitCon: rabbitCon,
		config:    config,
	}
}

func (r *Service) InitChannel() error {
	qChannel, err := r.RabbitCon.Channel()

	if err != nil {
		return err
	}

	err = qChannel.ExchangeDeclare(r.config.ExchangeName, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}

	q, err := qChannel.QueueDeclare(
		r.config.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = qChannel.QueueBind(
		q.Name,
		r.config.RoutingKey,
		r.config.ExchangeName,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	r.RabbitChannel = qChannel

	return err
}

func (r *Service) Publish(ctx context.Context, data any) error {
	message, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.RabbitChannel.PublishWithContext(ctx,
		r.config.ExchangeName,
		r.config.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
}

func (r *Service) Consume(ctx context.Context) (<-chan amqp.Delivery, error) {
	return r.RabbitChannel.ConsumeWithContext(ctx,
		r.config.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}

func (r *Service) Get() (amqp.Delivery, bool, error) {
	return r.RabbitChannel.Get(r.config.QueueName, true)
}
