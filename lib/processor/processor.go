package processor

import (
	"catbyte-golang-test/shared/entities"
	"catbyte-golang-test/shared/logger"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/rabbitmq/amqp091-go"
)

type processor struct {
	channel   *amqp091.Channel
	queueName string
	rc        *redis.Client
	logger    logger.Logger
}

func NewProcessor(channel *amqp091.Channel, queueName string, rc *redis.Client, l logger.Logger) *processor {
	if l == nil {
		l = logger.DefaultLogger{}
	}
	return &processor{channel: channel, queueName: queueName, rc: rc, logger: l}
}

func (p *processor) Start() error {
	messages, err := p.channel.Consume(p.queueName, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	for message := range messages {
		var m *entities.Message
		err := json.Unmarshal(message.Body, &m)
		if err != nil {
			p.logger.Error("cant unmarshal message from rabbitMQ: %s", err)
			continue
		}
		if m.Sender == "" || m.Receiver == "" || m.Message == "" {
			p.logger.Error(fmt.Sprintf("empty values found for rabbitMQ message id: %s", message.MessageId))
			continue
		}

		_, err = p.rc.LPush(context.Background(), fmt.Sprintf("catbyte_messages_%s_%s", m.Sender, m.Receiver), string(message.Body)).Result()
		if err != nil {
			return err
		}
	}

	return nil
}
