package reporting_api

import (
	"catbyte-golang-test/shared/entities"
	"catbyte-golang-test/shared/ginwriters"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *reportingAPI) ListMessagesHandler(c *gin.Context) {
	type ListMessageQuery struct {
		Sender   string `form:"sender"`
		Receiver string `form:"receiver"`
	}

	listMessageQuery := &ListMessageQuery{}

	err := c.Bind(&listMessageQuery)
	if err != nil {
		ginwriters.BadRequest(c, err.Error())
		return
	}

	if listMessageQuery.Sender == "" || listMessageQuery.Receiver == "" {
		ginwriters.BadRequest(c, "empty_parameters")
		return
	}

	key := fmt.Sprintf("catbyte_messages_%s_%s", listMessageQuery.Sender, listMessageQuery.Receiver)

	exists, err := a.rc.Exists(c, key).Result()
	if err != nil {
		ginwriters.InternalServerError(c, err.Error())
		return
	}

	if exists == 0 {
		ginwriters.BadRequest(c, "no messages found for provided sender and receiver")
		return
	}

	data, err := a.rc.LRange(c, key, 0, -1).Result()
	if err != nil {
		ginwriters.InternalServerError(c, err.Error())
		return
	}

	var messages []entities.Message
	for _, message := range data {
		var m entities.Message
		err := json.Unmarshal([]byte(message), &m)
		if err != nil {
			a.logger.Error(fmt.Errorf("cant decode message from redis: %s", err))
			continue
		}
		messages = append(messages, m)
	}

	c.JSON(http.StatusOK, messages)
}
