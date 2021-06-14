package telegram

import (
	"gopkg.in/telegram-bot-api.v4"
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
