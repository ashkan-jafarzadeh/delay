package rabbit

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ashkan-jafarzadeh/delay/config"
	"net"
)

func NewConnection(conf config.Rabbitmq) (*amqp.Connection, error) {
	host := net.JoinHostPort(conf.Host, conf.Port)

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s", conf.Username, conf.Password, host))

	if err != nil {
		return nil, err
	}

	return conn, nil
}
