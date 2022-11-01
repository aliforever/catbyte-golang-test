package publishing_api

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

type publishingAPI struct {
	server      *gin.Engine
	rbmqChannel *amqp091.Channel
}

func NewAPIService(channel *amqp091.Channel) *publishingAPI {
	return &publishingAPI{
		server:      gin.Default(),
		rbmqChannel: channel,
	}
}

func (a *publishingAPI) RegisterRouters() {
	a.server.POST("/message", a.MessageHandler)
}

func (a *publishingAPI) Start(address ...string) error {
	a.RegisterRouters()

	return a.server.Run(address...)
}
