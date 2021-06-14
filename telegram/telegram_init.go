package telegram

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/telegram-bot-api.v4"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	bot      *tgbotapi.BotAPI
	botToken = os.Getenv("TOKEN")
	baseURL  = os.Getenv("BASE_URL")
	err      error
)

func InitTelegram() (*tgbotapi.BotAPI, error) {
	bot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Println(err)
	}

	url := baseURL + bot.Token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))

	if err != nil {
		log.Println(err)
	}

	return bot, err
}

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

	msg := UpdateHandler(&update)
	_, err = bot.Send(msg)
	if err != nil {
		log.Println(err)
	}

}
