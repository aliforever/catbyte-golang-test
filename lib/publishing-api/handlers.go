package publishing_api

import (
	"catbyte-golang-test/shared/entities"
	"catbyte-golang-test/shared/ginwriters"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

func (a *publishingAPI) MessageHandler(c *gin.Context) {
	message := &entities.Message{}

	err := c.Bind(&message)
	if err != nil {
		ginwriters.BadRequest(c, err.Error())
		return
	}

	if message.Sender == "" || message.Receiver == "" || message.Message == "" {
		ginwriters.BadRequest(c, "empty_parameters")
		return
	}

	j, err := json.Marshal(message)
	if err != nil {
		ginwriters.InternalServerError(c, err.Error())
	}

	err = a.rbmqChannel.PublishWithContext(c, "catbyte", "", false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        j,
	})
	if err != nil {
		ginwriters.InternalServerError(c, err.Error())
		return
	}

	ginwriters.SuccessfulRequest(c, "written to rabbitmq successfully")
}
