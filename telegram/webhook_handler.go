package telegram

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/telegram-bot-api.v4"
	"io"
	"io/ioutil"
	"log"
)

func WebhookHandler(c *gin.Context) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(c.Request.Body)

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Println(err)
		return
	}

	UpdateHandler(&update)
}
